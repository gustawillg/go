package main

import "fmt"

//thread 1
func main() {
	canal := make(chan string) //vazio

	//thread 2
	go func() {
		canal <- "Ola mundo!" //esta cheio
	}()

	//thread 1
	msg := <-canal //canal esvazia
	fmt.Println(msg)
}
