// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now() // avoid unused import

// ContainerClusterInventoryView ContainerCluster
type ContainerClusterInventoryView struct {
	BaseInfoView
	BaseTimeView
	ContainerUuid string `json:"containerUuid,omitempty"`
	ClusterId int64 `json:"clusterId,omitempty"`
	ClusterName string `json:"clusterName,omitempty"`
	ContainerName string `json:"containerName,omitempty"`
	ContainerDescription string `json:"containerDescription,omitempty"`
	ProjectTag string `json:"projectTag,omitempty"`
}

