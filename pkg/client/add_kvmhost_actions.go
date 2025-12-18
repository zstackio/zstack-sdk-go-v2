// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddKVMHost 操作AddKVMHost
func (cli *ZSClient) AddKVMHost(params param.AddKVMHostParam) (*view.AddHostEventView, error) {
	resp := view.AddHostEventView{}
	if err := cli.Post("v1/hosts/kvm", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

