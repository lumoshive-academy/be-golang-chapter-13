package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Membuat parent context
	parentCtx := context.Background()

	// Membuat context dengan cancel
	ctx, cancel := context.WithCancel(parentCtx)

	// Fungsi defer untuk memastikan cancel dipanggil untuk membersihkan sumber daya
	defer cancel()

	// Menjalankan beberapa operasi yang menggunakan context
	go worker(ctx, "Worker 1")
	go worker(ctx, "Worker 2")
	go worker(ctx, "Worker 3")

	// Menunggu sebentar sebelum membatalkan context
	time.Sleep(2 * time.Second)
	fmt.Println("Membatalkan context...")
	cancel() // Membatalkan context

	// Menunggu sebentar untuk melihat hasil pembatalan
	time.Sleep(1 * time.Second)
	fmt.Println("Aplikasi selesai")
}

func worker(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			// Sinyal pembatalan diterima
			fmt.Printf("%s dibatalkan\n", name)
			return
		default:
			// Melakukan pekerjaan
			fmt.Printf("%s sedang bekerja\n", name)
			time.Sleep(500 * time.Millisecond)
		}
	}
}
