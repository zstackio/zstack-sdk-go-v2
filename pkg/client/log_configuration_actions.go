// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateLogConfiguration updates LogConfiguration
func (cli *ZSClient) UpdateLogConfiguration(ctx context.Context, params param.UpdateLogConfigurationParam) (*view.JsonLabelInventoryView, error) {
	resp := view.JsonLabelInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/log/configurations", "", "", map[string]interface{}{
		"updateLogConfiguration": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// GetLogConfiguration gets LogConfiguration by uuid
func (cli *ZSClient) GetLogConfiguration(ctx context.Context) (*view.GetLogConfigurationView, error) {
	var resp view.GetLogConfigurationView
	if err := cli.GetWithRespKey(ctx, "v1/log/configurations", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteLogConfiguration deletes LogConfiguration
func (cli *ZSClient) DeleteLogConfiguration(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/log/configurations/log4j2", uuid, string(deleteMode))
}
// AddLogConfiguration adds LogConfiguration
func (cli *ZSClient) AddLogConfiguration(ctx context.Context, params param.AddLogConfigurationParam) (*view.JsonLabelInventoryView, error) {
	resp := view.JsonLabelInventoryView{}
	if err := cli.Post(ctx, "v1/log/configurations", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
