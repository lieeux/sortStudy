package main

import (
	"fmt"
	"math/rand"
	"sortStudy/tools"
	"time"
)

func main() {
	var arr [100000]int
	for i := 0; i < 100000; i++ {
		arr[i] = rand.Intn(900000)
	}
	start := time.Now().Unix()
	tools.Quick(0, len(arr)-1, &arr)
	end := time.Now().Unix()
	fmt.Printf("耗时%d秒", end-start)
}
