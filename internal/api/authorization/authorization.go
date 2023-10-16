package authorization

import (
	"TalkHub/internal/api/authorization/pb"
	"TalkHub/internal/config"
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"time"
)

type GRPCClient struct {
	Client pb.AuthorizationClient
}

var (
	errConnectionToServer = errors.New("connection failed to grpc server")
)

func initGRPCClient(cfg *config.GRPCConfig) (*GRPCClient, error) {
	conn, err := grpc.Dial(cfg.Address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil || !ping(conn) {
		return nil, errConnectionToServer
	}

	client := pb.NewAuthorizationClient(conn)
	return &GRPCClient{Client: client}, nil
}

func ping(conn *grpc.ClientConn) bool {
	for conn.GetState() == connectivity.Connecting {
		time.Sleep(time.Millisecond * 10)
	}
	return conn.GetState() == connectivity.Ready
}

func (c *GRPCClient) Register(login, password string) (ket string, strErr string) {
	return authorization(login, password, c.Client.Register)
}

func (c *GRPCClient) LogIn(login, password string) (key string, strErr string) {
	return authorization(login, password, c.Client.LogIn)
}

func authorization(login, password string, ff func(ctx context.Context, in *pb.User, opts ...grpc.CallOption) (*pb.SessionData, error)) (key string, strErr string) {
	user := &pb.User{
		Login:    login,
		Password: password,
	}

	session, err := ff(context.Background(), user)
	if err != nil {
		strErr = grpcErrToString(err)
		return "", strErr
	}

	return session.Key, ""
}

func grpcErrToString(err error) string {
	return err.Error()[33:]
}

func (c *GRPCClient) IsAuthenticated(key string) (login string, strErr string) {
	session := &pb.SessionData{
		Key: key,
	}

	data, err := c.Client.IsAuthenticated(context.Background(), session)
	if err != nil {
		strErr = grpcErrToString(err)
		return "", strErr
	}

	return data.Login, ""
}

func (c *GRPCClient) ChangePassword(login, password, newPassword string) (strErr string) {
	data := &pb.ChangePasswordData{
		Login:       login,
		Password:    password,
		NewPassword: newPassword,
	}

	_, err := c.Client.ChangePassword(context.Background(), data)
	strErr = grpcErrToString(err)
	return strErr
}
