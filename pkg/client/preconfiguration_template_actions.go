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
	return cli.Delete("v1/baremetal/preconfigurations", uuid, string(deleteMode))
}
// AddPreconfigurationTemplate adds PreconfigurationTemplate
func (cli *ZSClient) AddPreconfigurationTemplate(params param.AddPreconfigurationTemplateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	resp := view.PreconfigurationTemplateInventoryView{}
	if err := cli.Post("v1/baremetal/preconfigurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdatePreconfigurationTemplate updates PreconfigurationTemplate
func (cli *ZSClient) UpdatePreconfigurationTemplate(uuid string, params param.UpdatePreconfigurationTemplateParam) (*view.PreconfigurationTemplateInventoryView, error) {
	resp := view.PreconfigurationTemplateInventoryView{}
	if err := cli.Put("v1/baremetal/preconfigurations", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PagePreconfigurationTemplate Pagination
func (cli *ZSClient) PagePreconfigurationTemplate(params *param.QueryParam) ([]view.PreconfigurationTemplateInventoryView, int, error) {
	var preconfigurationTemplates []view.PreconfigurationTemplateInventoryView
	total, err := cli.Page("v1/baremetal/preconfigurations", params, &preconfigurationTemplates)
	return preconfigurationTemplates, total, err
}
