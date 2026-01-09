// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySlbGroup queries SlbGroup list
func (cli *ZSClient) QuerySlbGroup(params *param.QueryParam) ([]view.SlbGroupInventoryView, error) {
	var resp []view.SlbGroupInventoryView
	return resp, cli.List("v1/load-balancers/slb/groups", params, &resp)
}

func (cli *ZSClient) GetSlbGroup(uuid string) (*view.SlbGroupInventoryView, error) {
	var resp view.SlbGroupInventoryView
	if err := cli.Get("v1/load-balancers/slb/groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSlbGroup creates SlbGroup
func (cli *ZSClient) CreateSlbGroup(params param.CreateSlbGroupParam) (*view.SlbGroupInventoryView, error) {
	var resp view.CreateSlbGroupEventView
	if err := cli.Post("v1/load-balancers/slb/groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSlbGroup deletes SlbGroup
func (cli *ZSClient) DeleteSlbGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/load-balancers/slb/group", uuid, string(deleteMode))
}
// UpdateSlbGroup updates SlbGroup
func (cli *ZSClient) UpdateSlbGroup(uuid string, params param.UpdateSlbGroupParam) (*view.SlbGroupInventoryView, error) {
	var resp view.UpdateSlbGroupEventView
	if err := cli.Put("v1/load-balancers/slb/group/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
