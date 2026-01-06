// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddImage adds Image
func (cli *ZSClient) AddImage(params param.AddImageParam) (*view.ImageInventoryView, error) {
	var resp view.AddImageEventView
	if err := cli.Post("v1/images", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryImage queries Image list
func (cli *ZSClient) QueryImage(params *param.QueryParam) ([]view.ImageInventoryView, error) {
	var resp []view.ImageInventoryView
	return resp, cli.List("v1/images", params, &resp)
}
// SyncImage operates on Image
func (cli *ZSClient) SyncImage(uuid string, params param.SyncImageParam) (*view.ImageInventoryView, error) {
	resp := view.ImageInventoryView{}
	if err := cli.Put("v1/backup-storage/image-store/{imageStoreUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// RecoverImage operates on Image
func (cli *ZSClient) RecoverImage(uuid string, params param.RecoverImageParam) (*view.ImageInventoryView, error) {
	var resp view.RecoverImageEventView
	if err := cli.Put("v1/images/{imageUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CloneImage operates on Image
func (cli *ZSClient) CloneImage(params param.CloneImageParam) (*view.ImageInventoryView, error) {
	var resp view.CloneImageEventView
	if err := cli.Post("v1/image/clone/{imageUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteImage deletes Image
func (cli *ZSClient) DeleteImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/images/{uuid}", uuid, string(deleteMode))
}
// UpdateImage updates Image
func (cli *ZSClient) UpdateImage(uuid string, params param.UpdateImageParam) (*view.ImageInventoryView, error) {
	var resp view.UpdateImageEventView
	if err := cli.Put("v1/images/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ExpungeImage operates on Image
func (cli *ZSClient) ExpungeImage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/images/{imageUuid}/actions", uuid, string(deleteMode))
}
