// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// UpdateOssBucket updates OssBucket
func (cli *ZSClient) UpdateOssBucket(uuid string, params param.UpdateOssBucketParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/oss-bucket", uuid, "", map[string]interface{}{
		"updateOssBucket": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOssBucketFileRemote requires additional arguments.
//
// Deprecated: use DeleteOssBucketFileRemoteWithParams to supply the required fields.
func (cli *ZSClient) DeleteOssBucketFileRemote(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("DeleteOssBucketFileRemote requires additional arguments; use DeleteOssBucketFileRemoteWithParams")
}

// DeleteOssBucketFileRemoteWithParams sends the DELETE request with its operation parameters.
func (cli *ZSClient) DeleteOssBucketFileRemoteWithParams(uuid string, deleteMode param.DeleteMode, params param.DeleteOssBucketFileRemoteParamDetail) error {
	query := url.Values{
		"fileName": {params.FileName},
	}
	return cli.deleteResourceWithQuery("v1/hybrid/aliyun/oss-bucket-file/remote", uuid, deleteModeOrDefault(deleteMode, params.DeleteMode), query, "fileName")
}

// DeleteOssBucketRemote deletes OssBucketRemote
func (cli *ZSClient) DeleteOssBucketRemote(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket/remote", uuid, string(deleteMode))
}

// CreateOssBucketRemote creates OssBucketRemote
func (cli *ZSClient) CreateOssBucketRemote(params param.CreateOssBucketRemoteParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket/remote", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddOssBucketFromRemote adds OssBucketFromRemote
func (cli *ZSClient) AddOssBucketFromRemote(params param.AddOssBucketFromRemoteParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.Post("v1/hybrid/aliyun/oss-bucket", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachOssBucketToEcsDataCenter operates on OssBucketToEcsDataCenter
func (cli *ZSClient) AttachOssBucketToEcsDataCenter(ossBucketUuid string, params param.AttachOssBucketToEcsDataCenterParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/oss-bucket", ossBucketUuid, "", map[string]interface{}{
		"attachOssBucketToEcsDataCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOssBucketFileFromRemote gets OssBucketFileFromRemote by uuid
func (cli *ZSClient) GetOssBucketFileFromRemote() (*view.GetOssBucketFileFromRemoteView, error) {
	var resp view.GetOssBucketFileFromRemoteView
	if err := cli.GetWithRespKey("v1/hybrid/oss/file/remote", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryOssBucketFileName queries OssBucketFileName list
func (cli *ZSClient) QueryOssBucketFileName(params *param.QueryParam) ([]view.OssBucketInventoryView, error) {
	var resp []view.OssBucketInventoryView
	return resp, cli.List("v1/hybrid/aliyun/oss-bucket", params, &resp)
}

func (cli *ZSClient) GetOssBucketFileName(uuid string) (*view.OssBucketInventoryView, error) {
	var resp view.OssBucketInventoryView
	if err := cli.Get("v1/hybrid/aliyun/oss-bucket", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageOssBucketFileName Pagination
func (cli *ZSClient) PageOssBucketFileName(params *param.QueryParam) ([]view.OssBucketInventoryView, int, error) {
	var ossBucketFileNames []view.OssBucketInventoryView
	total, err := cli.Page("v1/hybrid/aliyun/oss-bucket", params, &ossBucketFileNames)
	return ossBucketFileNames, total, err
}

// DeleteOssBucketNameLocal deletes OssBucketNameLocal
func (cli *ZSClient) DeleteOssBucketNameLocal(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/hybrid/aliyun/oss-bucket", uuid, string(deleteMode))
}

// GetOssBucketNameFromRemote gets OssBucketNameFromRemote by uuid
func (cli *ZSClient) GetOssBucketNameFromRemote(uuid string) (*view.OssBucketPropertyView, error) {
	var resp view.OssBucketPropertyView
	if err := cli.GetWithRespKey("v1/hybrid/oss", uuid, "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachOssBucketFromEcsDataCenter operates on OssBucketFromEcsDataCenter
func (cli *ZSClient) DetachOssBucketFromEcsDataCenter(ossBucketUuid string, params param.DetachOssBucketFromEcsDataCenterParam) (*view.OssBucketInventoryView, error) {
	resp := view.OssBucketInventoryView{}
	if err := cli.PutWithRespKey("v1/hybrid/aliyun/oss-bucket", ossBucketUuid, "", map[string]interface{}{
		"detachOssBucketFromEcsDataCenter": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
