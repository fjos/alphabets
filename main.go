package main

import (
	"github.com/fjos/alphabets/model"
	"github.com/julienschmidt/httprouter"
	"log"
	"expvar"
	"net/http"

)

func main() {
	router := httprouter.New()
	router.GET("/api/v1/pangram/:alphabet/:input", model.GETPangramAlphabetInput)
	router.POST("/api/v1/pangram", model.POSTPangram)

	router.GET("/debug/vars", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		expvar.Handler().ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
