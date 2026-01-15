// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePolicyRouteTableRouteEntry creates PolicyRouteTableRouteEntry
func (cli *ZSClient) CreatePolicyRouteTableRouteEntry(params param.CreatePolicyRouteTableRouteEntryParam) (*view.PolicyRouteTableRouteEntryInventoryView, error) {
	resp := view.PolicyRouteTableRouteEntryInventoryView{}
	if err := cli.Post("v1/policy-routes/routes", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteTableRouteEntry deletes PolicyRouteTableRouteEntry
func (cli *ZSClient) DeletePolicyRouteTableRouteEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/routes", uuid, string(deleteMode))
}
// QueryPolicyRouteTableRouteEntry queries PolicyRouteTableRouteEntry list
func (cli *ZSClient) QueryPolicyRouteTableRouteEntry(params *param.QueryParam) ([]view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp []view.PolicyRouteTableRouteEntryInventoryView
	return resp, cli.List("v1/policy-routes/routes", params, &resp)
}

func (cli *ZSClient) GetPolicyRouteTableRouteEntry(uuid string) (*view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp view.PolicyRouteTableRouteEntryInventoryView
	if err := cli.Get("v1/policy-routes/routes", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PagePolicyRouteTableRouteEntry Pagination
func (cli *ZSClient) PagePolicyRouteTableRouteEntry(params *param.QueryParam) ([]view.PolicyRouteTableRouteEntryInventoryView, int, error) {
	var policyRouteTableRouteEntries []view.PolicyRouteTableRouteEntryInventoryView
	total, err := cli.Page("v1/policy-routes/routes", params, &policyRouteTableRouteEntries)
	return policyRouteTableRouteEntries, total, err
}
