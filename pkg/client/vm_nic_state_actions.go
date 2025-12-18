// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVmNicState 操作VmNicState
func (cli *ZSClient) ChangeVmNicState(uuid string, params param.ChangeVmNicStateParam) (*view.ChangeVmNicStateEventView, error) {
	resp := view.ChangeVmNicStateEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

