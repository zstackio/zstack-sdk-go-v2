// Copyright (c) ZStack.io, Inc.

package client

import (
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteBuildAppExportHistory sends the action with its default options.
func (cli *ZSClient) DeleteBuildAppExportHistory(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteBuildAppExportHistoryWithParams(uuid, deleteMode, param.DeleteBuildAppExportHistoryParamDetail{})
}

// DeleteBuildAppExportHistoryWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteBuildAppExportHistoryWithParams(buildAppUuid string, deleteMode param.DeleteMode, params param.DeleteBuildAppExportHistoryParamDetail) error {
	query := url.Values{}
	if params.ExportId != nil {
		query.Set("exportId", *params.ExportId)
	}
	if params.BuildSystemUuid != nil {
		query.Set("buildSystemUuid", *params.BuildSystemUuid)
	}
	return cli.deleteResourceWithQuery("v1/buildapp/exports", buildAppUuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}
// QueryBuildAppExportHistory queries BuildAppExportHistory list
func (cli *ZSClient) QueryBuildAppExportHistory(params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, error) {
	var resp []view.BuildAppExportHistoryInventoryView
	return resp, cli.List("v1/appcenter/exportapp", params, &resp)
}

func (cli *ZSClient) GetBuildAppExportHistory(uuid string) (*view.BuildAppExportHistoryInventoryView, error) {
	var resp view.BuildAppExportHistoryInventoryView
	if err := cli.Get("v1/appcenter/exportapp", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageBuildAppExportHistory Pagination
func (cli *ZSClient) PageBuildAppExportHistory(params *param.QueryParam) ([]view.BuildAppExportHistoryInventoryView, int, error) {
	var buildAppExportHistories []view.BuildAppExportHistoryInventoryView
	total, err := cli.Page("v1/appcenter/exportapp", params, &buildAppExportHistories)
	return buildAppExportHistories, total, err
}

// ExportBuildApp operates on BuildApp
func (cli *ZSClient) ExportBuildApp(uuid string, params param.ExportBuildAppParam) (*view.BuildAppExportHistoryInventoryView, error) {
	resp := view.BuildAppExportHistoryInventoryView{}
	if err := cli.PutWithRespKey("v1/appcenter/buildapp", uuid, "", map[string]interface{}{
		"exportBuildApp": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
