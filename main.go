package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/99designs/httpsignatures-go"
)

func main() {
	body := strings.NewReader("{\"path\": \"droneci/global\", \"name\": \"slack_webhook\"}")
	r, err := http.NewRequest("POST", "http://localhost:3000", body)

	// Sign using the 'Signature' header
	httpsignatures.DefaultSha256Signer.SignRequest("KeyId", "bea26a2221fd8090ea38720fc445eca6", r)
	// OR Sign using the 'Authorization' header
	// httpsignatures.DefaultSha256Signer.AuthRequest("KeyId", "bea26a2221fd8090ea38720fc445eca6", r)

	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	respBody, _ := ioutil.ReadAll(resp.Body)

	fmt.Printf("RESPONSE: %v \n", string(respBody))
}
