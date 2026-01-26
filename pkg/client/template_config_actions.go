// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateTemplateConfig updates TemplateConfig
func (cli *ZSClient) UpdateTemplateConfig(templateUuid string, params param.UpdateTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithRespKey("v1/template-configurations", templateUuid, "", map[string]interface{}{
		"updateTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryTemplateConfig queries TemplateConfig list
func (cli *ZSClient) QueryTemplateConfig(params *param.QueryParam) ([]view.TemplateConfigInventoryView, error) {
	var resp []view.TemplateConfigInventoryView
	return resp, cli.List("v1/template-configurations/configs", params, &resp)
}

func (cli *ZSClient) GetTemplateConfig(uuid string) (*view.TemplateConfigInventoryView, error) {
	var resp view.TemplateConfigInventoryView
	if err := cli.Get("v1/template-configurations/configs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTemplateConfig Pagination
func (cli *ZSClient) PageTemplateConfig(params *param.QueryParam) ([]view.TemplateConfigInventoryView, int, error) {
	var templateConfigs []view.TemplateConfigInventoryView
	total, err := cli.Page("v1/template-configurations/configs", params, &templateConfigs)
	return templateConfigs, total, err
}
// RevertTemplateConfig operates on TemplateConfig
func (cli *ZSClient) RevertTemplateConfig(templateUuid string, params param.RevertTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithRespKey("v1/template-configurations", templateUuid, "", map[string]interface{}{
		"revertTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ApplyTemplateConfig operates on TemplateConfig
func (cli *ZSClient) ApplyTemplateConfig(templateUuid string, params param.ApplyTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithRespKey("v1/template-configurations", templateUuid, "", map[string]interface{}{
		"applyTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// ResetTemplateConfig operates on TemplateConfig
func (cli *ZSClient) ResetTemplateConfig(templateUuid string, params param.ResetTemplateConfigParam) (*view.TemplateConfigInventoryView, error) {
	resp := view.TemplateConfigInventoryView{}
	if err := cli.PutWithRespKey("v1/template-configurations", templateUuid, "", map[string]interface{}{
		"resetTemplateConfig": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
