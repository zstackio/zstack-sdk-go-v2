// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryNfvInstGroup queries NfvInstGroup list
func (cli *ZSClient) QueryNfvInstGroup(params *param.QueryParam) ([]view.NfvInstGroupInventoryView, error) {
	var resp []view.NfvInstGroupInventoryView
	return resp, cli.List("v1/nfvinstgroup/group", params, &resp)
}
