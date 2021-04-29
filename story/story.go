package main

import (
	"fmt"
)

type data struct {
}

func main() {
	var nama string
	var nama_kasir string
	var tanggal string
	var harga int
	var jumlah_item = make(map[string]int)
	var item string

	fmt.Println("")
	fmt.Println("Masukkan Nama Toko")
	fmt.Scanln(&nama)
	fmt.Print("Masukkan Nama Kasir")
	fmt.Scanln(nama_kasir)

	fmt.Println("Tanggal")
	fmt.Scanln(&tanggal)
	for {
		fmt.Scanln(&item)
		if item == "keluar" {
			break
		}
		fmt.Println("Masukkan " + item + " harga: ")
		fmt.Scanln(&harga)
		jumlah_item[item] = harga
	}
	var jumlah int

	fmt.Println("Total...................")
	fmt.Println()
	fmt.Println(nama)
	fmt.Println("Tanggal :          ", tanggal)
	fmt.Println("Nama Kasir :     ", nama_kasir)
	fmt.Println("============================")
	for i, j := range jumlah_item {
		fmt.Println(i + "...................")
		jumlah += j
	}

}
