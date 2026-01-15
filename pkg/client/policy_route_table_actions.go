// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreatePolicyRouteTable creates PolicyRouteTable
func (cli *ZSClient) CreatePolicyRouteTable(params param.CreatePolicyRouteTableParam) (*view.PolicyRouteTableInventoryView, error) {
	resp := view.PolicyRouteTableInventoryView{}
	if err := cli.Post("v1/policy-routes/tables", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeletePolicyRouteTable deletes PolicyRouteTable
func (cli *ZSClient) DeletePolicyRouteTable(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/policy-routes/tables", uuid, string(deleteMode))
}
// QueryPolicyRouteTable queries PolicyRouteTable list
func (cli *ZSClient) QueryPolicyRouteTable(params *param.QueryParam) ([]view.PolicyRouteTableInventoryView, error) {
	var resp []view.PolicyRouteTableInventoryView
	return resp, cli.List("v1/policy-routes/tables", params, &resp)
}

// PagePolicyRouteTable Pagination
func (cli *ZSClient) PagePolicyRouteTable(params *param.QueryParam) ([]view.PolicyRouteTableInventoryView, int, error) {
	var policyRouteTables []view.PolicyRouteTableInventoryView
	total, err := cli.Page("v1/policy-routes/tables", params, &policyRouteTables)
	return policyRouteTables, total, err
}
