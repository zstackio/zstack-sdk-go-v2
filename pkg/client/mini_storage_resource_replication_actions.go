// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryMiniStorageResourceReplication queries MiniStorageResourceReplication list
func (cli *ZSClient) QueryMiniStorageResourceReplication(params *param.QueryParam) ([]view.MiniStorageResourceReplicationInventoryView, error) {
	var resp []view.MiniStorageResourceReplicationInventoryView
	return resp, cli.List("v1/primary-storage/mini/replications", params, &resp)
}

func (cli *ZSClient) GetMiniStorageResourceReplication(uuid string) (*view.MiniStorageResourceReplicationInventoryView, error) {
	var resp view.MiniStorageResourceReplicationInventoryView
	if err := cli.Get("v1/primary-storage/mini/replications", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMiniStorageResourceReplication Pagination
func (cli *ZSClient) PageMiniStorageResourceReplication(params *param.QueryParam) ([]view.MiniStorageResourceReplicationInventoryView, int, error) {
	var miniStorageResourceReplications []view.MiniStorageResourceReplicationInventoryView
	total, err := cli.Page("v1/primary-storage/mini/replications", params, &miniStorageResourceReplications)
	return miniStorageResourceReplications, total, err
}
