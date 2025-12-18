// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// MigrateVm 操作Vm
func (cli *ZSClient) MigrateVm(uuid string, params param.MigrateVmParam) (*view.MigrateVmEventView, error) {
	resp := view.MigrateVmEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

