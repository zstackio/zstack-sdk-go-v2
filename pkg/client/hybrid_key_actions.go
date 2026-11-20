// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// AttachHybridKey operates on HybridKey
func (cli *ZSClient) AttachHybridKey(uuid string, params param.AttachHybridKeyParam) (*view.AttachHybridKeyEventView, error) {
	resp := view.AttachHybridKeyEventView{}
	if err := cli.PutWithRespKey("v1/hybrid/hybrid/key", uuid, "", map[string]interface{}{
		"attachHybridKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachHybridKey operates on HybridKey
func (cli *ZSClient) DetachHybridKey(uuid string, params param.DetachHybridKeyParam) (*view.DetachHybridKeyEventView, error) {
	resp := view.DetachHybridKeyEventView{}
	if err := cli.PutWithRespKey("v1/hybrid/hybrid/key", uuid, "", map[string]interface{}{
		"detachHybridKey": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryHybridKeySecret queries HybridKeySecret list
func (cli *ZSClient) QueryHybridKeySecret(params *param.QueryParam) ([]view.HybridAccountInventoryView, error) {
	var resp []view.HybridAccountInventoryView
	return resp, cli.List("v1/hybrid/hybrid/key", params, &resp)
}

func (cli *ZSClient) GetHybridKeySecret(uuid string) (*view.HybridAccountInventoryView, error) {
	var resp view.HybridAccountInventoryView
	if err := cli.Get("v1/hybrid/hybrid/key", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageHybridKeySecret Pagination
func (cli *ZSClient) PageHybridKeySecret(params *param.QueryParam) ([]view.HybridAccountInventoryView, int, error) {
	var hybridKeySecrets []view.HybridAccountInventoryView
	total, err := cli.Page("v1/hybrid/hybrid/key", params, &hybridKeySecrets)
	return hybridKeySecrets, total, err
}

// DeleteHybridKeySecret deletes HybridKeySecret
func (cli *ZSClient) DeleteHybridKeySecret(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/hybrid/key", uuid, string(deleteMode))
}

// AddHybridKeySecret adds HybridKeySecret
func (cli *ZSClient) AddHybridKeySecret(params param.AddHybridKeySecretParam) (*view.HybridAccountInventoryView, error) {
	resp := view.HybridAccountInventoryView{}
	if err := cli.Post("v1/hybrid/hybrid/key", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHybridKeySecret updates HybridKeySecret
func (cli *ZSClient) UpdateHybridKeySecret(uuid string, params param.UpdateHybridKeySecretParam) (*view.HybridAccountInventoryView, error) {
	resp := view.HybridAccountInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/hybrid/hybrid/%s/key", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
