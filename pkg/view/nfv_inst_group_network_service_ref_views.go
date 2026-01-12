// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// NfvInstGroupNetworkServiceRefInventoryView NfvInstGroupNetworkServiceRef
type NfvInstGroupNetworkServiceRefInventoryView struct {
	BaseInfoView
	BaseTimeView
	Id *int64 `json:"id,omitempty"`
	NfvInstGroupUuid *string `json:"nfvInstGroupUuid,omitempty"`
	NetworkServiceName *string `json:"networkServiceName,omitempty"`
	NetworkServiceUuid *string `json:"networkServiceUuid,omitempty"`
}

