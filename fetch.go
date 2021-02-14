package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

func makeURLList(data string) []string {
	// Consider regex
	data = strings.ReplaceAll(data, "[", "")
	data = strings.ReplaceAll(data, "]", "")
	data = strings.ReplaceAll(data, "\"", "")

	// Split file
	arr := strings.Split(data, ",")

	return arr
}

func controlFolder(folder string) {
	_, err := os.Open(folder)
	if err != nil {
		if os.IsNotExist(err) {
			os.Mkdir(folder, 0755)
		}
	}
}

func makeURL(tag string) string {
	return "https://www.rcsb.org/fasta/entry/" + tag
}

func extract(url string, channel chan string, finished chan bool) {
	resp, err := http.Get(url)

	defer func() {
		finished <- true
	}()

	if err != nil {
		fmt.Println(err)
		return
	}

	data, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()

	// Check for errors in data
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	channel <- string(data)
}

func main() {
	// Request `entry identities`
	resp, err := http.Get("https://data.rcsb.org/rest/v1/holdings/current/entry_ids")

	// Check for request errors
	if err != nil {
		print(err)
		return
	}

	// Get requested data
	data, err := ioutil.ReadAll(resp.Body)

	// Check for errors in data
	if err != nil {
		fmt.Println("File reading error", err)
		return
	}

	// Close `entry identities`
	resp.Body.Close()

	// Clean file
	arr := makeURLList(string(data))

	folder := "fasta/"

	// Ensure folder exist
	controlFolder(folder)

	// Channels
	URLChannel := make(chan string)
	EOLChannel := make(chan bool)

	// Kick off the extraction process (concurrently)
	for _, tag := range arr {
		go extract(makeURL(tag), URLChannel, EOLChannel)
	}

	for _, tag := range arr {
		file, _ := os.Create(folder + tag + ".fasta")

		_, _ = io.WriteString(file, <-URLChannel)

		file.Close()
	}

	close(URLChannel)
	close(EOLChannel)
}
