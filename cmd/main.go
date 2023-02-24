package main

import (
	"fmt"

	"github.com/aldoclg/webcrawler/bfs"
	"github.com/aldoclg/webcrawler/queue"
)

func main() {
	fmt.Println("Webcrawler world")

	bfs := bfs.NewBFS(queue.NewQueue[string](), make(map[string]bool, 0))

	bfs.Traverse("https://gobyexample.com/regular-expressions")
}
