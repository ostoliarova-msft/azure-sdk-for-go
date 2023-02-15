package azcommunicationidentity

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
)

type policyFunc func(*policy.Request) (*http.Response, error)

// Do implements the Policy interface on policyFunc.
func (pf policyFunc) Do(req *policy.Request) (*http.Response, error) {
	return pf(req)
}

func CreateCommunicationAccessKeyCredentialPolicy(key string) policy.Policy {
	return policyFunc(func(req *policy.Request) (*http.Response, error) {
		process(req, key)
		resp, err := req.Next()
		return resp, err
	})
}

func process(request *policy.Request, key string) {

	serializedBodyBytes := SerializeRequestBody(request)
	contentHash := ComputeContentHash(serializedBodyBytes)
	AddAuthenticationHeaders(request, contentHash, key)
}

func SerializeRequestBody(request *policy.Request) []byte {

	serializedBodyBytes := []byte{}
	if request.Body() != nil {
		ioReader := request.Body()
		body, _ := io.ReadAll(request.Body())
		serializedBodyBytes = body
		request.SetBody(ioReader, "application/json")
	}
	return serializedBodyBytes
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

func AddAuthenticationHeaders(request *policy.Request, contentHash string, key string) {

	requestUri := request.Raw().URL
	date := time.Now().UTC().Format(http.TimeFormat)
	stringToSign := fmt.Sprintf("%s\n%s\n%s;%s;%s", request.Raw().Method, requestUri.Path+"?"+requestUri.RawQuery, date, requestUri.Host, contentHash)
	signature := ComputeSignature(stringToSign, key)
	authorizationHeader := fmt.Sprintf("HMAC-SHA256 SignedHeaders=x-ms-date;host;x-ms-content-sha256&Signature=%s", signature)

	request.Raw().Header.Add("x-ms-date", date)
	request.Raw().Header.Add("x-ms-content-sha256", contentHash)
	request.Raw().Header.Add("Authorization", authorizationHeader)
}
