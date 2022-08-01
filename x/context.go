package x

import (
	"context"
	"time"
)

type Context func() (context.Context, context.CancelFunc)

func ContextTimeoutDefault() Context {
	return ContextTimeout(5)
}

func ContextTimeout(second int) Context {
	return func() (context.Context, context.CancelFunc) {
		return context.WithTimeout(context.Background(), time.Second*time.Duration(second))
	}
}
