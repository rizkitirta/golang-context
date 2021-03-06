package golangcontext

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

func TestContext(t *testing.T) {
	background := context.Background()
	fmt.Println(background)

	todo := context.TODO()
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	contextB := context.WithValue(contextA, "b", "A")
	contextC := context.WithValue(contextA, "c", "C")

	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	fmt.Println(contextA)
	fmt.Println(contextB)
	fmt.Println(contextC)
	fmt.Println(contextD)
	fmt.Println(contextE)
}

func RunCounter(ctx context.Context) chan int {
	destination := make(chan int)
	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return

			default:
				destination <- counter
				counter++
			}
		}
	}()

	return destination
}

func TestContextWithCancel(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	ctx, cancel := context.WithCancel(context.Background())
	destination := RunCounter(ctx)

	for v := range destination {
		fmt.Println("counter", v)
		if v == 10 {
			break
		}
	}

	cancel() // Mengirim sinyal cancel ke goroutine yang menjalankan counter
	fmt.Println("total goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeOut(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // cancel akan otomatis diajalankan setelah waktu 5 detik
	destination := RunCounter(ctx)
	defer cancel()

	for v := range destination {
		fmt.Println("counter", v)
	}

	fmt.Println("total goroutine", runtime.NumGoroutine())
}

func TestContextWithDeadline(t *testing.T) {
	fmt.Println("total goroutine", runtime.NumGoroutine())

	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second)) // cancel akan otomatis diajalankan sesuai waktu deadline yang ditentukan
	destination := RunCounter(ctx)
	defer cancel()

	for v := range destination {
		fmt.Println("counter", v)
	}

	fmt.Println("total goroutine", runtime.NumGoroutine())
}
