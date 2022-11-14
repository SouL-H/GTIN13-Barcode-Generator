package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	//First 8 digit
	barcodeFlag := "12345678"

	//Create Last 4 Digit
	numberFlag := lastFourDigits()

	//BarcodeFlag + NumberFlag + CheckSum
	//12345678 + 0000 + 0 example

	//Create 100000 Barcode
	for i := 0; i <= 9999; i++ {
		barcode := barcodeCreate(barcodeFlag + numberFlag[i])
		fmt.Println(barcode)
	}
}

func lastFourDigits() []string {
	var first string = "0000"
	var numbers []string
	for i := 0; i <= 9999; i++ {
		first = "0000" + strconv.Itoa(i)
		numbers = append(numbers, first[len(first)-4:])
	}
	return numbers

}

func barcodeCreate(a string) string {
	totalOdd := 0
	totalEven := 0
	odd := make([]int, 0, 10)
	even := make([]int, 0, 10)

	for i := 0; i < len(a); i++ {
		if i%2 == 0 {
			intVar, _ := strconv.Atoi(string(a[i]))
			even = append(even, intVar)
		} else {
			intVar, _ := strconv.Atoi(string(a[i]))
			odd = append(odd, intVar)
		}
	}
	for i := 0; i < len(even); i++ {
		totalEven += even[i]
	}

	for i := 0; i < len(odd); i++ {
		totalOdd += odd[i]
	}
	oddMultiplication := totalOdd * 3

	total := oddMultiplication + totalEven
	nearestNum := math.Ceil(float64(total)/10) * 10
	checkSum := int(nearestNum) - total
	barcode := a + strconv.Itoa(checkSum)
	return barcode
}
