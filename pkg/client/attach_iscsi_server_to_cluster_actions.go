// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachIscsiServerToCluster operates on IscsiServerToCluster
func (cli *ZSClient) AttachIscsiServerToCluster(params param.AttachIscsiServerToClusterParam) (*view.AttachIscsiServerToClusterEventView, error) {
	resp := view.AttachIscsiServerToClusterEventView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/storage-devices/iscsi/servers/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
