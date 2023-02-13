// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License.

package identity

import (
	"context"
	"fmt"
	"testing"

	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/stretchr/testify/require"
)

func TestCreateClientFromConnectionString(t *testing.T) {
	client := NewCommunicationIdentityClient(
		"",
		runtime.NewPipeline("module", "version", runtime.PipelineOptions{}, nil))
	require.NotNil(t, client)
	identityResponse, err := client.Create(context.TODO(), nil)
	require.Nil(t, err)
	fmt.Println(identityResponse)
}
