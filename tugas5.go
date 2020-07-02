package main

import (
	"fmt"
	"math/rand"
)

type pelajar struct {
	nama string
	umur int
}

func main() {
	var n1, n2, n3 pelajar
	n1.nama = "Ahmad"
	n2.nama = "Abi"
	n3.nama = "Zaka"

	n1.umur = rand.Intn(100)
	n2.umur = rand.Intn(100)
	n3.umur = rand.Intn(100)

	n1.tampil()
	n2.tampil()
	n3.tampil()

}

func (g pelajar) tampil() {
	fmt.Println(g.nama, g.umur)
}
