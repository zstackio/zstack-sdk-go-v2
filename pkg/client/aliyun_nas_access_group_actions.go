// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddAliyunNasAccessGroup adds AliyunNasAccessGroup
func (cli *ZSClient) AddAliyunNasAccessGroup(params param.AddAliyunNasAccessGroupParam) (*view.AliyunNasAccessGroupInventoryView, error) {
	var resp view.AddAliyunNasAccessGroupEventView
	if err := cli.Post("v1/nas/aliyun/access", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateAliyunNasAccessGroup updates AliyunNasAccessGroup
func (cli *ZSClient) UpdateAliyunNasAccessGroup(uuid string, params param.UpdateAliyunNasAccessGroupParam) (*view.AliyunNasAccessGroupInventoryView, error) {
	var resp view.UpdateAliyunNasAccessGroupEventView
	if err := cli.Put("v1/nas/aliyun/access", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateAliyunNasAccessGroup creates AliyunNasAccessGroup
func (cli *ZSClient) CreateAliyunNasAccessGroup(params param.CreateAliyunNasAccessGroupParam) (*view.AliyunNasAccessGroupInventoryView, error) {
	var resp view.CreateAliyunNasAccessGroupEventView
	if err := cli.Post("v1/nas/aliyun/access", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteAliyunNasAccessGroup deletes AliyunNasAccessGroup
func (cli *ZSClient) DeleteAliyunNasAccessGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/nas/access", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// QueryAliyunNasAccessGroup queries AliyunNasAccessGroup list
func (cli *ZSClient) QueryAliyunNasAccessGroup(params *param.QueryParam) ([]view.AliyunNasAccessGroupInventoryView, error) {
	var resp []view.AliyunNasAccessGroupInventoryView
	return resp, cli.List("v1/nas/aliyun/access", params, &resp)
}

func (cli *ZSClient) GetAliyunNasAccessGroup(uuid string) (*view.AliyunNasAccessGroupInventoryView, error) {
	var resp view.AliyunNasAccessGroupInventoryView
	if err := cli.Get("v1/nas/aliyun/access", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
