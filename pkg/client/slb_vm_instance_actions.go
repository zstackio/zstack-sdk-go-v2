// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySlbVmInstance queries SlbVmInstance list
func (cli *ZSClient) QuerySlbVmInstance(params *param.QueryParam) ([]view.SlbVmInstanceInventoryView, error) {
	var resp []view.SlbVmInstanceInventoryView
	return resp, cli.List("v1/load-balancers/slb/instances", params, &resp)
}

func (cli *ZSClient) GetSlbVmInstance(uuid string) (*view.SlbVmInstanceInventoryView, error) {
	var resp view.SlbVmInstanceInventoryView
	if err := cli.Get("v1/load-balancers/slb/instances", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSlbVmInstance Pagination
func (cli *ZSClient) PageSlbVmInstance(params *param.QueryParam) ([]view.SlbVmInstanceInventoryView, int, error) {
	var slbVmInstances []view.SlbVmInstanceInventoryView
	total, err := cli.Page("v1/load-balancers/slb/instances", params, &slbVmInstances)
	return slbVmInstances, total, err
}

// CreateSlbInstance creates SlbInstance
func (cli *ZSClient) CreateSlbInstance(params param.CreateSlbInstanceParam) (*view.SlbVmInstanceInventoryView, error) {
	resp := view.SlbVmInstanceInventoryView{}
	if err := cli.Post("v1/load-balancers/slb/instances", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProvisionSlbInstance operates on ProvisionSlbInstance
func (cli *ZSClient) ProvisionSlbInstance(uuid string, params param.ProvisionSlbInstanceParam) (*view.SlbGroupInventoryView, error) {
	resp := view.SlbGroupInventoryView{}
	if err := cli.PutWithRespKey("v1/load-balancers/slb/instances", uuid, "", map[string]interface{}{
		"provisionSlbInstance": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
