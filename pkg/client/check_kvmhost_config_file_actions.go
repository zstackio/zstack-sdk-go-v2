// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckKVMHostConfigFile operates on CheckKVMHostConfigFile
func (cli *ZSClient) CheckKVMHostConfigFile(params param.CheckKVMHostConfigFileParam) (*view.CheckHostConfigFileView, error) {
	resp := view.CheckHostConfigFileView{}
	if err := cli.Post("v1/hosts/kvm/from-file/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
