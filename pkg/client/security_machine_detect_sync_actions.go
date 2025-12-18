// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SecurityMachineDetectSync 操作SecurityMachineDetectSync
func (cli *ZSClient) SecurityMachineDetectSync(params param.SecurityMachineDetectSyncParam) (*view.SecurityMachineDetectSyncEventView, error) {
	resp := view.SecurityMachineDetectSyncEventView{}
	if err := cli.Post("v1/security-machine/{uuid}/detect/sync/actions", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

