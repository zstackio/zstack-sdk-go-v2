// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PauseVmInstance 操作PauseVmInstance
func (cli *ZSClient) PauseVmInstance(uuid string, params param.PauseVmInstanceParam) (*view.PauseVmInstanceEventView, error) {
	resp := view.PauseVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

