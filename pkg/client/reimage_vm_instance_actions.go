// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReimageVmInstance operates on ReimageVmInstance
func (cli *ZSClient) ReimageVmInstance(uuid string, params param.ReimageVmInstanceParam) (*view.ReimageVmInstanceEventView, error) {
	resp := view.ReimageVmInstanceEventView{}
	if err := cli.Put("v1/vm-instances/{vmInstanceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
