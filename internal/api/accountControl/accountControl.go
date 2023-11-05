package accountControl

import (
	"TalkHub/internal/api/accountControl/pb"
	"TalkHub/internal/config"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GRPCClient struct {
	Client pb.AccountControlClient
}

var (
	errConnectionToServer = errors.New("connection failed to grpc server")
)

func initGRPCClient(cfg *config.GRPCConfig) (*GRPCClient, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil || !ping(conn) {
		return nil, errConnectionToServer
	}

	client := pb.NewAccountControlClient(conn)
	return &GRPCClient{Client: client}, nil
}

func ping(conn *grpc.ClientConn) bool {
	for conn.GetState() == connectivity.Connecting {
		time.Sleep(time.Millisecond * 10)
	}
	return conn.GetState() == connectivity.Ready
}

func (c *GRPCClient) Registration(email, password string) (session *pb.SessionData, strErr string) {
	return authentication(email, password, c.Client.RegistrationAccount)
}

func (c *GRPCClient) Authorization(email, password string) (session *pb.SessionData, strErr string) {
	return authentication(email, password, c.Client.AuthorizationAccount)
}

func authentication(email, password string, ff func(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.SessionData, error)) (session *pb.SessionData, strErr string) {
	user := &pb.User{
		Email:    email,
		Password: password,
	}

	session, err := ff(context.Background(), user)
	if err != nil {
		strErr = grpcErrToString(err)
		return nil, strErr
	}

	return session, ""
}

func grpcErrToString(err error) string {
	return err.Error()[33:]
}

func (c *GRPCClient) ChangePassword(email, password, newPassword string) (strErr string) {
	data := &pb.ChangePasswordData{
		Email:       email,
		Password:    password,
		NewPassword: newPassword,
	}

	_, err := c.Client.ChangePasswordAccount(context.Background(), data)
	if err != nil {
		return grpcErrToString(err)
	}

	return ""
}

func (c *GRPCClient) DeleteAccount(id, email, password string) {
	info := &pb.FullInfoUser{
		Id:       id,
		Email:    email,
		Password: password,
	}

	_, _ = c.Client.DeleteAccount(context.Background(), info)
}

func (c *GRPCClient) IsAuthenticated(key string) (id string, strErr string) {
	session := &pb.SessionData{
		Key: key,
	}

	data, err := c.Client.IsAuthorizedSessionData(context.Background(), session)
	if err != nil {
		strErr = grpcErrToString(err)
		return "", strErr
	}

	return data.Id, ""
}

func (c *GRPCClient) DeleteSession(key string) {
	session := &pb.SessionData{
		Key: key,
	}

	_, _ = c.Client.DeleteSessionData(context.Background(), session)
}
