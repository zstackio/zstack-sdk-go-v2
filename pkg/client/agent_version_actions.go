// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAgentVersion queries AgentVersion list
func (cli *ZSClient) QueryAgentVersion(params *param.QueryParam) ([]view.AgentVersionInventoryView, error) {
	var resp []view.AgentVersionInventoryView
	return resp, cli.List("v1/agent-version", params, &resp)
}

// PageAgentVersion Pagination
func (cli *ZSClient) PageAgentVersion(params *param.QueryParam) ([]view.AgentVersionInventoryView, int, error) {
	var agentVersions []view.AgentVersionInventoryView
	total, err := cli.Page("v1/agent-version", params, &agentVersions)
	return agentVersions, total, err
}
