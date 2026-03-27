// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryExternalServiceConfiguration queries ExternalServiceConfiguration list
func (cli *ZSClient) QueryExternalServiceConfiguration(ctx context.Context, params *param.QueryParam) ([]view.ExternalServiceConfigurationInventoryView, error) {
	var resp []view.ExternalServiceConfigurationInventoryView
	return resp, cli.List(ctx, "v1/external/service/configuration", params, &resp)
}

func (cli *ZSClient) GetExternalServiceConfiguration(ctx context.Context, uuid string) (*view.ExternalServiceConfigurationInventoryView, error) {
	var resp view.ExternalServiceConfigurationInventoryView
	if err := cli.Get(ctx, "v1/external/service/configuration", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageExternalServiceConfiguration Pagination
func (cli *ZSClient) PageExternalServiceConfiguration(ctx context.Context, params *param.QueryParam) ([]view.ExternalServiceConfigurationInventoryView, int, error) {
	var externalServiceConfigurations []view.ExternalServiceConfigurationInventoryView
	total, err := cli.Page(ctx, "v1/external/service/configuration", params, &externalServiceConfigurations)
	return externalServiceConfigurations, total, err
}
// UpdateExternalServiceConfiguration updates ExternalServiceConfiguration
func (cli *ZSClient) UpdateExternalServiceConfiguration(ctx context.Context, uuid string, params param.UpdateExternalServiceConfigurationParam) (*view.ExternalServiceConfigurationInventoryView, error) {
	resp := view.ExternalServiceConfigurationInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/external/service/configuration", uuid, "", map[string]interface{}{
		"updateExternalServiceConfiguration": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// AddExternalServiceConfiguration adds ExternalServiceConfiguration
func (cli *ZSClient) AddExternalServiceConfiguration(ctx context.Context, params param.AddExternalServiceConfigurationParam) (*view.ExternalServiceConfigurationInventoryView, error) {
	resp := view.ExternalServiceConfigurationInventoryView{}
	if err := cli.Post(ctx, "v1/external/service/configuration", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteExternalServiceConfiguration deletes ExternalServiceConfiguration
func (cli *ZSClient) DeleteExternalServiceConfiguration(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/external/service/configuration", uuid, string(deleteMode))
}
