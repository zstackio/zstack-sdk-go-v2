// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLogConfiguration updates LogConfiguration
func (cli *ZSClient) UpdateLogConfiguration(uuid string, params param.UpdateLogConfigurationParam) (*view.JsonLabelInventoryView, error) {
	resp := view.JsonLabelInventoryView{}
	if err := cli.Put("v1/log/configurations", uuid, map[string]interface{}{
		"updateLogConfiguration": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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
	resp := view.JsonLabelInventoryView{}
	if err := cli.Post("v1/log/configurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
