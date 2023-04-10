package main

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

func hello(w http.ResponseWriter, req *http.Request) {
	logrus.Info("hello from function 1 ####------- ")
	fmt.Fprintf(w, "hello from poonams first website....\n")
}

func hello2(w http.ResponseWriter, req *http.Request) {
	logrus.Info("hello from function 2  ##### ")
	fmt.Fprintf(w, "hello again from poonams first website....\n")
}

func headers(w http.ResponseWriter, req *http.Request) {
	logrus.Info("hello from function headers..... ")
	for name, headers := range req.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/hello2", hello2)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8090", nil)
}
