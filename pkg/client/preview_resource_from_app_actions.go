// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// PreviewResourceFromApp operates on PreviewResourceFromApp
func (cli *ZSClient) PreviewResourceFromApp(params param.PreviewResourceFromAppParam) (*view.PreviewResourceStackView, error) {
	resp := view.PreviewResourceStackView{}
	if err := cli.Post("v1/appcenter/app/preview", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
