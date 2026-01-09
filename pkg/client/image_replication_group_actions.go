// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteImageReplicationGroup deletes ImageReplicationGroup
func (cli *ZSClient) DeleteImageReplicationGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-replication-groups", uuid, string(deleteMode))
}
// QueryImageReplicationGroup queries ImageReplicationGroup list
func (cli *ZSClient) QueryImageReplicationGroup(params *param.QueryParam) ([]view.ImageReplicationGroupInventoryView, error) {
	var resp []view.ImageReplicationGroupInventoryView
	return resp, cli.List("v1/image-replication-groups", params, &resp)
}

func (cli *ZSClient) GetImageReplicationGroup(uuid string) (*view.ImageReplicationGroupInventoryView, error) {
	var resp view.ImageReplicationGroupInventoryView
	if err := cli.Get("v1/image-replication-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateImageReplicationGroup creates ImageReplicationGroup
func (cli *ZSClient) CreateImageReplicationGroup(params param.CreateImageReplicationGroupParam) (*view.ImageReplicationGroupInventoryView, error) {
	var resp view.CreateImageReplicationGroupEventView
	if err := cli.Post("v1/image-replication-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
