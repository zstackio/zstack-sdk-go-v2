// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetTextTemplateArg 获取TextTemplateArg详情
func (cli *ZSClient) GetTextTemplateArg(uuid string) (*view.GetTextTemplateArgView, error) {
	var resp view.GetTextTemplateArgView
	if err := cli.Get("v1/zwatch/textTemplateArg", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

