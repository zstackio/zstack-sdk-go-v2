// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PublishApp 操作PublishApp
func (cli *ZSClient) PublishApp(params param.PublishAppParam) (*view.PublishAppEventView, error) {
	resp := view.PublishAppEventView{}
	if err := cli.Post("v1/appcenter/app", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

