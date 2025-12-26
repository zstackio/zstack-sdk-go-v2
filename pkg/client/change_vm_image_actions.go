// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmImage changes VmImage
func (cli *ZSClient) ChangeVmImage(uuid string, params param.ChangeVmImageParam) (*view.ChangeVmImageEventView, error) {
	resp := view.ChangeVmImageEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
