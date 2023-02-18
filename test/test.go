package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func get_test(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	param := ps.ByName("param")
	time_stamp := ps.ByName("time_stamp")

	fmt.Fprintf(w, "param = \"%s\", time_stamp = \"%s\".", param, time_stamp)
}

func main() {
	router := httprouter.New()

	router.GET("/test/:param", get_test)

	http.ListenAndServe(":12345", router)
}
