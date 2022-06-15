package main

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"sync"
)

func main() {
	//problem1
	colorSlice()

	//problem2
	readChanelValues()

	//problem3
	mutexRaceCondition()

	//problem4
	chanelRaceCondition()

	//problem5 is in main_test file

	//problem6
	httpFunction()
}

func colorSlice() {
	c1 := []string{"Red", "Black", "White"}
	c2 := []string{"Black", "Yellow", "Orange"}

	c1 = append(c1, c2...)

	colorMap := make(map[string]int)

	for i, v := range c1 {
		colorMap[v] = i
	}

	fmt.Println(reflect.ValueOf(colorMap).MapKeys())
}
func readChanelValues() {
	ch := make(chan int, 10)
	values := []int{10, 20, 35, 100, 200, 502}

	go func() {
		for i := 0; i < len(values); i++ {
			ch <- values[i]
		}
		close(ch)
	}()

	fmt.Println("Values reading from channel :")
	for v := range ch {
		fmt.Println(v)
	}
}

type MutexStruct struct {
	xLock sync.Mutex
	x     int
}

func mutexRaceCondition() {

	var w sync.WaitGroup
	ms := MutexStruct{
		x: int(0),
	}
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go ms.increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", ms.x)
}

func (ms *MutexStruct) increment(wg *sync.WaitGroup) {
	ms.xLock.Lock()
	defer ms.xLock.Unlock()
	ms.x = ms.x + 1
	wg.Done()
}

var x = 0

func chanelRaceCondition() {

	ch := make(chan struct{})
	for i := 0; i < 1000; i++ {
		go increment(ch)
		<-ch
	}
	fmt.Println("final value of x", x)
}

func increment(ch chan struct{}) {
	x = x + 1
	ch <- struct{}{}
}

func Subtract(x, y int) (res int) {
	return x - y
}
func Add(x, y int) (res int) {
	return x + y
}

func httpFunction() {

	log.Print("Listening...")

	handleResponse := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
	http.Handle("/", LogStatusInfo(handleResponse))
	http.ListenAndServe(":8000", nil)

}

func LogStatusInfo(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		recorder := &ResponseStatus{
			ResponseWriter: w,
			Status:         http.StatusServiceUnavailable,
		}
		h.ServeHTTP(recorder, r)
		fmt.Fprintf(w, "status %d : %s ", recorder.Status, http.StatusText(recorder.Status))
	})
}

type ResponseStatus struct {
	http.ResponseWriter
	Status int
}
