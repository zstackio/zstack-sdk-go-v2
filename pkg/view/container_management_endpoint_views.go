// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ContainerManagementEndpointInventoryView ContainerManagementEndpoint
type ContainerManagementEndpointInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"accessKeyId,omitempty"`
	rest string `json:"managementIp,omitempty"`
	rest int `json:"managementPort,omitempty"`
	rest string `json:"vendor,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

