// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeVmNicType changes VmNicType
func (cli *ZSClient) ChangeVmNicType(uuid string, params param.ChangeVmNicTypeParam) (*view.ChangeVmNicTypeEventView, error) {
	resp := view.ChangeVmNicTypeEventView{}
	if err := cli.Put("v1/vm-instances/nics/{vmNicUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
