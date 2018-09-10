package main

import (
	"encoding/json"
	"expvar"
	"github.com/fjos/alphabets/pangram"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strconv"
	"strings"
)

// exported variables for server request metrics
var (
	PangramRequests    = expvar.NewMap("requests")
)

type Response struct {
	Status int    `json:"status"`
	Err    string `json:"error"`
}

type PangramResponse struct {
	Response
	Data PangramResult `json:"data"`
}

type PangramResult struct {
	Pangram  bool             `json:"pangram"`
	Alphabet pangram.Alphabet `json:"alphabet"`
}

type PangramRequest struct {
	Input    string           `json:"input"`
	Alphabet pangram.Alphabet `json:"alphabet"`
}

func CheckIfPangram(alphabet pangram.Alphabet, input io.Reader) PangramResponse {
	var response PangramResponse

	err := alphabet.SetAlphabetContents()
	alphabetString := alphabet.Contents
	alphabet.Contents = alphabetString
	if err != nil {
		response = PangramResponse{
			Data: PangramResult{
				Alphabet: alphabet,
				Pangram:  false,
			},
		}
		response.Status = http.StatusNotFound
		response.Err = err.Error()

		return response
	}

	result, err := pangram.IsPangram(input, alphabetString, 64)
	response = PangramResponse{
		Data: PangramResult{
			Alphabet: alphabet,
			Pangram:  result,
		},
	}

	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Err = err.Error()
	} else {
		response.Status = http.StatusOK

	}

	return response
}

func GETPangramAlphabetInput(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	alphabet := pangram.Alphabet{
		Name: ps.ByName("alphabet"),
	}
	testString := ps.ByName("input")
	log.Printf("Checking if `%s` alphabet is contained in `%s`\n", alphabet.Name, testString)
	response := CheckIfPangram(alphabet, strings.NewReader(testString))

	WriteResponse(w, response)

}

func POSTPangram(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var request PangramRequest
	var response PangramResponse
	err := decoder.Decode(&request)
	if err != nil {
		response = PangramResponse{}
		response.Status = http.StatusBadRequest
		response.Err = err.Error()
		WriteResponse(w, response)
		return
	}
	alphabet := request.Alphabet
	testString := request.Input

	response = CheckIfPangram(alphabet, strings.NewReader(testString))
	response.Status = http.StatusCreated

	WriteResponse(w, response)
}

func WriteResponse(w http.ResponseWriter, response PangramResponse) {
	w.Header().Set("Content-Type", "application/json")

	PangramRequests.Add(strconv.Itoa(response.Status), 1)

	w.WriteHeader(response.Status)
	marshalled, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	w.Write(marshalled)
}

func main() {
	router := httprouter.New()
	router.GET("/api/v1/pangram/:alphabet/:input", GETPangramAlphabetInput)
	router.POST("/api/v1/pangram", POSTPangram)

	router.GET("/debug/vars", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Set("Content-Type", "application/json")
		expvar.Handler().ServeHTTP(w, r)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
