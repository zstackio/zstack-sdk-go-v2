// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteCluster deletes Cluster
func (cli *ZSClient) DeleteCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/{uuid}", uuid, string(deleteMode))
}
// UpdateCluster updates Cluster
func (cli *ZSClient) UpdateCluster(uuid string, params param.UpdateClusterParam) (*view.ClusterInventoryView, error) {
	var resp view.UpdateClusterEventView
	if err := cli.Put("v1/clusters/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CreateCluster creates Cluster
func (cli *ZSClient) CreateCluster(params param.CreateClusterParam) (*view.ClusterInventoryView, error) {
	var resp view.CreateClusterEventView
	if err := cli.Post("v1/clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QueryCluster queries Cluster list
func (cli *ZSClient) QueryCluster(params *param.QueryParam) ([]view.ClusterInventoryView, error) {
	var resp []view.ClusterInventoryView
	return resp, cli.List("v1/clusters", params, &resp)
}
