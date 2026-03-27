// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePortForwardingRule deletes PortForwardingRule
func (cli *ZSClient) DeletePortForwardingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/port-forwarding", uuid, string(deleteMode))
}
// QueryPortForwardingRule queries PortForwardingRule list
func (cli *ZSClient) QueryPortForwardingRule(ctx context.Context, params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, error) {
	var resp []view.PortForwardingRuleInventoryView
	return resp, cli.List(ctx, "v1/port-forwarding", params, &resp)
}

func (cli *ZSClient) GetPortForwardingRule(ctx context.Context, uuid string) (*view.PortForwardingRuleInventoryView, error) {
	var resp view.PortForwardingRuleInventoryView
	if err := cli.Get(ctx, "v1/port-forwarding", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePortForwardingRule Pagination
func (cli *ZSClient) PagePortForwardingRule(ctx context.Context, params *param.QueryParam) ([]view.PortForwardingRuleInventoryView, int, error) {
	var portForwardingRules []view.PortForwardingRuleInventoryView
	total, err := cli.Page(ctx, "v1/port-forwarding", params, &portForwardingRules)
	return portForwardingRules, total, err
}
// UpdatePortForwardingRule updates PortForwardingRule
func (cli *ZSClient) UpdatePortForwardingRule(ctx context.Context, uuid string, params param.UpdatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/port-forwarding", uuid, "", map[string]interface{}{
		"updatePortForwardingRule": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DetachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) DetachPortForwardingRule(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/port-forwarding", uuid, string(deleteMode))
}
// AttachPortForwardingRule operates on PortForwardingRule
func (cli *ZSClient) AttachPortForwardingRule(ctx context.Context, params param.AttachPortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.Post(ctx, "v1/port-forwarding/{ruleUuid}/vm-instances/nics/{vmNicUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreatePortForwardingRule creates PortForwardingRule
func (cli *ZSClient) CreatePortForwardingRule(ctx context.Context, params param.CreatePortForwardingRuleParam) (*view.PortForwardingRuleInventoryView, error) {
	resp := view.PortForwardingRuleInventoryView{}
	if err := cli.Post(ctx, "v1/port-forwarding", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
