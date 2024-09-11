package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// Membuat parent context
	parentCtx := context.Background()

	// Menentukan waktu deadline: 5 detik dari sekarang
	deadline := time.Now().Add(5 * time.Second)

	// Membuat context dengan deadline
	ctx, cancel := context.WithDeadline(parentCtx, deadline)
	// Fungsi defer untuk memastikan cancel dipanggil untuk membersihkan sumber daya
	defer cancel()

	// Menjalankan operasi yang menggunakan context
	result := process(ctx)
	fmt.Println("Hasil operasi:", result)
}

func process(ctx context.Context) string {
	select {
	case <-time.After(3 * time.Second):
		return "Proses selesai tepat waktu"
	case <-ctx.Done():
		// Context dibatalkan (termasuk karena deadline)
		return "Proses dibatalkan karena melewati deadline"
	}
}
