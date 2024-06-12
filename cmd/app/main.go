package main

import "net/http"

func main() {
	app, err := initialize()
	if err != nil {
		panic(err)
	}

	if err := http.ListenAndServe(":8080", app.handler); err != nil {
		panic(err)
	}

}
