// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SetVmCleanTraffic 操作SetVmCleanTraffic
func (cli *ZSClient) SetVmCleanTraffic(uuid string, params param.SetVmCleanTrafficParam) (*view.SetVmCleanTrafficEventView, error) {
	resp := view.SetVmCleanTrafficEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

