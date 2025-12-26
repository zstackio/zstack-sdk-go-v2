// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeL2NetworkVlanId changes L2NetworkVlanId
func (cli *ZSClient) ChangeL2NetworkVlanId(uuid string, params param.ChangeL2NetworkVlanIdParam) (*view.ChangeL2NetworkVlanIdEventView, error) {
	resp := view.ChangeL2NetworkVlanIdEventView{}
	if err := cli.Put("v1/l2-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
