package multiThread

import (
	"fmt"
	"net/http"
	"time"
)


func Multi(w http.ResponseWriter, r *http.Request) {

	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("Hello")
	}
}

// gorutine思い出すために用意しておく
func Hoge(w http.ResponseWriter, r *http.Request) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println("World")
	}
}