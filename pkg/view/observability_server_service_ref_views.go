// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ObservabilityServerServiceRefInventoryView ObservabilityServerServiceRef
type ObservabilityServerServiceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	ObservabilityServerOfferingUuid *string `json:"observabilityServerOfferingUuid,omitempty"`
	ObservabilityServerUuid *string `json:"observabilityServerUuid,omitempty"`
	ServiceUuid *string `json:"serviceUuid,omitempty"`
	ServiceType *string `json:"serviceType,omitempty"`
	ObservabilityServerPublicIp *string `json:"observabilityServerPublicIp,omitempty"`
	ServicePublicIp *string `json:"servicePublicIp,omitempty"`
}

