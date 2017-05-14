package main

//TODO: Move routs to another file.
//Reference: https://thenewstack.io/make-a-restful-json-api-go/
//Explains routes and return josn model(struct)

import (
	"encoding/json"
	"log"
	"net/http"
	"os/exec"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/godo/{args}", Index)

	log.Fatal(http.ListenAndServe(":8080", router))
}

func Index(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	commandArg := vars["args"]

	output, err := exec.Command("godoc", commandArg).Output()

	if err == nil {
		outputString := string(output[:])

		var godocOutput GoDocOutput

		if len(outputString) > 0 {
			godocOutput = GoDocOutput{Command: "godoc " + commandArg, Output: outputString}
		} else {
			godocOutput = GoDocOutput{Command: "godoc " + commandArg, Output: "doc for command " + commandArg + " not found."}
		}

		json.NewEncoder(w).Encode(godocOutput)
	}

}
