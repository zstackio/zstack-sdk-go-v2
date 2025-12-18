// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LogInByAccount 操作LogInByAccount
func (cli *ZSClient) LogInByAccount(uuid string, params param.LogInByAccountParam) (*view.LogInView, error) {
	resp := view.LogInView{}
	if err := cli.Put("v1/accounts/login", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

