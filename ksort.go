//Write a program to sort an array of integers. The program should partition the array into 4 parts, each of which is
//sorted by a different goroutine. Each partition should be of approximately equal size. Then the main goroutine should
//merge the 4 sorted subarrays into one large sorted array.

//The program should prompt the user to input a series of integers. Each goroutine which sorts ¼ of the array should
//print the subarray that it will sort. When sorting is complete, the main goroutine should print the entire sorted list.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type ArrayList struct {
	currentIndex int
	bucket []int
}

func (list *ArrayList) hasNext() bool {
	return list.currentIndex < len(list.bucket)
}

func (list *ArrayList) current() int {
	return list.bucket[list.currentIndex]
}

func (list * ArrayList) pop() {
	list.advancePointer()
}

func (list *ArrayList) advancePointer() {
	list.currentIndex++
}

func (list *ArrayList) retreatPointer() {
	if list.currentIndex > 0 {
		list.currentIndex--
	}
}

func main() {

	fmt.Println("!!! Welcome Merging K Sorted Arrays !!!")
	fmt.Println("Please enter a sequence of at least four or more integers to be sorted ")

	inputScanner := setUpConsoleScanner()
	unsortedStrings := getSequence(inputScanner)

	validateList(&unsortedStrings)

	unsortedSequence := convertToListOfInts(&unsortedStrings)

	chunk1, chunk2, chunk3, chunk4 := chunkUnsortedSequence(unsortedSequence)

	alpha := make(chan []int)
	beta := make(chan []int)
	charlie := make(chan []int)
	delta := make(chan []int)

	go sortArray(chunk1, alpha)
	go sortArray(chunk2, beta)
	go sortArray(chunk3, charlie)
	go sortArray(chunk4, delta)

	sortedChunk1 := <- alpha
	sortedChunk2 := <- beta
	sortedChunk3 := <- charlie
	sortedChunk4 := <- delta

	var sortBucketOne = ArrayList{0, sortedChunk1}
	var sortBucketTwo = ArrayList{0, sortedChunk2}
	var sortBucketThree = ArrayList{0, sortedChunk3}
	var sortBucketFour = ArrayList{0, sortedChunk4}

	mergedSequence := make([]int, 0)

	var lowestSortArray *ArrayList
	lowestSortArray = &sortBucketOne

	for i:=0; i < len(unsortedSequence) ; i++ {

		val := lowestSortBucketValue(lowestSortArray, &sortBucketOne, &sortBucketTwo, &sortBucketThree, &sortBucketFour)
		mergedSequence = append(mergedSequence, val)
	}
	fmt.Println(mergedSequence)
}

func lowestSortBucketValue(lowest *ArrayList ,one *ArrayList, two *ArrayList, three *ArrayList, four *ArrayList) int {

	if one.hasNext() && one.current() < lowest.current() {
		lowest = one
	}

	if two.hasNext() && two.current() < lowest.current(){
		lowest = two
	}

	if three.hasNext() && three.current() < lowest.current() {
		lowest = three
	}

	if four.hasNext() && four.current() < lowest.current() {
		lowest = four
	}

	lowestValue := lowest.current()
	lowest.pop()

	if !lowest.hasNext() {
		lowest = one

		if !lowest.hasNext() {
			lowest = two

			if !lowest.hasNext() {
				lowest = three

				if !lowest.hasNext(){
					lowest = four
				}
			}
		}
	}

	return lowestValue
}


func chunkUnsortedSequence(sequence []int) ([]int,[]int,[]int,[]int)  {

	var chunks [4][]int
	
	for i := 0; i < len(sequence); i++ {

		chunks[i % 4] = append(chunks[i % 4], sequence[i])
	}

	return chunks[0] , chunks[1], chunks[2], chunks[3]
}

func sortArray( unsortedArray []int, channel chan []int) {
	
	fmt.Println(channel, "is sorting the following array", unsortedArray)
	
	BubbleSort(&unsortedArray)
	channel <- unsortedArray
}


func BubbleSort(unsortedSequence *[]int) {

	unsorted := true

	for unsorted {

		for index  := 0; index < len(*unsortedSequence) - 1 ; index++ {

			Swap(index , unsortedSequence)
		}

		unsorted = isSequenceUnSorted(unsortedSequence)
	}
}


func Swap(index int, ints *[]int) {

	firstNum := (*ints)[index]
	secondNum := (*ints)[index + 1]

	if firstNum > secondNum {
		(*ints)[index] = secondNum
		(*ints)[index + 1] = firstNum
	}
}

func isSequenceUnSorted(unsortedSequence *[]int) (unsorted bool) {

	unsorted = false

	for index  := 1; index < len(*unsortedSequence); index++ {

		if (*unsortedSequence)[index - 1] > (*unsortedSequence)[index] {
			unsorted = true
			break
		}
	}

	return
}

func convertToListOfInts(list *[]string) (ints []int) {

	ints = make([]int, 0)

	for _, str := range *list {

		num , _ := strconv.ParseInt(str, 10, 64 )

		ints = append(ints , int(num))
	}
	return
}

func setUpConsoleScanner() *bufio.Scanner {

	inputScanner := bufio.NewScanner(os.Stdin)
	return inputScanner
}

func getSequence(inputScanner *bufio.Scanner) []string {
	inputScanner.Scan()
	return strings.Split(strings.TrimSpace(inputScanner.Text()), " ")
}

func validateList(list *[]string) {

	for _,str := range *list {
		_,err := strconv.ParseInt(str, 10, 64)

		if err != nil {
			panic(err)
		}
	}
}