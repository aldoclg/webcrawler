package bfs

import (
	"bytes"
	"crypto/tls"
	"io"
	"log"
	"net/http"
	"regexp"

	"github.com/aldoclg/webcrawler/queue"
)

type BreathFirstSearch struct {
	queue               queue.Queue[string]
	discorveredWebsites map[string]bool
}

const urlPattern string = "https://(\\w+\\.)*(\\w+)"

var client *http.Client

func NewBFS(queue queue.Queue[string], discorveredWebsites map[string]bool) BreathFirstSearch {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	client = &http.Client{Transport: tr}
	return BreathFirstSearch{queue: queue, discorveredWebsites: discorveredWebsites}
}

func (bfs *BreathFirstSearch) Traverse(root string) {
	bfs.queue.Enqueue(root)
	bfs.discorveredWebsites[root] = true

	for bfs.queue.IsNotEmpty() {
		actual := bfs.queue.Dequeue()
		log.Println("Dequeue", actual)
		rawHTML := readURL(actual)

		r, _ := regexp.Compile(urlPattern)
		newUrls := r.FindAllString(rawHTML, 5)

		for _, newUrl := range newUrls {
			if _, ok := bfs.discorveredWebsites[newUrl]; !ok {
				bfs.discorveredWebsites[newUrl] = true
				log.Printf("New website found: %s", newUrl)
				bfs.queue.Enqueue(newUrl)
			}
		}
	}
}

func readURL(url string) string {

	resp, err := client.Get(url)
	if err != nil {
		log.Println("Web site search error", url, err)
	}
	if resp == nil {
		log.Println("Response nil", resp)
		return ""
	}
	defer resp.Body.Close()

	log.Println("Reading response")
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		log.Println("Read response error")
		log.Fatal(err)
	}
	var buffer bytes.Buffer
	for _, c := range body {
		buffer.WriteByte(c)
	}
	return buffer.String()
}
