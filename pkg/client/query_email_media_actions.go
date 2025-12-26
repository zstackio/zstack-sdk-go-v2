// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryEmailMedia queries EmailMedia list
func (cli *ZSClient) QueryEmailMedia(params *param.QueryParam) ([]view.MediaInventoryView, error) {
	var resp []view.MediaInventoryView
	return resp, cli.List("v1/media/emails", params, &resp)
}
