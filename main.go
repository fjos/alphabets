package main

import (
	"encoding/json"
	"github.com/fjos/alphabets/pangram"
	"github.com/julienschmidt/httprouter"
	"io"
	"log"
	"net/http"
	"strings"
)

type Response struct {
	Status int             `json:"status"`
	Err    string          `json:"error"`
	Data   PangramResponse `json:"data"`
}

type PangramResponse struct {
	Pangram  bool   `json:"pangram"`
	Alphabet string `json:"alphabet"`
}

type PangramRequest struct {
	Input    string `json:"input"`
	Alphabet string `json:"alphabet"`
}

func CheckIfPangram(alphabetName string, input io.Reader) Response {
	var response Response

	alphabet, err := pangram.GetAlphabet(alphabetName)
	if err != nil {
		response = Response{
			Status: http.StatusBadRequest,
			Err:    err.Error(),
			Data: PangramResponse{
				Alphabet: alphabetName,
				Pangram:  false,
			},
		}
		return response
	}

	result, err := pangram.IsPangram(input, alphabet, 64)
	response = Response{
		Status: http.StatusOK,
		Data: PangramResponse{
			Alphabet: alphabetName,
			Pangram:  result,
		},
	}
	if err != nil {
		response.Status = http.StatusInternalServerError
		response.Err = err.Error()
	}

	return response
}

func SelectAlphabet(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	alphabet := ps.ByName("alphabet")
	testString := ps.ByName("input")
	log.Printf("Checking if `%s` alphabet is contained in `%s`\n", alphabet, testString)

	response := CheckIfPangram(alphabet, strings.NewReader(testString))
	w.WriteHeader(response.Status)

	marshalled, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}

	w.Write(marshalled)

}

func RequestViaJson(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	decoder := json.NewDecoder(r.Body)
	var request PangramRequest
	var response Response
	err := decoder.Decode(&request)
	if err != nil {
		response = Response{
			Status: http.StatusBadRequest,
			Err:    err.Error(),
		}
		w.WriteHeader(response.Status)
		marshalled, err := json.Marshal(response)
		if err != nil {
			log.Println(err)
		}

		w.Write(marshalled)
		return
	}
	alphabet := request.Alphabet
	testString := request.Input

	response = CheckIfPangram(alphabet, strings.NewReader(testString))
	response.Status = http.StatusCreated

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.Status)

	marshalled, err := json.Marshal(response)
	if err != nil {
		log.Println(err)
	}
	w.Write(marshalled)
}

func main() {
	router := httprouter.New()
	// router.GET("/api/v1/pangram/:input", DefaultLatin)
	router.GET("/api/v1/pangram/:alphabet/:input", SelectAlphabet)
	router.POST("/api/v1/pangram", RequestViaJson)

	log.Fatal(http.ListenAndServe(":8080", router))
}
