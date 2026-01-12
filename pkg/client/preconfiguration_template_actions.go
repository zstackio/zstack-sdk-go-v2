// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeletePreconfigurationTemplate deletes PreconfigurationTemplate
func (cli *ZSClient) DeletePreconfigurationTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/baremetal/preconfigurations", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
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
	err := cli.PutWithSpec("v1/baremetal/preconfigurations", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryPreconfigurationTemplate queries PreconfigurationTemplate list
func (cli *ZSClient) QueryPreconfigurationTemplate(params *param.QueryParam) ([]view.PreconfigurationTemplateInventoryView, error) {
	var resp []view.PreconfigurationTemplateInventoryView
	return resp, cli.List("v1/baremetal/preconfigurations", params, &resp)
}

func (cli *ZSClient) GetPreconfigurationTemplate(uuid string) (*view.PreconfigurationTemplateInventoryView, error) {
	var resp view.PreconfigurationTemplateInventoryView
	if err := cli.Get("v1/baremetal/preconfigurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
