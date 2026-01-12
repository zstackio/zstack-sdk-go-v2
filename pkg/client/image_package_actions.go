// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateImagePackage updates ImagePackage
func (cli *ZSClient) UpdateImagePackage(uuid string, params param.UpdateImagePackageParam) (*view.ImagePackageInventoryView, error) {
	var resp view.UpdateImagePackageEventView
	err := cli.PutWithSpec("v1/image-packages", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryImagePackage queries ImagePackage list
func (cli *ZSClient) QueryImagePackage(params *param.QueryParam) ([]view.ImagePackageInventoryView, error) {
	var resp []view.ImagePackageInventoryView
	return resp, cli.List("v1/image-packages", params, &resp)
}

func (cli *ZSClient) GetImagePackage(uuid string) (*view.ImagePackageInventoryView, error) {
	var resp view.ImagePackageInventoryView
	if err := cli.Get("v1/image-packages", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteImagePackage deletes ImagePackage
func (cli *ZSClient) DeleteImagePackage(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/image-packages", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
