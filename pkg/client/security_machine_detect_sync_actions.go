// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SecurityMachineDetectSync operates on SecurityMachineDetectSync
func (cli *ZSClient) SecurityMachineDetectSync(params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.Post("v1/security-machine/{uuid}/detect/sync/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
