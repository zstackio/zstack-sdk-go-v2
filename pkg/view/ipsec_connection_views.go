// Copyright (c) ZStack.io, Inc.

package view

import "time"

// IPsecConnectionInventoryView IPsecConnection
type IPsecConnectionInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"peerAddress,omitempty"`
	rest string `json:"authMode,omitempty"`
	rest string `json:"authKey,omitempty"`
	rest string `json:"vipUuid,omitempty"`
	rest string `json:"ikeAuthAlgorithm,omitempty"`
	rest string `json:"ikeEncryptionAlgorithm,omitempty"`
	rest int `json:"ikeDhGroup,omitempty"`
	rest string `json:"policyAuthAlgorithm,omitempty"`
	rest string `json:"policyEncryptionAlgorithm,omitempty"`
	rest string `json:"pfs,omitempty"`
	rest string `json:"policyMode,omitempty"`
	rest string `json:"transformProtocol,omitempty"`
	rest string `json:"ikeVersion,omitempty"`
	rest string `json:"idType,omitempty"`
	rest string `json:"localId,omitempty"`
	rest string `json:"remoteId,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"status,omitempty"`
	rest int `json:"ikeLifeTime,omitempty"`
	rest int `json:"lifeTime,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest []IPsecPeerCidrInventoryView `json:"peerCidrs,omitempty"`
	rest []IPsecL3NetworkRefInventoryView `json:"l3NetworkRefs,omitempty"`
}

