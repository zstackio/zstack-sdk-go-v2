// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteCluster deletes Cluster
func (cli *ZSClient) DeleteCluster(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters", uuid, string(deleteMode))
}
// UpdateCluster updates Cluster
func (cli *ZSClient) UpdateCluster(uuid string, params param.UpdateClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PutWithRespKey("v1/clusters", uuid, "", map[string]interface{}{
		"updateCluster": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateCluster creates Cluster
func (cli *ZSClient) CreateCluster(params param.CreateClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.Post("v1/clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCluster queries Cluster list
func (cli *ZSClient) QueryCluster(params *param.QueryParam) ([]view.ClusterInventoryView, error) {
	var resp []view.ClusterInventoryView
	return resp, cli.List("v1/clusters", params, &resp)
}

func (cli *ZSClient) GetCluster(uuid string) (*view.ClusterInventoryView, error) {
	var resp view.ClusterInventoryView
	if err := cli.Get("v1/clusters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCluster Pagination
func (cli *ZSClient) PageCluster(params *param.QueryParam) ([]view.ClusterInventoryView, int, error) {
	var clusters []view.ClusterInventoryView
	total, err := cli.Page("v1/clusters", params, &clusters)
	return clusters, total, err
}
