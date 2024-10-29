package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	fmt.Println("Doing something!", ctx.Err())
}

func main() {
	ctx := context.TODO()
	doSomething(ctx)
}
