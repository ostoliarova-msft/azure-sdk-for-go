package azcommunicationidentity

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func sendRequestUsingHmac() {

	const host = "resource_endpoint_name"
	const resourceEndpoint = "https://" + host
	const pathAndQuery = "/identities?api-version=2021-03-07"
	const resourceSecret = "resource_endpoint_secret"

	// Create a uri you are going to call.
	requestUri := resourceEndpoint + pathAndQuery

	body := struct {
		CreateTokenWithScopes []string `json:"createTokenWithScopes"`
	}{
		CreateTokenWithScopes: []string{"chat"},
	}

	serializedBody, err := json.Marshal(body)
	if err != nil {
		fmt.Println(err)
		return
	}

	stringifiedBody := string(serializedBody)

	// Specify the 'x-ms-date' header as the current UTC timestamp according to the RFC1123 standard
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")

	// Compute a content hash for the 'x-ms-content-sha256' header.
	contentHash := ComputeContentHash(stringifiedBody)

	// Prepare a string to sign.
	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", pathAndQuery, date, host, contentHash)

	// Compute the signature.
	signature := ComputeSignature(stringToSign, resourceSecret)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := fmt.Sprintf("HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s", signature)

	request, err := http.NewRequest(http.MethodPost, requestUri, bytes.NewBuffer(serializedBody))
	if err != nil {
		fmt.Println(err)
		return
	}

	// Set headers
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("x-ms-date", date)
	request.Header.Set("x-ms-content-sha256", contentHash)
	request.Header.Set("Authorization", authorizationHeader)

	// Create client
	client := &http.Client{}

	// Send request
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	// Read response
	result, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	stringResult := string(result)

	fmt.Println(stringResult)

	// Close response body
	defer response.Body.Close()

}

func ComputeContentHash(content string) string {
	hash := sha256.Sum256([]byte(content))
	return base64.StdEncoding.EncodeToString(hash[:])
}

func ComputeSignature(stringToSign string, secret string) string {
	key, _ := base64.StdEncoding.DecodeString(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}
