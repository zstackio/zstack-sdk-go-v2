// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmSshKey 操作SetVmSshKey
func (cli *ZSClient) SetVmSshKey(uuid string, params param.SetVmSshKeyParam) (*view.SetVmSshKeyEventView, error) {
	resp := view.SetVmSshKeyEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

