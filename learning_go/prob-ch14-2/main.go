package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	const format = "Sum: %d, Loops: %d, Reason: %s\n"

	var sum uint64

	for i := 1; ; i++ {
		n := r.Int63n(100_000_000)
		sum += uint64(n)

		// time.Sleep(2 * time.Millisecond) // 2秒経過させやすくするため

		if ctx.Err() != nil {
			fmt.Printf(format, sum, i, "2 seconds passed.")
			return
		} else if n == 1234 {
			fmt.Printf(format, sum, i, "1234 generated.")
			return
		}
	}
}
