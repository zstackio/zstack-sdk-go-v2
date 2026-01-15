// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryClusterDRS queries ClusterDRS list
func (cli *ZSClient) QueryClusterDRS(params *param.QueryParam) ([]view.ClusterDRSInventoryView, error) {
	var resp []view.ClusterDRSInventoryView
	return resp, cli.List("v1/clusters/drs", params, &resp)
}

// PageClusterDRS Pagination
func (cli *ZSClient) PageClusterDRS(params *param.QueryParam) ([]view.ClusterDRSInventoryView, int, error) {
	var clusterDRSs []view.ClusterDRSInventoryView
	total, err := cli.Page("v1/clusters/drs", params, &clusterDRSs)
	return clusterDRSs, total, err
}
// DeleteClusterDRS deletes ClusterDRS
func (cli *ZSClient) DeleteClusterDRS(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/clusters/drs", uuid, string(deleteMode))
}
// CreateClusterDRS creates ClusterDRS
func (cli *ZSClient) CreateClusterDRS(params param.CreateClusterDRSParam) (*view.ClusterDRSInventoryView, error) {
	resp := view.ClusterDRSInventoryView{}
	if err := cli.Post("v1/clusters/{clusterUuid}/drs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateClusterDRS updates ClusterDRS
func (cli *ZSClient) UpdateClusterDRS(uuid string, params param.UpdateClusterDRSParam) (*view.ClusterDRSInventoryView, error) {
	resp := view.ClusterDRSInventoryView{}
	if err := cli.Put("v1/clusters/drs", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
