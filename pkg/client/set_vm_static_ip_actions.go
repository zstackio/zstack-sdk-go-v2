// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmStaticIp 操作SetVmStaticIp
func (cli *ZSClient) SetVmStaticIp(uuid string, params param.SetVmStaticIpParam) (*view.SetVmStaticIpEventView, error) {
	resp := view.SetVmStaticIpEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

