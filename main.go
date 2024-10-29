package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	doAnother(anotherCtx)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func doAnother(ctx context.Context) {
	fmt.Printf("doAnother: myKey's value is %s\n", ctx.Value("myKey"))
}

func main() {
	// ctx := context.TODO()

	// The context.Background function creates an empty context like context.TODO does, but it’s designed to be used where you intend to start a known context.
	ctx := context.Background()

  // The values stored in a specific context.Context are immutable, meaning they can’t be changed.
	ctx = context.WithValue(ctx, "myKey", "myValue")

	doSomething(ctx)
}
