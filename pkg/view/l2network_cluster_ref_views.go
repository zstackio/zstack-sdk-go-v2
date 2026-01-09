// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// L2NetworkClusterRefInventoryView L2NetworkClusterRef
type L2NetworkClusterRefInventoryView struct {
	ClusterUuid *string `json:"clusterUuid,omitempty"`
	L2NetworkUuid *string `json:"l2NetworkUuid,omitempty"`
	L2ProviderType *string `json:"l2ProviderType,omitempty"`
	CreateDate *time.Time `json:"createDate,omitempty"`
	LastOpDate *time.Time `json:"lastOpDate,omitempty"`
}

