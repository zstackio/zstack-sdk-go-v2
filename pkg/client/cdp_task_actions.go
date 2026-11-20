// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"
	"strconv"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// CreateCdpTask creates CdpTask
func (cli *ZSClient) CreateCdpTask(params param.CreateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Post("v1/cdp-backup-storage/task", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryCdpTask queries CdpTask list
func (cli *ZSClient) QueryCdpTask(params *param.QueryParam) ([]view.CdpTaskInventoryView, error) {
	var resp []view.CdpTaskInventoryView
	return resp, cli.List("v1/cdp-task", params, &resp)
}

func (cli *ZSClient) GetCdpTask(uuid string) (*view.CdpTaskInventoryView, error) {
	var resp view.CdpTaskInventoryView
	if err := cli.Get("v1/cdp-task", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageCdpTask Pagination
func (cli *ZSClient) PageCdpTask(params *param.QueryParam) ([]view.CdpTaskInventoryView, int, error) {
	var cdpTasks []view.CdpTaskInventoryView
	total, err := cli.Page("v1/cdp-task", params, &cdpTasks)
	return cdpTasks, total, err
}
// UpdateCdpTask updates CdpTask
func (cli *ZSClient) UpdateCdpTask(uuid string, params param.UpdateCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.PutWithRespKey("v1/cdp-backup-storage/task", uuid, "", map[string]interface{}{
		"updateCdpTask": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// DeleteCdpTask sends the action with its default options.
func (cli *ZSClient) DeleteCdpTask(uuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteCdpTaskWithParams(uuid, deleteMode, param.DeleteCdpTaskParamDetail{})
}

// DeleteCdpTaskWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteCdpTaskWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteCdpTaskParamDetail) error {
	query := url.Values{}
	if params.Force != nil {
		query.Set("force", strconv.FormatBool(*params.Force))
	}
	return cli.deleteResourceWithQuery("v1/cdp-task", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query)
}

// UpgradeBackupStorageCdpTasks operates on UpgradeBackupStorageCdpTasks
func (cli *ZSClient) UpgradeBackupStorageCdpTasks(backupStorageUuid string, params param.UpgradeBackupStorageCdpTasksParam) (*view.UpgradeBackupStorageCdpTasksEventView, error) {
	resp := view.UpgradeBackupStorageCdpTasksEventView{}
	if err := cli.PutWithRespKey("v1/cdp-task/upgrade", backupStorageUuid, "", map[string]interface{}{
		"upgradeBackupStorageCdpTasks": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCdpTaskData deletes CdpTaskData
func (cli *ZSClient) DeleteCdpTaskData(uuid string, params param.DeleteCdpTaskDataParam) (*view.DeleteCdpTaskDataEventView, error) {
	resp := view.DeleteCdpTaskDataEventView{}
	if err := cli.Post(fmt.Sprintf("v1/cdp-task/%s/data", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableCdpTask operates on DisableCdpTask
func (cli *ZSClient) DisableCdpTask(uuid string, params param.DisableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/cdp-task/disable/%s", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableCdpTask operates on EnableCdpTask
func (cli *ZSClient) EnableCdpTask(uuid string, params param.EnableCdpTaskParam) (*view.CdpTaskInventoryView, error) {
	resp := view.CdpTaskInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/cdp-task/enable/%s", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableCdpTaskAsync Async
func (cli *ZSClient) EnableCdpTaskAsync(uuid string, params param.EnableCdpTaskParam) (string, error) {

	resource := fmt.Sprintf("v1/cdp-task/enable/%s", uuid)
	responseKey := ""
	var retVal interface{}

	apiId, err := cli.PostWithAsync(resource, responseKey, params, retVal, true)
	if err != nil {
		return "", err
	}

	return apiId, nil
}
