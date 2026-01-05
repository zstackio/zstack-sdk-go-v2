// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeNfvInstGroupOperationMode changes NfvInstGroupOperationMode
func (cli *ZSClient) ChangeNfvInstGroupOperationMode(uuid string, params param.ChangeNfvInstGroupOperationModeParam) (*view.ChangeNfvInstGroupOperationModeEventView, error) {
	resp := view.ChangeNfvInstGroupOperationModeEventView{}
	if err := cli.Put("v1/nfvinstgroup/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
