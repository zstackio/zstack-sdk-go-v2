// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLogConfiguration updates LogConfiguration
func (cli *ZSClient) UpdateLogConfiguration(uuid string, params param.UpdateLogConfigurationParam) (*view.JsonLabelInventoryView, error) {
	var resp view.UpdateLogConfigurationEventView
	if err := cli.Put("v1/log/configurations", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
// GetLogConfiguration gets LogConfiguration by uuid
func (cli *ZSClient) GetLogConfiguration(uuid string) (*view.JsonLabelInventoryView, error) {
	var resp view.JsonLabelInventoryView
	if err := cli.Get("v1/log/configurations", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLogConfiguration deletes LogConfiguration
func (cli *ZSClient) DeleteLogConfiguration(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/log/configurations/log4j2", uuid, string(deleteMode))
}
// AddLogConfiguration adds LogConfiguration
func (cli *ZSClient) AddLogConfiguration(params param.AddLogConfigurationParam) (*view.JsonLabelInventoryView, error) {
	var resp view.AddLogConfigurationEventView
	if err := cli.Post("v1/log/configurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
