// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateMonitorTemplate updates MonitorTemplate
func (cli *ZSClient) UpdateMonitorTemplate(uuid string, params param.UpdateMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	resp := view.MonitorTemplateInventoryView{}
	if err := cli.Put("v1/zwatch/monitortemplates", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// CloneMonitorTemplate operates on MonitorTemplate
func (cli *ZSClient) CloneMonitorTemplate(params param.CloneMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	resp := view.MonitorTemplateInventoryView{}
	if err := cli.Post("v1/zwatch/monitortemplates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryMonitorTemplate queries MonitorTemplate list
func (cli *ZSClient) QueryMonitorTemplate(params *param.QueryParam) ([]view.MonitorTemplateInventoryView, error) {
	var resp []view.MonitorTemplateInventoryView
	return resp, cli.List("v1/zwatch/monitortemplates", params, &resp)
}

func (cli *ZSClient) GetMonitorTemplate(uuid string) (*view.MonitorTemplateInventoryView, error) {
	var resp view.MonitorTemplateInventoryView
	if err := cli.Get("v1/zwatch/monitortemplates", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageMonitorTemplate Pagination
func (cli *ZSClient) PageMonitorTemplate(params *param.QueryParam) ([]view.MonitorTemplateInventoryView, int, error) {
	var monitorTemplates []view.MonitorTemplateInventoryView
	total, err := cli.Page("v1/zwatch/monitortemplates", params, &monitorTemplates)
	return monitorTemplates, total, err
}
// CreateMonitorTemplate creates MonitorTemplate
func (cli *ZSClient) CreateMonitorTemplate(params param.CreateMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	resp := view.MonitorTemplateInventoryView{}
	if err := cli.Post("v1/zwatch/monitortemplates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteMonitorTemplate deletes MonitorTemplate
func (cli *ZSClient) DeleteMonitorTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates", uuid, string(deleteMode))
}
