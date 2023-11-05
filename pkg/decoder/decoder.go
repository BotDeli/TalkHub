package decoder

import (
	"encoding/json"
	"io"
)

func JSONDecoder[T any](r io.ReadCloser) (*T, error) {
	var data T

	bytes, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}
