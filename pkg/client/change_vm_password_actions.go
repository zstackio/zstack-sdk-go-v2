// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmPassword changes VmPassword
func (cli *ZSClient) ChangeVmPassword(uuid string, params param.ChangeVmPasswordParam) (*view.ChangeVmPasswordEventView, error) {
	resp := view.ChangeVmPasswordEventView{}
	if err := cli.Put("v1/vm-instances/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
