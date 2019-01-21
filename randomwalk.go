package main

import (
	"awesomeProject/plot"
	"awesomeProject/pricedata"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("hogehoge")

	data := pricedata.CreateRandomData(10000, 100000)
	err := plot.PlotGraph("example", "x", "y", "price", data, "points.png")
	if err!=nil {
		os.Exit(1)
	}

	//numTry := 10000
	//stock := 0
	//paid := 0
	//
	//// ドルコスト平均法
	//value := numTry
	//for i := 0; i < numTry; i++ {
	//	_ = changeValue(&value)
	//
	//	paid += value
	//	stock += 1
	//}
	//printResult(paid, stock, value)
}


func changeValue(value *int) bool {
	if (rand.Intn(2) == 0) {
		// up
		*value++
		return true
	} else {
		*value--
		return false
	}
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printResult(paid, stock, value int) {
	fmt.Printf("paid=%d, stock=%d, current value=%d, profit=%d", paid, stock, value, calcProfit(paid, stock, value))
}

func calcProfit(paid, stock, value int) int {
	return (stock * value) - paid
}
