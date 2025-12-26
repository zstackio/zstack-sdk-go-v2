// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PreviewResourceStack operates on PreviewResourceStack
func (cli *ZSClient) PreviewResourceStack(params param.PreviewResourceStackParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/cloudformation/stack/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
