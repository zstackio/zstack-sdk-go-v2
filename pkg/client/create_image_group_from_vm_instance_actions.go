// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// CreateImageGroupFromVmInstance creates ImageGroupFromVmInstance
func (cli *ZSClient) CreateImageGroupFromVmInstance(params param.CreateImageGroupFromVmInstanceParam) (*view.CreateImageGroupFromVmInstanceEventView, error) {
	resp := view.CreateImageGroupFromVmInstanceEventView{}
	if err := cli.Post("v1/images/groups/from/vm-instance", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
