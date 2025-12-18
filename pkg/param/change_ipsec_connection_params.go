// Copyright (c) ZStack.io, Inc.

package param

// ChangeIPsecConnectionDetailParam ChangeIPsecConnection detail param
type ChangeIPsecConnectionDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	PeerAddress string `json:"peerAddress" validate:"required"`
	AuthMode string `json:"authMode,omitempty"`
	AuthKey string `json:"authKey" validate:"required"`
	IdType string `json:"idType,omitempty"`
	LocalId string `json:"localId,omitempty"`
	RemoteId string `json:"remoteId,omitempty"`
	IkeVersion string `json:"ikeVersion,omitempty"`
	IkeAuthAlgorithm string `json:"ikeAuthAlgorithm,omitempty"`
	IkeEncryptionAlgorithm string `json:"ikeEncryptionAlgorithm,omitempty"`
	IkeDhGroup int `json:"ikeDhGroup,omitempty"`
	IkeLifeTime int `json:"ikeLifeTime,omitempty"`
	PolicyAuthAlgorithm string `json:"policyAuthAlgorithm,omitempty"`
	PolicyEncryptionAlgorithm string `json:"policyEncryptionAlgorithm,omitempty"`
	Pfs string `json:"pfs,omitempty"`
	PolicyMode string `json:"policyMode,omitempty"`
	TransformProtocol string `json:"transformProtocol,omitempty"`
	LifeTime int `json:"lifeTime,omitempty"`
}

// ChangeIPsecConnectionParam ChangeIPsecConnection request param
type ChangeIPsecConnectionParam struct {
	BaseParam
	Params ChangeIPsecConnectionDetailParam `json:"params"`
}
