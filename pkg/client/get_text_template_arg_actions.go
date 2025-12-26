// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetTextTemplateArg gets TextTemplateArg by uuid
func (cli *ZSClient) GetTextTemplateArg(uuid string) (*view.GetTextTemplateArgView, error) {
	var resp view.GetTextTemplateArgView
	if err := cli.Get("v1/zwatch/textTemplateArg", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
