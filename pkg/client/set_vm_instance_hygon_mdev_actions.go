// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// SetVmInstanceHygonMdev operates on SetVmInstanceHygonMdev
func (cli *ZSClient) SetVmInstanceHygonMdev(params param.SetVmInstanceHygonMdevParam) (*view.SetVmInstanceHygonMdevEventView, error) {
	resp := view.SetVmInstanceHygonMdevEventView{}
	if err := cli.Post("v1/vm-instances/{uuid}/hygon-mdev", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
