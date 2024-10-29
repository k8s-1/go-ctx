package main

import (
	"context"
	"fmt"
)

func doSomething(ctx context.Context) {
	ctx, cancelCtx := context.WithCancel(ctx)
	
	printCh := make(chan int)
	go doAnother(ctx, printCh)

	for num := 1; num <= 3; num++ {
		printCh <- num
	}

	cancelCtx()

	time.Sleep(100 * time.Millisecond)

	fmt.Printf("doSomething: finished\n")
}

func doAnother(ctx context.Context, printCh <-chan int) {
	for {
		select {
		case <-ctx.Done():
			if err := ctx.Err(); err != nil {
				fmt.Printf("doAnother err: %s\n", err)
			}
			fmt.Printf("doAnother: finished\n")
			return
		case num := <-printCh:
			fmt.Printf("doAnother: %d\n", num)
		}
	}
}

func processCtxValues(ctx context.Context) {
	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))

	anotherCtx := context.WithValue(ctx, "myKey", "anotherValue")
	processOtherCtxValues(anotherCtx)

	fmt.Printf("doSomething: myKey's value is %s\n", ctx.Value("myKey"))
}

func processOtherCtxValues(ctx context.Context) {
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
