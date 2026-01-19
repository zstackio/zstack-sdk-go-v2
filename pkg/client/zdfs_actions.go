// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectZdfs operates on Zdfs
func (cli *ZSClient) ReconnectZdfs(uuid string, params param.ReconnectZdfsParam) (*view.ZdfsInventoryView, error) {
	resp := view.ZdfsInventoryView{}
	if err := cli.Put("v1/zdfs", uuid, map[string]interface{}{
		"reconnectZdfs": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryZdfs queries Zdfs list
func (cli *ZSClient) QueryZdfs(params *param.QueryParam) ([]view.ZdfsInventoryView, error) {
	var resp []view.ZdfsInventoryView
	return resp, cli.List("v1/zdfs", params, &resp)
}

func (cli *ZSClient) GetZdfs(uuid string) (*view.ZdfsInventoryView, error) {
	var resp view.ZdfsInventoryView
	if err := cli.Get("v1/zdfs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageZdfs Pagination
func (cli *ZSClient) PageZdfs(params *param.QueryParam) ([]view.ZdfsInventoryView, int, error) {
	var zdfs []view.ZdfsInventoryView
	total, err := cli.Page("v1/zdfs", params, &zdfs)
	return zdfs, total, err
}
