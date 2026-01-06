// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePreconfigurationTemplate deletes PreconfigurationTemplate
func (cli *ZSClient) DeletePreconfigurationTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/baremetal/preconfigurations/{uuid}", uuid, string(deleteMode))
}
// AddPreconfigurationTemplate adds PreconfigurationTemplate
func (cli *ZSClient) AddPreconfigurationTemplate(params param.AddPreconfigurationTemplateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	var resp view.AddPreconfigurationTemplateEventView
	if err := cli.Post("v1/baremetal/preconfigurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdatePreconfigurationTemplate updates PreconfigurationTemplate
func (cli *ZSClient) UpdatePreconfigurationTemplate(uuid string, params param.UpdatePreconfigurationTemplateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	var resp view.UpdatePreconfigurationTemplateEventView
	if err := cli.Put("v1/baremetal/preconfigurations/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryPreconfigurationTemplate queries PreconfigurationTemplate list
func (cli *ZSClient) QueryPreconfigurationTemplate(params *param.QueryParam) ([]view.PreconfigurationTemplateInventoryView, error) {
	var resp []view.PreconfigurationTemplateInventoryView
	return resp, cli.List("v1/baremetal/preconfigurations", params, &resp)
}
