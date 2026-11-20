// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// DeleteHybridEipRemote requires additional arguments.
//
// Deprecated: use DeleteHybridEipRemoteWithParams to supply the required fields.
func (cli *ZSClient) DeleteHybridEipRemote(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteHybridEipRemote requires additional arguments; use DeleteHybridEipRemoteWithParams")
}

// DeleteHybridEipRemoteWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteHybridEipRemoteWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteHybridEipRemoteParamDetail) error {
	query := url.Values{
		"type": {params.Type},
	}
	return cli.deleteActionWithQuery("v1/hybrid/eip", uuid, "remote", deleteModeOrDefault(deleteMode, params.DeleteMode), query, "type")
}

// DeleteHybridEipFromLocal requires additional arguments.
//
// Deprecated: use DeleteHybridEipFromLocalWithParams to supply the required fields.
func (cli *ZSClient) DeleteHybridEipFromLocal(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteHybridEipFromLocal requires additional arguments; use DeleteHybridEipFromLocalWithParams")
}

// DeleteHybridEipFromLocalWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteHybridEipFromLocalWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteHybridEipFromLocalParamDetail) error {
	query := url.Values{
		"type": {params.Type},
	}
	return cli.deleteResourceWithQuery("v1/hybrid/eip", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query, "type")
}

// SyncHybridEipFromRemote operates on HybridEipFromRemote
func (cli *ZSClient) SyncHybridEipFromRemote(dataCenterUuid string, params param.SyncHybridEipFromRemoteParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/eip", dataCenterUuid, "", map[string]interface{}{
		"syncHybridEipFromRemote": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHybridEip updates HybridEip
func (cli *ZSClient) UpdateHybridEip(uuid string, params param.UpdateHybridEipParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/eip", uuid, "", map[string]interface{}{
		"updateHybridEip": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHybridEipFromEcs operates on HybridEipFromEcs
func (cli *ZSClient) DetachHybridEipFromEcs(eipUuid string, params param.DetachHybridEipFromEcsParam) (*view.DetachHybridEipFromEcsEventView, error) {
	resp := view.DetachHybridEipFromEcsEventView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/eip/%s/detach", eipUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryHybridEipFromLocal queries HybridEipFromLocal list
func (cli *ZSClient) QueryHybridEipFromLocal(params *param.QueryParam) ([]view.HybridEipAddressInventoryView, error) {
	var resp []view.HybridEipAddressInventoryView
	return resp, cli.List("v1/hybrid/eip", params, &resp)
}

func (cli *ZSClient) GetHybridEipFromLocal(uuid string) (*view.HybridEipAddressInventoryView, error) {
	var resp view.HybridEipAddressInventoryView
	if err := cli.Get("v1/hybrid/eip", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHybridEipFromLocal Pagination
func (cli *ZSClient) PageHybridEipFromLocal(params *param.QueryParam) ([]view.HybridEipAddressInventoryView, int, error) {
	var hybridEipFromLocals []view.HybridEipAddressInventoryView
	total, err := cli.Page("v1/hybrid/eip", params, &hybridEipFromLocals)
	return hybridEipFromLocals, total, err
}

// AttachHybridEipToEcs operates on HybridEipToEcs
func (cli *ZSClient) AttachHybridEipToEcs(eipUuid string, params param.AttachHybridEipToEcsParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/eip/%s/attach", eipUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateHybridEip creates HybridEip
func (cli *ZSClient) CreateHybridEip(params param.CreateHybridEipParam) (*view.HybridEipAddressInventoryView, error) {
	resp := view.HybridEipAddressInventoryView{}
	if err := cli.Post("v1/hybrid/eip", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
