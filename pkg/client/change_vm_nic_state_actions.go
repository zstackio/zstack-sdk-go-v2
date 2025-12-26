// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmNicState changes VmNicState
func (cli *ZSClient) ChangeVmNicState(uuid string, params param.ChangeVmNicStateParam) (*view.ChangeVmNicStateEventView, error) {
	resp := view.ChangeVmNicStateEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
