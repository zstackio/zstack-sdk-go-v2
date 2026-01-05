// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ProvisionNfvInstGroup operates on ProvisionNfvInstGroup
func (cli *ZSClient) ProvisionNfvInstGroup(uuid string, params param.ProvisionNfvInstGroupParam) (*view.ProvisionNfvInstGroupEventView, error) {
	resp := view.ProvisionNfvInstGroupEventView{}
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
