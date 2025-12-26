// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ChangeVmNicType changes VmNicType
func (cli *ZSClient) ChangeVmNicType(uuid string, params param.ChangeVmNicTypeParam) (*view.ChangeVmNicTypeEventView, error) {
	resp := view.ChangeVmNicTypeEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
