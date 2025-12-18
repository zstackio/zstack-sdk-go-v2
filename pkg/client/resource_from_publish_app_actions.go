// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetResourceFromPublishApp 获取ResourceFromPublishApp详情
func (cli *ZSClient) GetResourceFromPublishApp(uuid string) (*view.GetResourceFromPublishAppView, error) {
	var resp view.GetResourceFromPublishAppView
	if err := cli.Get("v1/appcenter/app/resources", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

