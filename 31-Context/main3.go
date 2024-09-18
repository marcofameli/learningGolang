package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "senha")
	bookHotell(ctx, "Hotel")
}

// CONTEXTO SEMPRE VEM COMO PRIMEIRO NO PARAMETRO
func bookHotell(ctx context.Context, hotel string) {
	token := ctx.Value("token")
	fmt.Println(token)
}
