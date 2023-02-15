package main

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
	"net/url"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity"
)

func main() {
	fmt.Println("Hello, World!")
	CreateAcsHmacRequest()
}

func CreateAcsHmacRequest() {
	const resourceEndpoint = "https://m365-quickstart.communication.azure.com"
	const pathAndQuery = "/identities?api-version=2021-03-07"
	const resourceSecret = "boh0QPbeVWFa8ZKEMOCTOs/580ZEKXZkO7mVP23HojojSvIuyEM65yI0IvL8Q55h6xBsC2vDuIFB59r3QyUM4Q=="

	// Create a uri you are going to call.
	requestUri, err := url.Parse(resourceEndpoint + pathAndQuery)
	fmt.Println("Request URI : ", requestUri)
	logError(err)

	var scope = azcommunicationidentity.CommunicationIdentityTokenScopeChat
	//Create Body
	body := azcommunicationidentity.CreateRequest{
		CreateTokenWithScopes: []*azcommunicationidentity.CommunicationIdentityTokenScope{&scope},
	}

	serializedBodyBytes, err := json.Marshal(body)
	fmt.Println("Serialized Body : ", string(serializedBodyBytes))
	logError(err)

	// Specify the 'x-ms-date' header as the current UTC timestamp according to the RFC1123 standard
	date := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 GMT")

	// Get the host name corresponding with the 'host' header.
	fmt.Println("Host : ", requestUri.Host)
	host := requestUri.Host

	// Compute a content hash for the 'x-ms-content-sha256' header.
	contentHash := ComputeContentHash(serializedBodyBytes)

	// Prepare a string to sign.
	fmt.Println("Path : ", requestUri.Path+"?"+requestUri.RawQuery)
	//stringToSign := "POST\n" + requestUri.Path + "?" + requestUri.RawQuery + date + ";" + host + ";" + contentHash
	stringToSign := fmt.Sprintf("POST\n%s\n%s;%s;%s", requestUri.Path+"?"+requestUri.RawQuery, date, host, contentHash)

	// Compute the signature.
	signature := ComputeSignature(stringToSign, resourceSecret)

	// Concatenate the string, which will be used in the authorization header.
	authorizationHeader := fmt.Sprintf("HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s", signature)

	requestMessage, err := http.NewRequest(http.MethodPost, requestUri.String(), bytes.NewBuffer(serializedBodyBytes))
	logError(err)

	// Add a content type header.
	requestMessage.Header.Add("Content-Type", "application/json")

	// Add a date header.
	requestMessage.Header.Add("x-ms-date", date)

	// Add a content hash header.
	requestMessage.Header.Add("x-ms-content-sha256", contentHash)

	// Add an authorization header.
	requestMessage.Header.Add("Authorization", authorizationHeader)

	// Create client
	client := &http.Client{}

	// Send request
	response, err := client.Do(requestMessage)
	logError(err)

	// Read response
	responseBody, err := io.ReadAll(response.Body)
	logError(err)

	responseString := string(responseBody)
	fmt.Println(responseString)

	// Close response body
	defer response.Body.Close()
}

func ComputeContentHash(content []byte) string {
	hash := sha256.Sum256(content)
	return base64.StdEncoding.EncodeToString(hash[:])
}

func ComputeSignature(stringToSign string, secret string) string {
	key, _ := base64.StdEncoding.DecodeString(secret)
	h := hmac.New(sha256.New, key)
	h.Write([]byte(stringToSign))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func logError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
