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
	var resp view.ReconnectIPsecConnectionEventView
	err := cli.PutWithSpec("v1/ipsec", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateIPsecConnection updates IPsecConnection
func (cli *ZSClient) UpdateIPsecConnection(uuid string, params param.UpdateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.UpdateIPsecConnectionEventView
	err := cli.PutWithSpec("v1/ipsec", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteIPsecConnection deletes IPsecConnection
func (cli *ZSClient) DeleteIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/ipsec", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
// CreateIPsecConnection creates IPsecConnection
func (cli *ZSClient) CreateIPsecConnection(params param.CreateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.CreateIPsecConnectionEventView
	if err := cli.Post("v1/ipsec", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// ChangeIPsecConnection changes IPsecConnection
func (cli *ZSClient) ChangeIPsecConnection(uuid string, params param.ChangeIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	var resp view.ChangeIPsecConnectionEventView
	err := cli.PutWithSpec("v1/ipsec/config", fmt.Sprintf(\"%s\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
