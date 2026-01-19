// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectIPsecConnection operates on IPsecConnection
func (cli *ZSClient) ReconnectIPsecConnection(uuid string, params param.ReconnectIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Put("v1/ipsec", uuid, map[string]interface{}{
		"reconnectIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIPsecConnection updates IPsecConnection
func (cli *ZSClient) UpdateIPsecConnection(uuid string, params param.UpdateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Put("v1/ipsec", uuid, map[string]interface{}{
		"updateIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteIPsecConnection deletes IPsecConnection
func (cli *ZSClient) DeleteIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec", uuid, string(deleteMode))
}
// CreateIPsecConnection creates IPsecConnection
func (cli *ZSClient) CreateIPsecConnection(params param.CreateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post("v1/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ChangeIPsecConnection changes IPsecConnection
func (cli *ZSClient) ChangeIPsecConnection(uuid string, params param.ChangeIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Put("v1/ipsec/config", uuid, map[string]interface{}{
		"changeIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
