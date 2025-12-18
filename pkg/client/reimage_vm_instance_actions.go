// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReimageVmInstance 操作ReimageVmInstance
func (cli *ZSClient) ReimageVmInstance(uuid string, params param.ReimageVmInstanceParam) (*view.ReimageVmInstanceEventView, error) {
	resp := view.ReimageVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

