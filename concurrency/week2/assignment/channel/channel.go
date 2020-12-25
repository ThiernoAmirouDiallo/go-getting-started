package main

import (
	"fmt"
	sortpkg "sort"
	"strconv"
	"strings"
)

const PARTITION_NUMBER int = 4

func main() {

	ints := GetPopulatedSlice()

	if len(ints) <= 1 {
		fmt.Println("No need of sorting")
	} else {
		ints = PartitionAndSort(ints, PARTITION_NUMBER)
	}

	fmt.Println(ints)
}

func PartitionAndSort(ints []int, partNumber int) []int {
	c := make(chan []int, partNumber)

	partitionSize := len(ints) / partNumber
	if len(ints)%partNumber != 0 {
		partitionSize++
	}
	var partitionNum, startIdx, endIdx = 0, 0, 0

	for i := 0; endIdx < len(ints); i++ {
		startIdx = i * partitionSize
		endIdx = (i + 1) * partitionSize
		partitionNum++
		if i+1 == partNumber {
			endIdx = len(ints)
		}
		var partition = make([]int, endIdx-startIdx)
		copy(partition, ints[startIdx:endIdx])

		go sort(&c, partition, partitionNum)
	}

	sortedPartitions := make([][]int, partitionNum)
	for i := 0; i < partitionNum; i++ {
		sortedPartitions[i] = <-c
	}

	mergedPartitions := make([]int, 0)

	for i := 0; i < partitionNum; i++ {
		mergedPartitions = merge(mergedPartitions, sortedPartitions[i])
	}

	return mergedPartitions
}

func sort(c *chan []int, t []int, id int) {
	fmt.Printf("Goroutine-%v - sorting %v\n", id, t)
	if len(t) > 1 {
		sortpkg.Ints(t)
	}
	*c <- t
}

func merge(t1 []int, t2 []int) []int {
	res := make([]int, len(t1)+len(t2))
	i, j := 0, 0

	for i < len(t1) && j < len(t2) {
		if t1[i] < t2[j] {
			res[i+j] = t1[i]
			i++
		} else {
			res[i+j] = t2[j]
			j++
		}
	}

	if i < len(t1) {
		copy(res[i+j:], t1[i:])
	} else {
		copy(res[i+j:], t2[j:])
	}

	return res
}

func GetPopulatedSlice() []int {
	numbers := make([]int, 0, 10)
	for {
		var enteredString string
		for {
			fmt.Print("Please enter a number or X to finish : ")
			_, err := fmt.Scanf("%s", &enteredString)

			if err == nil {
				break
			} else {
				fmt.Println("Error while reading your input")
			}
		}

		if strings.ToUpper(enteredString) == "X" {
			break
		}

		if n, errConv := strconv.Atoi(enteredString); errConv != nil {
			fmt.Println("Error while converting your input to integer")
			continue
		} else {
			numbers = append(numbers, n)
		}
	}

	return numbers
}
