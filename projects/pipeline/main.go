package main

import "fmt"

func removeDuplicates(inputStream chan string, outputStream chan string) {
	var a string
	for b := range inputStream {
		if b != a {
			outputStream <- b
			a = b
		}
	}
	close(outputStream)
}
func main() {
	inputStream := make(chan string)
	outputStream := make(chan string)
	go func() {
		inputStream <- "first"
		inputStream <- "first"
		inputStream <- "second"
		inputStream <- "third"
		inputStream <- "third"
		inputStream <- "third"
		inputStream <- "fourth"
		close(inputStream)
	}()
	go removeDuplicates(inputStream, outputStream)

	for i := range outputStream {
		fmt.Println(i)
	}
}
