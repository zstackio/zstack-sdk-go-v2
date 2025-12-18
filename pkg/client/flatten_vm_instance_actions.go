// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// FlattenVmInstance 操作FlattenVmInstance
func (cli *ZSClient) FlattenVmInstance(uuid string, params param.FlattenVmInstanceParam) (*view.FlattenVmInstanceEventView, error) {
	resp := view.FlattenVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

