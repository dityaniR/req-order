package skeleton

import (
	"context"
	"request-order/internal/entity/auth"
)

type Data interface {
	CheckAuth(ctx context.Context, _token, code string) (auth.Auth, error)
}

type Service struct {
	data Data
}

func New(data Data) Service {
	return Service{
		data: data,
	}
}
