// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryAudit queries Audit list
func (cli *ZSClient) QueryAudit(params *param.QueryParam) ([]view.AuditsInventoryView, error) {
	var resp []view.AuditsInventoryView
	return resp, cli.List("v1/zwatch/audit-records", params, &resp)
}
