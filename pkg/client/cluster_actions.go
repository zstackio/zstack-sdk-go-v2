// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

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

// ChangeClusterState changes ClusterState
func (cli *ZSClient) ChangeClusterState(uuid string, params param.ChangeClusterStateParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.PutWithRespKey("v1/clusters", uuid, "", map[string]interface{}{
		"changeClusterState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateClusterSupportDRS operates on ClusterSupportDRS
func (cli *ZSClient) ValidateClusterSupportDRS(uuid string) (*view.ValidateClusterSupportDRSView, error) {
	var resp view.ValidateClusterSupportDRSView
	if err := cli.GetWithRespKey("v1/clusters", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMiniCluster creates MiniCluster
func (cli *ZSClient) CreateMiniCluster(params param.CreateMiniClusterParam) (*view.ClusterInventoryView, error) {
	resp := view.ClusterInventoryView{}
	if err := cli.Post("v1/mini-clusters", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateClusterOS updates ClusterOS
func (cli *ZSClient) UpdateClusterOS(uuid string, params param.UpdateClusterOSParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.PutWithRespKey("v1/clusters", uuid, "", map[string]interface{}{
		"updateClusterOS": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateClusterOSAsync Async
func (cli *ZSClient) UpdateClusterOSAsync(uuid string, params param.UpdateClusterOSParam) (string, error) {

	resource := fmt.Sprintf("v1/clusters/%s/actions", uuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// GetClusterHostNetworkFacts gets ClusterHostNetworkFacts by uuid
func (cli *ZSClient) GetClusterHostNetworkFacts(uuid string) (*view.GetClusterHostNetworkFactsView, error) {
	var resp view.GetClusterHostNetworkFactsView
	if err := cli.GetWithRespKey("v1/cluster/hosts-network-facts", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
