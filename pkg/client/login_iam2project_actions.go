// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LoginIAM2Project operates on LoginIAM2Project
func (cli *ZSClient) LoginIAM2Project(uuid string, params param.LoginIAM2ProjectParam) (*view.LoginIAM2ProjectView, error) {
	resp := view.LoginIAM2ProjectView{}
	if err := cli.Put("v1/iam2/projects/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
