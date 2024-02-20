package main

import (
	"fmt"
)

func isOne(nomor int) {
	if nomor == 1 {
		fmt.Print("Ini satu")
	} else {
		fmt.Print("Ini bukan satu")
	}
}

func main() {
	var (
		nomor1 int
	)
	fmt.Println("Masukkan nomor 1:")
	fmt.Scan(&nomor1)
	isOne(nomor1)
}
