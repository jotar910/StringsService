package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"

	"joaor.dev.com/joao/stringservice/pkg/parser"
)

// HandleParse func handles all the requests about parsing a string
func HandleParse(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		bodyR := r.Body
		defer bodyR.Close()
		body, err := ioutil.ReadAll(bodyR)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
		var p Parser
		dec := json.NewDecoder(strings.NewReader(string(body)))
		err = dec.Decode(&p)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("4xx - Something bad happened!"))
		}
		res, err := json.Marshal(&ParserRes{parser.Parse(p.Text, strings.Join(p.Flags, ","))})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))
		}
		w.Write(res)
	case http.MethodOptions:
		setupRequest(&w)
		return
	default:
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 - Not found!"))
	}
}

func setupRequest(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
