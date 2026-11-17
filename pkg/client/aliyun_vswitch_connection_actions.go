// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch updates ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) UpdateConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, params param.UpdateConnectionBetweenL3NetWorkAndAliyunVSwitchParam) (*view.ConnectionRelationShipInventoryView, error) {
	resp := view.ConnectionRelationShipInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/connections", uuid, "", map[string]interface{}{
		"updateConnectionBetweenL3NetWorkAndAliyunVSwitch": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch deletes ConnectionBetweenL3NetWorkAndAliyunVSwitch
func (cli *ZSClient) DeleteConnectionBetweenL3NetWorkAndAliyunVSwitch(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/connections", uuid, string(deleteMode))
}

// CreateConnectionBetweenL3NetworkAndAliyunVSwitch creates ConnectionBetweenL3NetworkAndAliyunVSwitch
func (cli *ZSClient) CreateConnectionBetweenL3NetworkAndAliyunVSwitch(params param.CreateConnectionBetweenL3NetworkAndAliyunVSwitchParam) (*view.ConnectionRelationShipInventoryView, error) {
	resp := view.ConnectionRelationShipInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/connections", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryConnectionBetweenL3NetworkAndAliyunVSwitch queries ConnectionBetweenL3NetworkAndAliyunVSwitch list
func (cli *ZSClient) QueryConnectionBetweenL3NetworkAndAliyunVSwitch(params *param.QueryParam) ([]view.ConnectionRelationShipInventoryView, error) {
	var resp []view.ConnectionRelationShipInventoryView
	return resp, cli.List("v1/hybrid/aliyun/relationships", params, &resp)
}

func (cli *ZSClient) GetConnectionBetweenL3NetworkAndAliyunVSwitch(uuid string) (*view.ConnectionRelationShipInventoryView, error) {
	var resp view.ConnectionRelationShipInventoryView
	if err := cli.Get("v1/hybrid/aliyun/relationships", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageConnectionBetweenL3NetworkAndAliyunVSwitch Pagination
func (cli *ZSClient) PageConnectionBetweenL3NetworkAndAliyunVSwitch(params *param.QueryParam) ([]view.ConnectionRelationShipInventoryView, int, error) {
	var connectionBetweenL3NetworkAndAliyunVSwitchs []view.ConnectionRelationShipInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/relationships", params, &connectionBetweenL3NetworkAndAliyunVSwitchs)
	return connectionBetweenL3NetworkAndAliyunVSwitchs, total, err
}
