package main

import "fmt"

type Retiever interface {
	Get(url string) string
}

func download(r Retiever) string {
	return r.Get("www.baidu.com")
}

func main() {
	var r Retiever
	fmt.Println(download(r))
}
