## Go

### Generation

```ps
cd <azcommunicationidentity>
autorest autorest.md
```

```yaml
clear-output-folder: false
export-clients: true
go: true
input-file: https://raw.githubusercontent.com/Azure/azure-rest-api-specs/main/specification/communication/data-plane/Identity/stable/2022-10-01/CommunicationIdentity.json
license-header: MICROSOFT_MIT_NO_VERSION
module: github.com/Azure/azure-sdk-for-go/sdk/communication/azcommunicationidentity
openapi-type: "data-plane"
output-folder: ../azcommunicationidentity
override-client-name: Client
security: "AADToken"
security-scopes: "https://vault.azure.net/.default"
use: "@autorest/go@4.0.0-preview.43"
version: "v1.0.0-beta"

directive:
 # delete generated constructor
  - from: client.go
    where: $
    transform: return $.replace(/(?:\/\/.*\s)+func NewClient.+\{\s(?:.+\s)+\}\s/, "");
```
