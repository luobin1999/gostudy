package main

import (
	"gostudy/errhandling/filelistingserver/filelisting"
	"log"
	"net/http"
	"os"
)

type userError interface {
	error
	Message() string
}

type appHandler func(writer http.ResponseWriter, request *http.Request) error

func errWrapper(handler appHandler) func(http.ResponseWriter, *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			r := recover()
			if r != nil {
				log.Printf("Painc: %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
				return
			}
		}()

		err := handler(writer, request)

		if err != nil {

			log.Printf("Error occurred handling request: %s", err.Error())

			if err, ok := err.(userError); ok {
				http.Error(writer, err.Message(), http.StatusBadRequest)
				return
			}

			code := http.StatusOK
			switch {
			case os.IsNotExist(err):
				code = http.StatusNotFound
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errWrapper(filelisting.WriteFile))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
