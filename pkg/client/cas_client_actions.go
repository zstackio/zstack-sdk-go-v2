// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCasClient creates CasClient
func (cli *ZSClient) CreateCasClient(params param.CreateCasClientParam) (*view.CasClientInventoryView, error) {
	var resp view.CreateCasClientEventView
	if err := cli.Post("v1/create/cas/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateCasClient updates CasClient
func (cli *ZSClient) UpdateCasClient(uuid string, params param.UpdateCasClientParam) (*view.CasClientInventoryView, error) {
	var resp view.UpdateCasClientEventView
	if err := cli.Put("v1/update/cas/client", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
