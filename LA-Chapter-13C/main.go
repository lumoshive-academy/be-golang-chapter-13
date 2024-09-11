package main

import (
	"context"
	"fmt"
)

func main() {
	// Membuat parent context
	parentCtx := context.Background()

	// Mendefinisikan tipe unik untuk key
	type key string

	// Membuat child context pertama dengan nilai
	ctx1 := context.WithValue(parentCtx, key("key1"), "value1")

	// Membuat child context kedua dengan nilai
	ctx2 := context.WithValue(ctx1, key("key2"), "value2")

	// Membuat child context ketiga dengan nilai
	ctx3 := context.WithValue(ctx2, key("key3"), "value3")

	// Mengakses nilai dari context ketiga
	printContextValue(ctx3, key("key1"))
	printContextValue(ctx3, key("key2"))
	printContextValue(ctx3, key("key3"))
}

// Fungsi untuk mencetak nilai dari context
func printContextValue(ctx context.Context, k interface{}) {
	value := ctx.Value(k)
	if value != nil {
		fmt.Printf("Key: %v, Value: %v\n", k, value)
	} else {
		fmt.Printf("Key: %v tidak ditemukan dalam context\n", k)
	}
}
