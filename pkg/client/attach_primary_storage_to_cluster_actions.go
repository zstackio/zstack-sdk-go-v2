// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachPrimaryStorageToCluster operates on PrimaryStorageToCluster
func (cli *ZSClient) AttachPrimaryStorageToCluster(params param.AttachPrimaryStorageToClusterParam) (*view.AttachPrimaryStorageToClusterEventView, error) {
	resp := view.AttachPrimaryStorageToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/primary-storage/{primaryStorageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
