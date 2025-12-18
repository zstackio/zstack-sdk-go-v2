// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCdpPolicy queries CdpPolicy list
func (cli *ZSClient) QueryCdpPolicy(params param.QueryParam) ([]view.CdpPolicyInventoryView, error) {
	var resp []view.CdpPolicyInventoryView
	return resp, cli.List("v1/cdp-backup-storage/policy", &params, &resp)
}
