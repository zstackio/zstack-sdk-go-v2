// Copyright (c) ZStack.io, Inc.

package param

// DetachIscsiServerFromClusterDetailParam DetachIscsiServerFromCluster detail param
type DetachIscsiServerFromClusterDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	ClusterUuid string `json:"clusterUuid" validate:"required"`
}

// DetachIscsiServerFromClusterParam DetachIscsiServerFromCluster request param
type DetachIscsiServerFromClusterParam struct {
	BaseParam
	Params DetachIscsiServerFromClusterDetailParam `json:"params"`
}
