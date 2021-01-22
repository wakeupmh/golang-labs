package main

import (
	"fmt"
	"os"
	"log"
	http "net/http"
	"github.com/google/uuid"
)

func hostNameHandler(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.NewUUID()

	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	log.Printf("[INFO] ReqID: %v", id.String())

	myHostName, _ := os.Hostname()
	fmt.Fprintln(w, "Hostname: ", myHostName)
}

func main() {

	http.HandleFunc("/", hostNameHandler)

	log.Printf("[INFO] Listening on port :%d", 8080)

	http.ListenAndServe(fmt.Sprintf(":%d", 8080), nil)

}