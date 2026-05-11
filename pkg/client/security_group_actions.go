// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySecurityGroup queries SecurityGroup list
func (cli *ZSClient) QuerySecurityGroup(ctx context.Context, params *param.QueryParam) ([]view.SecurityGroupInventoryView, error) {
	var resp []view.SecurityGroupInventoryView
	return resp, cli.List(ctx, "v1/security-groups", params, &resp)
}

func (cli *ZSClient) GetSecurityGroup(ctx context.Context, uuid string) (*view.SecurityGroupInventoryView, error) {
	var resp view.SecurityGroupInventoryView
	if err := cli.Get(ctx, "v1/security-groups", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageSecurityGroup Pagination
func (cli *ZSClient) PageSecurityGroup(ctx context.Context, params *param.QueryParam) ([]view.SecurityGroupInventoryView, int, error) {
	var securityGroups []view.SecurityGroupInventoryView
	total, err := cli.Page(ctx, "v1/security-groups", params, &securityGroups)
	return securityGroups, total, err
}
// CreateSecurityGroup creates SecurityGroup
func (cli *ZSClient) CreateSecurityGroup(ctx context.Context, params param.CreateSecurityGroupParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PostWithRespKey(ctx, "v1/security-groups", "inventory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteSecurityGroup deletes SecurityGroup
func (cli *ZSClient) DeleteSecurityGroup(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/security-groups", uuid, string(deleteMode))
}
// UpdateSecurityGroup updates SecurityGroup
func (cli *ZSClient) UpdateSecurityGroup(ctx context.Context, uuid string, params param.UpdateSecurityGroupParam) (*view.SecurityGroupInventoryView, error) {
	resp := view.SecurityGroupInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/security-groups", uuid, "", map[string]interface{}{
		"updateSecurityGroup": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
