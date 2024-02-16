package main

/*
Write a web server (port: 3333) with following endpoints:
GET /count - returns value of "count" variable
POST /count (body: {"count": "n"} - increases count value for "n". If number is incorrect, returns 400
GET /hello?name=x - returns text `Hello, "x"!`
*/

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

var count = 0

var port = ":3333"
var baseURL = "http://127.0.0.1" + port

type varArgs map[string]string
type Count struct {
	Count string
}

func handlerCount(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		if _, err := w.Write([]byte(fmt.Sprintf("%d\n", count))); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	case "POST":
		var c Count
		if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		if i, err := strconv.Atoi(c.Count); err == nil {
			count += i
			if _, err := w.Write([]byte("OK")); err != nil {
				http.Error(w, err.Error(), http.StatusBadGateway)
				return
			}
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	default:
		msg := "Sorry, only GET and POST methods are supported."
		if _, err := w.Write([]byte(msg)); err != nil {
			http.Error(w, err.Error(), http.StatusBadGateway)
			return
		}
	}
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if _, err := w.Write([]byte(fmt.Sprintf("Hello, %s!\n", name))); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
}

func createServer() {
	http.HandleFunc("/count", handlerCount)
	http.HandleFunc("/hello", handlerHello)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		fmt.Println("Couldn't run HTTP server:", err)
	}
}

func getRequest(endpoint string, args varArgs) {
	fullURL := baseURL + endpoint
	if len(args) > 0 {
		params := url.Values{}
		for k, v := range args {
			params.Add(k, v)
		}
		fullURL += "?" + params.Encode()
	}

	resp, err := http.Get(fullURL)
	if err != nil {
		panic(err)
	}

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s", data)
}

func postRequest(endpoint string, args varArgs) {
	jsonBytes, err := json.Marshal(args)
	if err != nil {
		log.Fatalln(err)
	}

	resp, err := http.Post(baseURL+endpoint, "application/json", bytes.NewBuffer(jsonBytes))
	if err != nil {
		log.Fatalln(err)
	}

	bytesResp, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return
	}
	fmt.Println(string(bytesResp))
}

func main() {
	go createServer()
	time.Sleep(time.Second)

	var name = "Golang"
	var cnt = 123
	getRequest("/hello", varArgs{"name": name})
	postRequest("/count", varArgs{"count": strconv.Itoa(cnt)})
	getRequest("/count", varArgs{})
	postRequest("/count", varArgs{"count": "asd"}) // Error!
}
