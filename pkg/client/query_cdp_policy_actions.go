// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryCdpPolicy queries CdpPolicy list
func (cli *ZSClient) QueryCdpPolicy(params *param.QueryParam) ([]view.CdpPolicyInventoryView, error) {
	var resp []view.CdpPolicyInventoryView
	return resp, cli.List("v1/cdp-backup-storage/policy", params, &resp)
}
