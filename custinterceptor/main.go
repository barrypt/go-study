package main

import (
	"context"
	"fmt"
)

type invoker func(ctx context.Context, interceptors []interceptor) error

//type handler func(ctx context.Context)
type interceptor func(ctx context.Context, ivk invoker) error

func getInvoker(ctx context.Context, interceptors []interceptor, curr int, ivk invoker) invoker {
	if curr == len(interceptors)-1 {
		return ivk
	}
	return func(ctx context.Context, interceptors []interceptor) error {
		return interceptors[curr+1](ctx, getInvoker(ctx, interceptors, curr+1, ivk))
	}
}

func getChainInterceptor(ctx context.Context, interceptors []interceptor, ivk invoker) interceptor {
	if len(interceptors) == 0 {
		return nil
	} else if len(interceptors) == 1 {
		return interceptors[0]
	} else {
		return func(ctx context.Context, ivk invoker) error {
			return interceptors[0](ctx, getInvoker(ctx, interceptors, 0, ivk))
		}
	}
}

func main() {
	var ctx context.Context
	var ceps []interceptor

	var interceptor1 = func(ctx context.Context, ivk invoker) error {
		fmt.Println("loginc1")
		ivk(ctx, ceps)
		fmt.Println("loginc1after")
		return nil
	}
	var interceptor2 = func(ctx context.Context, ivk invoker) error {
		fmt.Println("loginc2")
		ivk(ctx, ceps)
		fmt.Println("loginc2after")
		return nil
	}
	ceps = append(ceps, interceptor1, interceptor2)

	var ivk = func(ctx context.Context, interceptors []interceptor) error {
		fmt.Println("invoker start")
		return nil
	}

	cep := getChainInterceptor(ctx, ceps, ivk)
	cep(ctx, ivk)
}
