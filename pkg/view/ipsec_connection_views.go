// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// IPsecConnectionInventoryView IPsecConnection
type IPsecConnectionInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	PeerAddress string `json:"peerAddress,omitempty"`
	AuthMode string `json:"authMode,omitempty"`
	AuthKey string `json:"authKey,omitempty"`
	VipUuid string `json:"vipUuid,omitempty"`
	IkeAuthAlgorithm string `json:"ikeAuthAlgorithm,omitempty"`
	IkeEncryptionAlgorithm string `json:"ikeEncryptionAlgorithm,omitempty"`
	IkeDhGroup int `json:"ikeDhGroup,omitempty"`
	PolicyAuthAlgorithm string `json:"policyAuthAlgorithm,omitempty"`
	PolicyEncryptionAlgorithm string `json:"policyEncryptionAlgorithm,omitempty"`
	Pfs string `json:"pfs,omitempty"`
	PolicyMode string `json:"policyMode,omitempty"`
	TransformProtocol string `json:"transformProtocol,omitempty"`
	IkeVersion string `json:"ikeVersion,omitempty"`
	IdType string `json:"idType,omitempty"`
	LocalId string `json:"localId,omitempty"`
	RemoteId string `json:"remoteId,omitempty"`
	State string `json:"state,omitempty"`
	Status string `json:"status,omitempty"`
	IkeLifeTime int `json:"ikeLifeTime,omitempty"`
	LifeTime int `json:"lifeTime,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	PeerCidrs []IPsecPeerCidrInventoryView `json:"peerCidrs,omitempty"`
	L3NetworkRefs []IPsecL3NetworkRefInventoryView `json:"l3NetworkRefs,omitempty"`
}

