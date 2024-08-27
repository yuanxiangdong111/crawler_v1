package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var chan1 chan int

//var chan2 chan int
var wg sync.WaitGroup

func Consume() {
	defer wg.Done()

	for {
		time.Sleep(time.Millisecond * 200)
		select {
		case nums := <-chan1:
			fmt.Printf("consume : %d\n", nums)
		//case <-time.After(time.Second * 3):
		//	fmt.Println("consume timeout")
		default:
			<-time.After(time.Second * 2)
			fmt.Println("consume timeout")
			return
		}
	}
}

func Producer() {
	defer wg.Done()

	for {
		time.Sleep(time.Millisecond * 100)
		rand.Seed(time.Now().UnixNano())
		nums := rand.Intn(10000)
		select {
		case chan1 <- nums:
			fmt.Printf("produce : %d\n", nums)
			//case <-time.After(time.Second * 2):
			//fmt.Println("producer timeout")
		}
	}
}

func main() {
	wg = sync.WaitGroup{}
	//wg.Add(2)
	chan1 = make(chan int, 10)
	//chan2 = make(chan int, 8)

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go Producer()
	}

	for i := 0; i < 2; i++ {
		wg.Add(1)
		go Consume()
	}

	wg.Wait()
	//fmt.Println(chan1)
}
