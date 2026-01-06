// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAgentVersion queries AgentVersion list
func (cli *ZSClient) QueryAgentVersion(params *param.QueryParam) ([]view.AgentVersionInventoryView, error) {
	var resp []view.AgentVersionInventoryView
	return resp, cli.List("v1/agent-version", params, &resp)
}
