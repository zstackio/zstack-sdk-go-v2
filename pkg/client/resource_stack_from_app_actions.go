// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateResourceStackFromApp 创建ResourceStackFromApp
func (cli *ZSClient) CreateResourceStackFromApp(params param.CreateResourceStackFromAppParam) (*view.CreateResourceStackEventView, error) {
	resp := view.CreateResourceStackEventView{}
	if err := cli.Post("v1/appcenter/app/stack", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

