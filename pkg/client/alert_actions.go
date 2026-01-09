// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryAlert queries Alert list
func (cli *ZSClient) QueryAlert(params *param.QueryParam) ([]view.AlertInventoryView, error) {
	var resp []view.AlertInventoryView
	return resp, cli.List("v1/monitoring/alerts", params, &resp)
}

func (cli *ZSClient) GetAlert(uuid string) (*view.AlertInventoryView, error) {
	var resp view.AlertInventoryView
	if err := cli.Get("v1/monitoring/alerts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteAlert deletes Alert
func (cli *ZSClient) DeleteAlert(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/monitoring/alerts", uuid, string(deleteMode))
}
