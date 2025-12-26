// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AddKVMHostFromConfigFile adds KVMHostFromConfigFile
func (cli *ZSClient) AddKVMHostFromConfigFile(params param.AddKVMHostFromConfigFileParam) (*view.AddHostFromConfigFileEventView, error) {
	resp := view.AddHostFromConfigFileEventView{}
	if err := cli.Post("v1/hosts/kvm/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
