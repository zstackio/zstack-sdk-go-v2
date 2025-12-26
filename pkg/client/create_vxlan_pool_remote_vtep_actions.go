// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateVxlanPoolRemoteVtep creates VxlanPoolRemoteVtep
func (cli *ZSClient) CreateVxlanPoolRemoteVtep(params param.CreateVxlanPoolRemoteVtepParam) (*view.CreateVxlanPoolRemoteVtepEventView, error) {
	resp := view.CreateVxlanPoolRemoteVtepEventView{}
	if err := cli.Post("v1/l2-networks/{l2NetworkUuid}/clusters/{clusterUuid}/remote-vtep-ip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
