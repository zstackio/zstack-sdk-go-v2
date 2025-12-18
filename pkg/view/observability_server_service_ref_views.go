// Copyright (c) ZStack.io, Inc.

package view

import "time"

// ObservabilityServerServiceRefInventoryView ObservabilityServerServiceRef
type ObservabilityServerServiceRefInventoryView struct {
	rest string `json:"observabilityServerOfferingUuid,omitempty"`
	rest string `json:"observabilityServerUuid,omitempty"`
	rest string `json:"serviceUuid,omitempty"`
	rest string `json:"serviceType,omitempty"`
	rest string `json:"observabilityServerPublicIp,omitempty"`
	rest string `json:"servicePublicIp,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
}

