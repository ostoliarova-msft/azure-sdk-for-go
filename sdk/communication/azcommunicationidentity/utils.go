//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.

package azcommunicationidentity

import "regexp"

const (
	moduleName = "azcommunicationidentity"
	version    = "0.1.0"
)

const (
	tokenScope = "https://communication.azure.com/.default"
)

func parseConnectionString(connectionString string) map[string]string {
	var rex = regexp.MustCompile("(\\w+)=([^;]*)")
	data := rex.FindAllStringSubmatch(connectionString, -1)

	connectionMap := make(map[string]string)
	for _, kv := range data {
		k := kv[1]
		v := kv[2]
		connectionMap[k] = v
	}
	return connectionMap
}
