package main

import (
	"fmt"
)

type data struct {
}

func main() {

	var jumlah_item = make(map[string]int)

	fmt.Println("")
	fmt.Println("Masukkan Nama Toko")
	var nama string
	fmt.Scanln(&nama)
	fmt.Println("Masukkan Nama Kasir")
	var nama_kasir string
	fmt.Scanln(nama_kasir)

	fmt.Println("Tanggal")
	var tanggal string
	fmt.Scanln(&tanggal)
	fmt.Println("Masukkan nama barang")
	for {
		var item string
		fmt.Scanln(&item)
		if item == "keluar" {
			break
		}
		fmt.Println("Masukkan " + item + " harga: ")
		var harga int
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
