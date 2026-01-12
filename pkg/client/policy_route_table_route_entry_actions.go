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
	var resp view.CreatePolicyRouteTableRouteEntryEventView
	if err := cli.Post("v1/policy-routes/routes", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeletePolicyRouteTableRouteEntry deletes PolicyRouteTableRouteEntry
func (cli *ZSClient) DeletePolicyRouteTableRouteEntry(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/policy-routes/routes", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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
