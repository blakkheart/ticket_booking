package httpx

import (
	"context"
)

type contextKey string

const requestIDKey contextKey = "requestID"
const userIDKey contextKey = "userID"

func WithRequstID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func GetRequestID(ctx context.Context) (string, bool) {
	v := ctx.Value(requestIDKey)
	id, ok := v.(string)
	return id, ok
}

func WithUserID(ctx context.Context, id *int64) context.Context {
	if id == nil {
		return ctx
	}
	return context.WithValue(ctx, userIDKey, *id)
}

func GetUserID(ctx context.Context) (int64, bool) {
	v := ctx.Value(userIDKey)
	id, ok := v.(int64)
	return id, ok
}
