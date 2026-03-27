// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAgentVersion queries AgentVersion list
func (cli *ZSClient) QueryAgentVersion(ctx context.Context, params *param.QueryParam) ([]view.AgentVersionInventoryView, error) {
	var resp []view.AgentVersionInventoryView
	return resp, cli.List(ctx, "v1/agent-version", params, &resp)
}

func (cli *ZSClient) GetAgentVersion(ctx context.Context, uuid string) (*view.AgentVersionInventoryView, error) {
	var resp view.AgentVersionInventoryView
	if err := cli.Get(ctx, "v1/agent-version", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageAgentVersion Pagination
func (cli *ZSClient) PageAgentVersion(ctx context.Context, params *param.QueryParam) ([]view.AgentVersionInventoryView, int, error) {
	var agentVersions []view.AgentVersionInventoryView
	total, err := cli.Page(ctx, "v1/agent-version", params, &agentVersions)
	return agentVersions, total, err
}
