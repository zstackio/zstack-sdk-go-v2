// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PreviewResourceFromApp 操作PreviewResourceFromApp
func (cli *ZSClient) PreviewResourceFromApp(params param.PreviewResourceFromAppParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/appcenter/app/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

