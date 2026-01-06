// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePolicyRouteTableRouteEntry creates PolicyRouteTableRouteEntry
func (cli *ZSClient) CreatePolicyRouteTableRouteEntry(params param.CreatePolicyRouteTableRouteEntryParam) (*view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp view.CreatePolicyRouteTableRouteEntryEventView
	if err := cli.Post("v1/policy-routes/routes", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePolicyRouteTableRouteEntry deletes PolicyRouteTableRouteEntry
func (cli *ZSClient) DeletePolicyRouteTableRouteEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/routes/{uuid}", uuid, string(deleteMode))
}
// QueryPolicyRouteTableRouteEntry queries PolicyRouteTableRouteEntry list
func (cli *ZSClient) QueryPolicyRouteTableRouteEntry(params *param.QueryParam) ([]view.PolicyRouteTableRouteEntryInventoryView, error) {
	var resp []view.PolicyRouteTableRouteEntryInventoryView
	return resp, cli.List("v1/policy-routes/routes", params, &resp)
}
