// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryVCenterCluster queries VCenterCluster list
func (cli *ZSClient) QueryVCenterCluster(params *param.QueryParam) ([]view.VCenterClusterInventoryView, error) {
	var resp []view.VCenterClusterInventoryView
	return resp, cli.List("v1/vcenters/clusters", params, &resp)
}

func (cli *ZSClient) GetVCenterCluster(uuid string) (*view.VCenterClusterInventoryView, error) {
	var resp view.VCenterClusterInventoryView
	if err := cli.Get("v1/vcenters/clusters", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageVCenterCluster Pagination
func (cli *ZSClient) PageVCenterCluster(params *param.QueryParam) ([]view.VCenterClusterInventoryView, int, error) {
	var vCenterClusters []view.VCenterClusterInventoryView
	total, err := cli.Page("v1/vcenters/clusters", params, &vCenterClusters)
	return vCenterClusters, total, err
}
