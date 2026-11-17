// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// AddV2VConversionHost adds V2VConversionHost
func (cli *ZSClient) AddV2VConversionHost(params param.AddV2VConversionHostParam) (*view.V2VConversionHostInventoryView, error) {
	resp := view.V2VConversionHostInventoryView{}
	if err := cli.Post("v1/v2v-conversion-hosts", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// UpdateV2VConversionHost updates V2VConversionHost
func (cli *ZSClient) UpdateV2VConversionHost(uuid string, params param.UpdateV2VConversionHostParam) (*view.V2VConversionHostInventoryView, error) {
	resp := view.V2VConversionHostInventoryView{}
	if err := cli.PutWithRespKey("v1/v2v-conversion-hosts", uuid, "", map[string]interface{}{
		"updateV2VConversionHost": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryV2VConversionHost queries V2VConversionHost list
func (cli *ZSClient) QueryV2VConversionHost(params *param.QueryParam) ([]view.V2VConversionHostInventoryView, error) {
	var resp []view.V2VConversionHostInventoryView
	return resp, cli.List("v1/v2v-conversion-hosts", params, &resp)
}

func (cli *ZSClient) GetV2VConversionHost(uuid string) (*view.V2VConversionHostInventoryView, error) {
	var resp view.V2VConversionHostInventoryView
	if err := cli.Get("v1/v2v-conversion-hosts", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageV2VConversionHost Pagination
func (cli *ZSClient) PageV2VConversionHost(params *param.QueryParam) ([]view.V2VConversionHostInventoryView, int, error) {
	var v2VConversionHosts []view.V2VConversionHostInventoryView
	total, err := cli.Page("v1/v2v-conversion-hosts", params, &v2VConversionHosts)
	return v2VConversionHosts, total, err
}
// DeleteV2VConversionHost deletes V2VConversionHost
func (cli *ZSClient) DeleteV2VConversionHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/v2v-conversion-hosts", uuid, string(deleteMode))
}

// ConvertVmFromForeignHypervisor operates on ConvertVmFromForeignHypervisor
func (cli *ZSClient) ConvertVmFromForeignHypervisor(params param.ConvertVmFromForeignHypervisorParam) (*view.LongJobInventoryView, error) {
	resp := view.LongJobInventoryView{}
	if err := cli.Post("v1/v2vs", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConvertVmFromForeignHypervisorAsync Async
func (cli *ZSClient) ConvertVmFromForeignHypervisorAsync(params param.ConvertVmFromForeignHypervisorParam) (string, error) {

	resource := "v1/v2vs"
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}

// ChangeV2VConversionHostState changes V2VConversionHostState
func (cli *ZSClient) ChangeV2VConversionHostState(uuid string, params param.ChangeV2VConversionHostStateParam) (*view.V2VConversionHostInventoryView, error) {
	resp := view.V2VConversionHostInventoryView{}
	if err := cli.PutWithRespKey("v1/v2v-conversion-hosts", uuid, "", map[string]interface{}{
		"changeV2VConversionHostState": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanV2VConversionCache operates on V2VConversionCache
func (cli *ZSClient) CleanV2VConversionCache(uuid string, params param.CleanV2VConversionCacheParam) (*view.CleanV2VConversionCacheEventView, error) {
	resp := view.CleanV2VConversionCacheEventView{}
	if err := cli.PutWithRespKey("v1/v2v/conversion/host", uuid, "", map[string]interface{}{
		"cleanV2VConversionCache": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
