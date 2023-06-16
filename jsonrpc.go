package hermesaddons

import (
	"context"
	"encoding/json"
)

type JsonRPCGetter interface {
	Form(id uint) (json.Marshaler, error)
	Call(ctx context.Context, id uint, data []byte, user UserGetter) (json.Marshaler, error)
}
