package httpx

import "context"

type contextKey string

const requestIDKey contextKey = "requestID"

func WithRequstID(ctx context.Context, id string) context.Context {
	return context.WithValue(ctx, requestIDKey, id)
}

func GetRequestID(ctx context.Context) (string, bool) {
	v := ctx.Value(requestIDKey)
	id, ok := v.(string)
	return id, ok
}
