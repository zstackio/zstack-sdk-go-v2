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
	var resp view.UpdateMonitorTemplateEventView
	err := cli.PutWithSpec("v1/zwatch/monitortemplates", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// CloneMonitorTemplate operates on MonitorTemplate
func (cli *ZSClient) CloneMonitorTemplate(uuid string, params param.CloneMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	var resp view.CloneMonitorTemplateEventView
	err := cli.PostWithSpec("v1/zwatch/monitortemplates", fmt.Sprintf(\"%s/actions\", uuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
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
// CreateMonitorTemplate creates MonitorTemplate
func (cli *ZSClient) CreateMonitorTemplate(params param.CreateMonitorTemplateParam) (*view.MonitorTemplateInventoryView, error) {
	var resp view.CreateMonitorTemplateEventView
	if err := cli.Post("v1/zwatch/monitortemplates", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// DeleteMonitorTemplate deletes MonitorTemplate
func (cli *ZSClient) DeleteMonitorTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/zwatch/monitortemplates", fmt.Sprintf(\"%s\", uuid), string(deleteMode))
}
