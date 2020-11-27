package test

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel01(t *testing.T) {
	ch := make(chan int, 100)

	fmt.Println(1)
	go func() {
		fmt.Println(2)

		for {

			time.Sleep(time.Second * time.Duration(1))

			fmt.Println(3)
			ch <- 22
		}
	}()

	fmt.Println(4)

	select {

	case <-ch:
		fmt.Println(5)
	}

	close(ch)

	fmt.Println(ch)
}

func TestChannel02(t *testing.T) {
	c := make(chan int)
	defer close(c)
	go func() { c <- 3 + 4 }()
	i, ok := <-c
	fmt.Println(i, ok)
}

func TestChannel03(t *testing.T) {
	c := make(chan int)
	go func() {
		for i := 0; i < 10; i = i + 1 {
			time.Sleep(time.Duration(1) * time.Second)
			c <- i
		}
		close(c)
	}()
	for i := range c {
		fmt.Println(i)
	}
	fmt.Println("Finished")
}

func TestChannel04(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestChannel05(t *testing.T) {
	ch := make(chan int, 3)
	cl := make(chan bool)
	go func() {
		time.Sleep(10 * time.Second)
		cl <- true
		fmt.Println("关闭->")
	}()

	go func() {
		i := 0
		for {
			time.Sleep(1 * time.Second)
			i++
			fmt.Println("发送->", i, len(ch))
			ch <- i
		}
	}()

	for {
		select {
		case i := <-ch:
			fmt.Println("接收<-", i, len(ch))
			time.Sleep(5 * time.Second)
		case <-cl:
			fmt.Println("关闭<-")
			return
		}
	}
}

func TestChannel06(t *testing.T) {
	ch:= make(chan int, 10)
	ch <- 1
	ch <- 2
	ch <- 3
	ch <- 4
	close(ch)
	for i:= range ch {
		fmt.Println(i)
	}
}

func TestChannel07(t *testing.T) {

	ch := make(chan int, 10)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(1 * time.Second)
			ch <- i
		}
		close(ch)
	}()
	for i := range ch{
		fmt.Println(i)
	}
}

func TestChannel08(t *testing.T) {
	done := make(chan bool, 1)
	go work(done)
	<-done
}

func work(ch chan bool){
	time.Sleep(3 * time.Second)
	ch <- true
}