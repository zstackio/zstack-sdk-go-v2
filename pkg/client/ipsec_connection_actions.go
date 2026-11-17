// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ReconnectIPsecConnection operates on IPsecConnection
func (cli *ZSClient) ReconnectIPsecConnection(uuid string, params param.ReconnectIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithRespKey("v1/ipsec", uuid, "", map[string]interface{}{
		"reconnectIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateIPsecConnection updates IPsecConnection
func (cli *ZSClient) UpdateIPsecConnection(uuid string, params param.UpdateIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithRespKey("v1/ipsec", uuid, "", map[string]interface{}{
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
	if err := cli.PutWithRespKey("v1/ipsec/config", uuid, "", map[string]interface{}{
		"changeIPsecConnection": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryIPSecConnection queries IPSecConnection list
func (cli *ZSClient) QueryIPSecConnection(params *param.QueryParam) ([]view.IPsecConnectionInventoryView, error) {
	var resp []view.IPsecConnectionInventoryView
	return resp, cli.List("v1/ipsec", params, &resp)
}

func (cli *ZSClient) GetIPSecConnection(uuid string) (*view.IPsecConnectionInventoryView, error) {
	var resp view.IPsecConnectionInventoryView
	if err := cli.Get("v1/ipsec", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageIPSecConnection Pagination
func (cli *ZSClient) PageIPSecConnection(params *param.QueryParam) ([]view.IPsecConnectionInventoryView, int, error) {
	var iPSecConnections []view.IPsecConnectionInventoryView
	total, err := cli.Page("v1/ipsec", params, &iPSecConnections)
	return iPSecConnections, total, err
}

// ChangeIPSecConnectionState changes IPSecConnectionState
func (cli *ZSClient) ChangeIPSecConnectionState(uuid string, params param.ChangeIPSecConnectionStateParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.PutWithRespKey("v1/ipsec", uuid, "", map[string]interface{}{
		"changeIPSecConnectionState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachL3NetworksToIPsecConnection operates on L3NetworksToIPsecConnection
func (cli *ZSClient) AttachL3NetworksToIPsecConnection(uuid string, params param.AttachL3NetworksToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/ipsec/%s/l3networks", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcAttachedIpsec gets VpcAttachedIpsec by uuid
func (cli *ZSClient) GetVpcAttachedIpsec(uuid string, params param.GetVpcAttachedIpsecParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/vpc/virtual-routers/%s/attached-ipsec", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpcIPsecLog gets VpcIPsecLog by uuid
func (cli *ZSClient) GetVpcIPsecLog() (*view.GetVpcIPsecLogView, error) {
	var resp view.GetVpcIPsecLogView
	if err := cli.GetWithRespKey("v1/vpc/virtual-routers/ipseclog", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRemoteCidrsToIPsecConnection adds RemoteCidrsToIPsecConnection
func (cli *ZSClient) AddRemoteCidrsToIPsecConnection(uuid string, params param.AddRemoteCidrsToIPsecConnectionParam) (*view.IPsecConnectionInventoryView, error) {
	resp := view.IPsecConnectionInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/ipsec/%s/remote-cidrs", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachL3NetworksFromIPsecConnection operates on L3NetworksFromIPsecConnection
func (cli *ZSClient) DetachL3NetworksFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec", uuid, string(deleteMode))
}

// RemoveRemoteCidrsFromIPsecConnection removes RemoteCidrsFromIPsecConnection
func (cli *ZSClient) RemoveRemoteCidrsFromIPsecConnection(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/ipsec", uuid, string(deleteMode))
}
