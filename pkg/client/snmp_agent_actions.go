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
	resp := view.SnmpAgentInventoryView{}
	if err := cli.Post("v1/snmp/agent", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StartSnmpAgent starts SnmpAgent
func (cli *ZSClient) StartSnmpAgent(params param.StartSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PutWithRespKey("v1/snmp/agent/actions", "", "", map[string]interface{}{
		"startSnmpAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// StopSnmpAgent stops SnmpAgent
func (cli *ZSClient) StopSnmpAgent(params param.StopSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PutWithRespKey("v1/snmp/agent/actions", "", "", map[string]interface{}{
		"stopSnmpAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateSnmpAgent updates SnmpAgent
func (cli *ZSClient) UpdateSnmpAgent(params param.UpdateSnmpAgentParam) (*view.SnmpAgentInventoryView, error) {
	resp := view.SnmpAgentInventoryView{}
	if err := cli.PutWithRespKey("v1/snmp/agent/actions", "", "", map[string]interface{}{
		"updateSnmpAgent": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// PageSnmpAgent Pagination
func (cli *ZSClient) PageSnmpAgent(params *param.QueryParam) ([]view.SnmpAgentInventoryView, int, error) {
	var snmpAgents []view.SnmpAgentInventoryView
	total, err := cli.Page("v1/snmp/agent", params, &snmpAgents)
	return snmpAgents, total, err
}
