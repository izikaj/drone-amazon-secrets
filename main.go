package main

import (
	"net/http"
	"strings"

	"github.com/99designs/httpsignatures-go"
)

func main() {
	body := strings.NewReader("{\"path\": \"droneci/global\", \"name\": \"slack_webhook\"}")
	r, _ := http.NewRequest("POST", "http://localhost:3000", body)

	// Sign using the 'Signature' header
	httpsignatures.DefaultSha256Signer.SignRequest("KeyId", "bea26a2221fd8090ea38720fc445eca6", r)
	// OR Sign using the 'Authorization' header
	// httpsignatures.DefaultSha256Signer.AuthRequest("KeyId", "bea26a2221fd8090ea38720fc445eca6", r)

	http.DefaultClient.Do(r)
}
