// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryImageGroupRef queries ImageGroupRef list
func (cli *ZSClient) QueryImageGroupRef(params *param.QueryParam) ([]view.ImageGroupRefInventoryView, error) {
	var resp []view.ImageGroupRefInventoryView
	return resp, cli.List("v1/imagegrouprefs", params, &resp)
}

func (cli *ZSClient) GetImageGroupRef(uuid string) (*view.ImageGroupRefInventoryView, error) {
	var resp view.ImageGroupRefInventoryView
	if err := cli.Get("v1/imagegrouprefs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageImageGroupRef Pagination
func (cli *ZSClient) PageImageGroupRef(params *param.QueryParam) ([]view.ImageGroupRefInventoryView, int, error) {
	var imageGroupRefs []view.ImageGroupRefInventoryView
	total, err := cli.Page("v1/imagegrouprefs", params, &imageGroupRefs)
	return imageGroupRefs, total, err
}
