// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LoginIAM2VirtualID operates on LoginIAM2VirtualID
func (cli *ZSClient) LoginIAM2VirtualID(uuid string, params param.LoginIAM2VirtualIDParam) (*view.LoginIAM2VirtualIDView, error) {
	resp := view.LoginIAM2VirtualIDView{}
	if err := cli.Put("v1/iam2/virtual-ids/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
