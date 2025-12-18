// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddKVMHostFromConfigFile 操作AddKVMHostFromConfigFile
func (cli *ZSClient) AddKVMHostFromConfigFile(params param.AddKVMHostFromConfigFileParam) (*view.AddHostFromConfigFileEventView, error) {
	resp := view.AddHostFromConfigFileEventView{}
	if err := cli.Post("v1/hosts/kvm/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

