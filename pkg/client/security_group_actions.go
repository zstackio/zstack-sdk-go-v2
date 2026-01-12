// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySecurityGroup queries SecurityGroup list
func (cli *ZSClient) QuerySecurityGroup(params *param.QueryParam) ([]view.SecurityGroupInventoryView, error) {
	var resp []view.SecurityGroupInventoryView
	return resp, cli.List("v1/security-groups", params, &resp)
}

func (cli *ZSClient) GetSecurityGroup(uuid string) (*view.SecurityGroupInventoryView, error) {
	var resp view.SecurityGroupInventoryView
	if err := cli.Get("v1/security-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CreateSecurityGroup creates SecurityGroup
func (cli *ZSClient) CreateSecurityGroup(params param.CreateSecurityGroupParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.CreateSecurityGroupEventView
	if err := cli.Post("v1/security-groups", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteSecurityGroup deletes SecurityGroup
func (cli *ZSClient) DeleteSecurityGroup(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/security-groups", uuid, string(deleteMode))
}
// UpdateSecurityGroup updates SecurityGroup
func (cli *ZSClient) UpdateSecurityGroup(uuid string, params param.UpdateSecurityGroupParam) (*view.SecurityGroupInventoryView, error) {
	var resp view.UpdateSecurityGroupEventView
	if err := cli.Put("v1/security-groups", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
