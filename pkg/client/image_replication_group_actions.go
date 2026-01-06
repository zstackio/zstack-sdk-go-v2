// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteImageReplicationGroup deletes ImageReplicationGroup
func (cli *ZSClient) DeleteImageReplicationGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/image-replication-groups/{uuid}", uuid, string(deleteMode))
}
// QueryImageReplicationGroup queries ImageReplicationGroup list
func (cli *ZSClient) QueryImageReplicationGroup(params *param.QueryParam) ([]view.ImageReplicationGroupInventoryView, error) {
	var resp []view.ImageReplicationGroupInventoryView
	return resp, cli.List("v1/image-replication-groups", params, &resp)
}
// CreateImageReplicationGroup creates ImageReplicationGroup
func (cli *ZSClient) CreateImageReplicationGroup(params param.CreateImageReplicationGroupParam) (*view.ImageReplicationGroupInventoryView, error) {
	var resp view.CreateImageReplicationGroupEventView
	if err := cli.Post("v1/image-replication-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
