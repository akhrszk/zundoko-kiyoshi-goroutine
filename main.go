package main

import "fmt"

const (
	zun     = "ズン"
	doko    = "ドコ"
	kiyoshi = "キ・ヨ・シ！"
)

func ZunDoko(ch chan string, quit chan int) {
	for {
		select {
		case ch <- zun:
		case ch <- doko:
		case <-quit:
			close(ch)
			return
		}
	}
}

func CheckZunDoko(s []string) bool {
	return s[0] == zun && s[1] == zun && s[2] == zun && s[3] == zun && s[4] == doko
}

func main() {
	ch := make(chan string)
	quit := make(chan int)
	go ZunDoko(ch, quit)
	store := make([]string, 5)
	for zundoko := range ch {
		fmt.Print(zundoko, "\n")
		store = append(store[1:], zundoko)
		if CheckZunDoko(store) {
			quit <- 0
		}
	}
	fmt.Print(kiyoshi)
}
