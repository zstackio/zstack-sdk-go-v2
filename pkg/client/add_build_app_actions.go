// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AddBuildApp 操作AddBuildApp
func (cli *ZSClient) AddBuildApp(params param.AddBuildAppParam) (*view.AddBuildAppEventView, error) {
	resp := view.AddBuildAppEventView{}
	if err := cli.Post("v1/appcenter/buildapp/add", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

