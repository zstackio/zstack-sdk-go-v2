// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SyncNfvInstGroup operates on SyncNfvInstGroup
func (cli *ZSClient) SyncNfvInstGroup(uuid string, params param.SyncNfvInstGroupParam) (*view.SyncNfvInstGroupEventView, error) {
	resp := view.SyncNfvInstGroupEventView{}
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
