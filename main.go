package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"time"
	"unsafe"
)

type Os int

const (
	Mac Os = iota + 1
	Windows
	Linux
)

type Task struct {
	Title    string
	Estimate int
}

type controller interface {
	speedUp() int
	speedDown() int
}
type vehicle struct {
	speed       int
	enginePower int
}
type bycycle struct {
	speed      int
	humanPower int
}

func (v *vehicle) speedUp() int {
	v.speed += 10 * v.enginePower
	return v.speed
}
func (v *vehicle) speedDown() int {
	v.speed -= 5 * v.enginePower
	return v.speed
}
func (b *bycycle) speedUp() int {
	b.speed += 3 * b.humanPower
	return b.speed
}
func (b *bycycle) speedDown() int {
	b.speed -= 1 * b.humanPower
	return b.speed
}
func speedUpAndDown(c controller) {
	fmt.Printf("current speed: %v\n", c.speedUp())
	fmt.Printf("current speed: %v\n", c.speedDown())
}
func (v vehicle) String() string {
	return fmt.Sprintf("Vehicle current speed is %v (enginePower %v)", v.speed, v.enginePower)
}

func Add(x, y int) int {
	return x + y
}
func Divide(x, y int) float32 {
	if y == 0 {
		return 0.
	}
	return float32(x) / float32(y)
}

func main() {
	file, err := os.Create("log.txt")
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()
	flags := log.Lshortfile
	warnLogger := log.New(io.MultiWriter(file, os.Stderr), "WARN: ", flags)

	warnLogger.Println("warning A")

	errorLogger := log.New(io.MultiWriter(file, os.Stderr), "ERROR: ", flags)

	errorLogger.Fatalln("critical error")

	err1 := errors.New("something wrong")
	fmt.Printf("%[1]p %[1]T %[1]v\n", err1)
	fmt.Printf(err1.Error())

	var n int
	for {
		if n > 3 {
			break
		}
		fmt.Println(n)
		n += 1
		time.Sleep(300 * time.Millisecond)
	}
	v := &vehicle{0, 5}
	speedUpAndDown(v)
	fmt.Println(v)
	task1 := Task{
		Title:    "Learn Go",
		Estimate: 3,
	}
	task1.Title = "Learning Go"
	task1.extendEstimate()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)
	task1.extendEstimatePointer()
	fmt.Printf("task1 value receiver: %+v\n", task1.Estimate)

	// fmt.Println("hello world")
	// sl := []int{1, 2, 3}
	// if len(sl) > 0 {
	// 	fmt.Println("unreachable code")
	// }
	// godotenv.Load()
	// fmt.Println(os.Getenv("GO_ENV"))
	// fmt.Println(calculator.Offset)
	// fmt.Println(calculator.Sum(1.0, 2.0))
	// fmt.Println(calculator.Multiply(1, 2))
	// var i = 2
	i := 1
	ui := uint16(2)

	fmt.Println(i)
	fmt.Printf("i: %[1]v %[1]T ui: %[2]v %[2]T\n", i, ui)

	fmt.Printf("Mac: %v Windows: %v Linux: %v\n", Mac, Windows, Linux)

	var ui1 uint16
	fmt.Printf("memory address of ui1: %p\n", &ui1)

	var p1 *uint16
	fmt.Printf("value of p1: %v\n", p1)
	p1 = &ui1
	fmt.Printf("value of p1: %v\n", p1)

	fmt.Printf("size of p1: %d[bytes]\n", unsafe.Sizeof(p1))
	fmt.Printf("memory address of p1: %p\n", &p1)
	fmt.Printf("value of ui1(dereference): %v\n", *p1)

	ok, result := true, "A"
	if ok {
		result := "B"
		println(result)
	} else {
		result := "C"
		println(result)
	}
	println(result)

	var s1 []int
	s2 := []int{1, 3, 30}
	s1 = append(s1, 1, 3, 19)
	fmt.Printf("s1: %[1]T %[1]v %v %v\n", s1, len(s1), cap(s1))
	fmt.Printf("s2: %[1]T %[1]v %v %v\n", s2, len(s2), cap(s2))

	s4 := make([]int, 0, 2)
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))
	s4 = append(s4, 1, 23, 4)
	fmt.Printf("s4: %[1]T %[1]v %v %v\n", s4, len(s4), cap(s4))

	var m1 map[string]int
	m2 := map[string]int{}
	fmt.Printf("%v %v \n", m1, m1 == nil)
	fmt.Printf("%v %v \n", m2, m2 == nil)
	m2["A"] = 10
	m2["B"] = 20
	m2["C"] = 30

	for k, v := range m2 {
		fmt.Printf("%v %v\n", k, v)
	}

	func(i int) {
		fmt.Println(i)
	}(i)

	f1 := func(i int) int {
		return i + 1
	}
	fmt.Println(f1(i))
	f3 := multiply()
	fmt.Println(f3(2))
}
func (task Task) extendEstimate() {
	task.Estimate += 10
}
func (taskp *Task) extendEstimatePointer() {
	taskp.Estimate += 10
}
func multiply() func(int) int {
	return func(n int) int {
		return n * 1000
	}
}
