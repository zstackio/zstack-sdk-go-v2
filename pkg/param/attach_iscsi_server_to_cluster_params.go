// Copyright (c) ZStack.io, Inc.

package param

// AttachIscsiServerToClusterDetailParam AttachIscsiServerToCluster detail param
type AttachIscsiServerToClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// AttachIscsiServerToClusterParam AttachIscsiServerToCluster request param
type AttachIscsiServerToClusterParam struct {
	BaseParam
	Params AttachIscsiServerToClusterDetailParam `json:"params"`
}
