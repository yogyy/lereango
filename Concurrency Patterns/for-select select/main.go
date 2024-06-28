package main

import (
	"fmt"
)

// buffered
func main() {
	charChannel := make(chan string)
	chars := []string{"a", "b", "c"}

	for _, s := range chars {
		select {
		case charChannel <- s:
		}
	}

	close(charChannel)

	for result := range charChannel {
		fmt.Println(result)
	}
}

/* unbuffered */

// func main() {
// 	go func() {
// 		for {
// 			select {
// 			default:
// 				fmt.Println("doing work")
// 			}
// 		}
// 	}()

// 	time.Sleep(time.Second * 10)
// }

/* done channel */

// func doWork(done <-chan bool) {
// 	for {
// 		select {
// 		case <-done:
// 			return
// 		default:
// 			fmt.Println("doing work")
// 		}
// 	}
// }

// func main() {
// 	done := make(chan bool)
// 	go doWork(done)

// 	time.Sleep(time.Second * 3)

// 	close(done)
// }
