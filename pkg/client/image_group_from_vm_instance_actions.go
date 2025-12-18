// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateImageGroupFromVmInstance 创建ImageGroupFromVmInstance
func (cli *ZSClient) CreateImageGroupFromVmInstance(params param.CreateImageGroupFromVmInstanceParam) (*view.CreateImageGroupFromVmInstanceEventView, error) {
	resp := view.CreateImageGroupFromVmInstanceEventView{}
	if err := cli.Post("v1/images/groups/from/vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

