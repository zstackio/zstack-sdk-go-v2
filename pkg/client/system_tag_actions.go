// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateSystemTag updates SystemTag
func (cli *ZSClient) UpdateSystemTag(uuid string, params param.UpdateSystemTagParam) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.Put("v1/system-tags", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSystemTag creates SystemTag
func (cli *ZSClient) CreateSystemTag(params param.CreateSystemTagParam) (*view.SystemTagInventoryView, error) {
	resp := view.SystemTagInventoryView{}
	if err := cli.Post("v1/system-tags", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QuerySystemTag queries SystemTag list
func (cli *ZSClient) QuerySystemTag(params *param.QueryParam) ([]view.SystemTagInventoryView, error) {
	var resp []view.SystemTagInventoryView
	return resp, cli.List("v1/system-tags", params, &resp)
}

func (cli *ZSClient) GetSystemTag(uuid string) (*view.SystemTagInventoryView, error) {
	var resp view.SystemTagInventoryView
	if err := cli.Get("v1/system-tags", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSystemTag Pagination
func (cli *ZSClient) PageSystemTag(params *param.QueryParam) ([]view.SystemTagInventoryView, int, error) {
	var systemTags []view.SystemTagInventoryView
	total, err := cli.Page("v1/system-tags", params, &systemTags)
	return systemTags, total, err
}
