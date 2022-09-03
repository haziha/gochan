package gochan

import (
	"fmt"
	"testing"
	"time"
)

func TestGoChan_Len(t *testing.T) {
	c := New[int](-1)
	for i := 0; i < 100; i++ {
		c.Push(i)
	}
	fmt.Println(c.Len())
	for i := 0; i < 50; i++ {
		c.Pop()
	}
	fmt.Println(c.Len())
}

func TestGoChan_Close(t *testing.T) {
	c := New[int](1)
	c.Push(1)
	c.Close()
	//c.Push(0)
	fmt.Println(c.Pop())
	fmt.Println(c.Pop())
}

func TestNew2(t *testing.T) {
	c := New[int](-1)

	go func() {
		for i := 0; i < 10; i++ {
			c.Push(i)
		}
		fmt.Println("go 0 exit")
	}()
	go func() {
		for i := 10; i < 20; i++ {
			c.Push(i)
		}
		fmt.Println("go 1 exit")
	}()
	go func() {
		for i := -10; i < 0; i++ {
			c.Push(i)
		}
		fmt.Println("go 2 exit")
	}()

	for i := 0; i < 30; i++ {
		fmt.Println(c.Pop())
		time.Sleep(time.Second)
	}
}

func TestNew_cap10(t *testing.T) {
	c := New[int](10)

	go func() {
		for i := 0; i < 10; i++ {
			c.Push(i)
		}
		fmt.Println("go 0 exit")
	}()
	go func() {
		for i := 10; i < 20; i++ {
			c.Push(i)
		}
		fmt.Println("go 1 exit")
	}()
	go func() {
		for i := -10; i < 0; i++ {
			c.Push(i)
		}
		fmt.Println("go 2 exit")
	}()

	for i := 0; i < 30; i++ {
		fmt.Println(c.Pop())
		time.Sleep(time.Second)
	}
}

func TestNew_Cap0_2(t *testing.T) {
	c := New[int](0)

	go func() {
		for {
			fmt.Println(c.Pop())
		}
	}()

	go func() {
		for {
			fmt.Println(c.Pop())
		}
	}()

	go func() {
		for {
			fmt.Println(c.Pop())
		}
	}()

	for i := 0; i < 30; i++ {
		c.Push(i)
		time.Sleep(time.Second)
	}
}

func TestNew_Cap0(t *testing.T) {
	c := New[int](0)

	go func() {
		for i := 0; i < 10; i++ {
			c.Push(i)
		}
	}()
	go func() {
		for i := 10; i < 20; i++ {
			c.Push(i)
		}
	}()
	go func() {
		for i := -10; i < 0; i++ {
			c.Push(i)
		}
	}()

	for i := 0; i < 30; i++ {
		fmt.Println(c.Pop())
	}
}

func TestNew(t *testing.T) {
	c := New[int](0)
	go func() {
		fmt.Println(c.Pop())
	}()
	time.Sleep(time.Second)
	c.Push(1)
	time.Sleep(time.Second)
}

func TestNew_Deadlock(t *testing.T) {
	c := New[int](0)
	c.Push(1)
}
