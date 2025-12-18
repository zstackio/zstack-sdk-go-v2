// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ListVMsFromKVMHost 操作ListVMsFromKVMHost
func (cli *ZSClient) ListVMsFromKVMHost(params param.ListVMsFromKVMHostParam) (*view.ListVMsFromKVMHostEventView, error) {
	resp := view.ListVMsFromKVMHostEventView{}
	if err := cli.Post("v1/v2v", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

