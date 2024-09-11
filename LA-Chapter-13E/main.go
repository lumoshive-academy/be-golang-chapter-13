package main

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

func main() {
	// Membuat parent context
	parentCtx := context.Background()

	// Membuat context dengan batas waktu 3 detik
	ctx, cancel := context.WithTimeout(parentCtx, 3*time.Second)
	defer cancel() // Pastikan untuk memanggil cancel untuk membersihkan sumber daya

	// URL layanan yang akan diperiksa
	url := "https://jsonplaceholder.typicode.com/todos/1"

	// Membuat request HTTP dengan context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Membuat client HTTP
	client := &http.Client{}

	// Melakukan request dengan context
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	}
	defer resp.Body.Close()

	// Memeriksa status response
	if resp.StatusCode == http.StatusOK {
		fmt.Println("Layanan tersedia")
	} else {
		fmt.Println("Layanan tidak tersedia")
	}
}
