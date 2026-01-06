// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateImagePackage updates ImagePackage
func (cli *ZSClient) UpdateImagePackage(uuid string, params param.UpdateImagePackageParam) (*view.ImagePackageInventoryView, error) {
	var resp view.UpdateImagePackageEventView
	if err := cli.Put("v1/image-packages/{uuid}", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryImagePackage queries ImagePackage list
func (cli *ZSClient) QueryImagePackage(params *param.QueryParam) ([]view.ImagePackageInventoryView, error) {
	var resp []view.ImagePackageInventoryView
	return resp, cli.List("v1/image-packages", params, &resp)
}
// DeleteImagePackage deletes ImagePackage
func (cli *ZSClient) DeleteImagePackage(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-packages/{uuid}", uuid, string(deleteMode))
}
