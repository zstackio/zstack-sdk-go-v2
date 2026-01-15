// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryGlobalConfigTemplate queries GlobalConfigTemplate list
func (cli *ZSClient) QueryGlobalConfigTemplate(params *param.QueryParam) ([]view.GlobalConfigTemplateInventoryView, error) {
	var resp []view.GlobalConfigTemplateInventoryView
	return resp, cli.List("v1/template-configurations/templates", params, &resp)
}

func (cli *ZSClient) GetGlobalConfigTemplate(uuid string) (*view.GlobalConfigTemplateInventoryView, error) {
	var resp view.GlobalConfigTemplateInventoryView
	if err := cli.Get("v1/template-configurations/templates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGlobalConfigTemplate Pagination
func (cli *ZSClient) PageGlobalConfigTemplate(params *param.QueryParam) ([]view.GlobalConfigTemplateInventoryView, int, error) {
	var globalConfigTemplates []view.GlobalConfigTemplateInventoryView
	total, err := cli.Page("v1/template-configurations/templates", params, &globalConfigTemplates)
	return globalConfigTemplates, total, err
}
