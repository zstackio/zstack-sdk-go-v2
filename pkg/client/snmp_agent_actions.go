// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateSnmpAgent creates SnmpAgent
func (cli *ZSClient) CreateSnmpAgent(params param.CreateSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	var resp view.CreateSnmpAgentEventView
	if err := cli.Post("v1/snmp/agent", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StartSnmpAgent starts SnmpAgent
func (cli *ZSClient) StartSnmpAgent(uuid string, params param.StartSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	var resp view.StartSnmpAgentEventView
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// StopSnmpAgent stops SnmpAgent
func (cli *ZSClient) StopSnmpAgent(uuid string, params param.StopSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	var resp view.StopSnmpAgentEventView
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// UpdateSnmpAgent updates SnmpAgent
func (cli *ZSClient) UpdateSnmpAgent(uuid string, params param.UpdateSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	var resp view.UpdateSnmpAgentEventView
	if err := cli.Put("v1/snmp/agent/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// QuerySnmpAgent queries SnmpAgent list
func (cli *ZSClient) QuerySnmpAgent(params *param.QueryParam) ([]view.SnmpAgentInventoryView, error) {
	var resp []view.SnmpAgentInventoryView
	return resp, cli.List("v1/snmp/agent", params, &resp)
}

func (cli *ZSClient) GetSnmpAgent(uuid string) (*view.SnmpAgentInventoryView, error) {
	var resp view.SnmpAgentInventoryView
	if err := cli.Get("v1/snmp/agent", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
