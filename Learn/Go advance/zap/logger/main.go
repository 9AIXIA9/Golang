package main

import (
	"log"
	"net/http"
)

func simpleHttpGet(url string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Printf("Error fetching url failed %s : %s", url, err.Error())
	} else {
		log.Printf("Status code for %s : %s", url, resp.Status)
		resp.Body.Close()
	}

}
func main() {
	simpleHttpGet("/index")
}
