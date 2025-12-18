// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LoginIAM2Project 操作LoginIAM2Project
func (cli *ZSClient) LoginIAM2Project(uuid string, params param.LoginIAM2ProjectParam) (*view.LoginIAM2ProjectView, error) {
	resp := view.LoginIAM2ProjectView{}
	if err := cli.Put("v1/iam2/projects/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

