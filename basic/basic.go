package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
	"os"
	"runtime"
	"strings"
	"time"
)

func basic() {
	for {
		fmt.Print("Enter Name: ")
		reader := bufio.NewReader(os.Stdin)
		text, _ := reader.ReadString('\n')
		fmt.Println("Welcome,", text)
		fmt.Print("Continue ? Press C: ")
		choice, _ := reader.ReadString('\n')
		if !strings.EqualFold(strings.TrimSpace(choice), "c") {
			fmt.Println("Good Bye")
			return
		}
	}
}

func types() {
	var ToBe bool = false
	var MaxInt uint64 = 1<<64 - 1
	var z complex128 = cmplx.Sqrt(-1 + 2i)

	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

func conversion() {
	var float float64 = 10.1
	fmt.Println(float, "become", int(float))
}

func forLoop() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

func forever() {
	for {
		fmt.Println("Hey")
	}
}

func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}
	return fmt.Sprint(math.Sqrt(x))
}

func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}

func getSystem() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s.", os)
	}
}

func getDay() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today + 1)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

func getTime() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

func goDefer() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		defer fmt.Println(i)
	}

	fmt.Println("done")
}
func goPointers() {
	i, j := 42, 2701

	p := &i         // point to i
	fmt.Println(*p) // read i through the pointer
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
	fmt.Println(j) // see the new value of j
}

type Vertex struct {
	name     string
	attended bool
}

func maps() {
	var x map[string]string
	x = make(map[string]string)
	x["10"] = "10"
	fmt.Println(x["10"])
}

func main() {
	// maps()
	// basic()
	// types()
	// conversion()
	// forLoop()
	// forever()

	// fmt.Println(
	// 	sqrt(4),
	// 	sqrt(-4),
	// )
	// fmt.Println(
	// 	pow(3, 2, 10),
	// 	pow(3, 3, 20),
	// )

	// getSystem()

	// getDay()

	// getTime()

	// goDefer()

	// goPointers()

	// goStruct()
	// arr()

	// x := recursion(4)
	// fmt.Println(x)
}

func recursion(x int) int {
	if x == 0 {
		return 1
	}
	return x * recursion(x-1)
}
func goStruct() {
	v := Vertex{"Ahmed", false}
	fmt.Println(v.name, v.attended)

	students := []Vertex{{"A", false}, {"B", true}}
	fmt.Println(students)
}
func arr() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println(primes)

	// slice
	primes = [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	// more about slicing
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)
	x := names[0:2]
	y := names[1:3]
	fmt.Println(x, y)
}
