// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateNfvInstGroup updates NfvInstGroup
func (cli *ZSClient) UpdateNfvInstGroup(uuid string, params param.UpdateNfvInstGroupParam) (*view.UpdateNfvInstGroupEventView, error) {
	resp := view.UpdateNfvInstGroupEventView{}
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
