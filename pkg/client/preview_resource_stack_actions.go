// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// PreviewResourceStack 操作PreviewResourceStack
func (cli *ZSClient) PreviewResourceStack(params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/cloudformation/stack/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

