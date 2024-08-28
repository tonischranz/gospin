package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	fmt.Println("\033[0;32mhey", "hey 🇨🇭 \u2626 \U0001f1e9\U0001f1ea\U0001f1e8\U0001f1ed", "\033[0m")
	// fmt.Println("\033[0;30m", "du")
	// fmt.Println("\033[0;31m", "du", "\033[2D", "\033[0;32m", "du")
	// fmt.Println("\033[0;33m", "du")
	// fmt.Println("\033[0;34m", "du")
	// fmt.Println("\033[0;35m", "du")
	// fmt.Println("\033[0;36m", "du")
	// fmt.Println("\033[0;37m", "du")

	defer say_bye()

	// go spin()
	// go spin_braille(128)
	go spin_block(256)

	// time.Sleep(16 * time.Second)

	// fs := http.FileServer(http.Dir("/"))

	http.HandleFunc("/", index)
	http.HandleFunc("POST /", do)
	//http.Handle("/img.JPG", fs)
	// http.Handle("/", fs)
	// http.ServeTLS()

	go http.ListenAndServe(":8090", nil)

	time.Sleep(64 * time.Second)
}

func index(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {

		fmt.Fprintln(w, "POST")
	}

	fmt.Fprintln(w, "hello world")
	fmt.Println("get")
}

func do(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, "do something", 'a', req.PostForm.Get("foo"))
}

func spin() {
	fmt.Println()

	defer fmt.Print("\033[1D")
	for {
		// fmt.Print("-", "\033[2D", "\033[0;32m")
		// time.Sleep(64 * time.Millisecond)
		// fmt.Print("/", "\033[2D")
		// time.Sleep(64 * time.Millisecond)
		// fmt.Print("|", "\033[2D", "\033[0;35m")
		// time.Sleep(64 * time.Millisecond)
		// fmt.Print("\\", "\033[2D")
		// time.Sleep(64 * time.Millisecond)

		fmt.Print("\033[0;32m-")
		time.Sleep(64 * time.Millisecond)
		fmt.Print("\033[1D/")
		time.Sleep(64 * time.Millisecond)
		fmt.Print("\033[1D\033[0;37m|")
		time.Sleep(64 * time.Millisecond)
		fmt.Print("\033[1D\\")
		time.Sleep(64 * time.Millisecond)
		fmt.Print("\033[1D")
	}
}

func spin_block(delay int) {
	fmt.Println()
	for {
		fmt.Print("\033[1D\033[0;33m\u2598")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2580")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u259d")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2590")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2597")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2582")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2596")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u258c")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		// fmt.Print("\033[1D\033[0;33m\u2596")
	}
}

func spin_braille(delay int) {
	fmt.Println()
	for {
		fmt.Print("\033[1D\033[0;33m\u2801") //1
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2803") //12
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2806")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2824")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2830")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2818")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		fmt.Print("\033[1D\033[0;33m\u2814")
		time.Sleep(time.Duration(delay) * time.Millisecond)
		// fmt.Print("\033[1D\033[0;33m\u2596")
		// time.Sleep(time.Duration(delay) * time.Millisecond)
		// fmt.Print("\033[1D\033[0;33m\u258c")
		// time.Sleep(time.Duration(delay) * time.Millisecond)
		// fmt.Print("\033[1D\033[0;33m\u2596")
	}
}

func say_bye() {
	fmt.Println()
	fmt.Println("bye")
}
