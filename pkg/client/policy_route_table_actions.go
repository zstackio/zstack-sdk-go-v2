// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePolicyRouteTable creates PolicyRouteTable
func (cli *ZSClient) CreatePolicyRouteTable(params param.CreatePolicyRouteTableParam) (*view.PolicyRouteTableInventoryView, error) {
	var resp view.CreatePolicyRouteTableEventView
	if err := cli.Post("v1/policy-routes/tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePolicyRouteTable deletes PolicyRouteTable
func (cli *ZSClient) DeletePolicyRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/tables/{uuid}", uuid, string(deleteMode))
}
// QueryPolicyRouteTable queries PolicyRouteTable list
func (cli *ZSClient) QueryPolicyRouteTable(params *param.QueryParam) ([]view.PolicyRouteTableInventoryView, error) {
	var resp []view.PolicyRouteTableInventoryView
	return resp, cli.List("v1/policy-routes/tables", params, &resp)
}
