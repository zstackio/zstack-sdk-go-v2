// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ListVMsFromKVMHost operates on ListVMsFromKVMHost
func (cli *ZSClient) ListVMsFromKVMHost(params param.ListVMsFromKVMHostParam) (*view.ListVMsFromKVMHostEventView, error) {
	resp := view.ListVMsFromKVMHostEventView{}
	if err := cli.Post("v1/v2v", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
