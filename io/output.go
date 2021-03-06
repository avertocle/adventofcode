package io

import "fmt"

func PrettyArray2DInt(arr [][]int) {
	for _, row := range arr {
		for _, cell := range row {
			fmt.Printf("%v ", cell)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrettyArray2DByte(arr [][]byte) {
	for _, row := range arr {
		fmt.Printf("%v\n", string(row))
	}
	fmt.Println()
}

func PrettyArray2DString(arr [][]string) {
	for _, row := range arr {
		for _, ele := range row {
			fmt.Printf("%v | ", ele)
		}
		fmt.Println()
	}
	fmt.Println()
}

func PrettyArray3DInt(arr [][][]int) {
	for _, d1 := range arr {
		PrettyArray2DInt(d1)
		fmt.Println()
	}
	fmt.Println()
}
