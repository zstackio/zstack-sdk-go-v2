// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"net/url"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{}    // avoid unused import

// UpdateDirectory updates Directory
func (cli *ZSClient) UpdateDirectory(params param.UpdateDirectoryParam) (*view.DirectoryInventoryView, error) {
	resp := view.DirectoryInventoryView{}
	if err := cli.PutWithRespKey("v1/update/directory", "", "", map[string]interface{}{
		"updateDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDirectory creates Directory
func (cli *ZSClient) CreateDirectory(params param.CreateDirectoryParam) (*view.DirectoryInventoryView, error) {
	resp := view.DirectoryInventoryView{}
	if err := cli.Post("v1/create/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// QueryDirectory queries Directory list
func (cli *ZSClient) QueryDirectory(params *param.QueryParam) ([]view.DirectoryInventoryView, error) {
	var resp []view.DirectoryInventoryView
	return resp, cli.List("v1/directories", params, &resp)
}

func (cli *ZSClient) GetDirectory(uuid string) (*view.DirectoryInventoryView, error) {
	var resp view.DirectoryInventoryView
	if err := cli.Get("v1/directories", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDirectory Pagination
func (cli *ZSClient) PageDirectory(params *param.QueryParam) ([]view.DirectoryInventoryView, int, error) {
	var directories []view.DirectoryInventoryView
	total, err := cli.Page("v1/directories", params, &directories)
	return directories, total, err
}

// DeleteDirectory deletes Directory
func (cli *ZSClient) DeleteDirectory(uuid string, deleteMode param.DeleteMode) error {
	mode := string(deleteMode)
	return cli.DeleteWithBody("v1/delete/directory", param.DeleteDirectoryParam{
		Params: param.DeleteDirectoryParamDetail{
			Uuid:       uuid,
			DeleteMode: &mode,
		},
	}, nil)
}

// AddResourcesToDirectory adds ResourcesToDirectory
func (cli *ZSClient) AddResourcesToDirectory(params param.AddResourcesToDirectoryParam) (*view.AddResourcesToDirectoryEventView, error) {
	resp := view.AddResourcesToDirectoryEventView{}
	if err := cli.Post("v1/add/resources/directory", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveResourcesFromDirectory requires additional arguments.
//
// Deprecated: use RemoveResourcesFromDirectoryWithParams to supply the required fields.
func (cli *ZSClient) RemoveResourcesFromDirectory(uuid string, deleteMode param.DeleteMode) error {
	return fmt.Errorf("RemoveResourcesFromDirectory requires additional arguments; use RemoveResourcesFromDirectoryWithParams")
}

// RemoveResourcesFromDirectoryWithParams sends the required route and query parameters for RemoveResourcesFromDirectory.
func (cli *ZSClient) RemoveResourcesFromDirectoryWithParams(deleteMode param.DeleteMode, params param.RemoveResourcesFromDirectoryParamDetail) error {
	query := url.Values{
		"directoryUuid": {params.DirectoryUuid},
		"resourceUuids": params.ResourceUuids,
	}
	return cli.deleteActionWithQuery("v1/remove/resources/directory", "", "", deleteMode, query, "directoryUuid", "resourceUuids")
}

// MoveDirectory operates on MoveDirectory
func (cli *ZSClient) MoveDirectory(params param.MoveDirectoryParam) (*view.MoveDirectoryEventView, error) {
	resp := view.MoveDirectoryEventView{}
	if err := cli.PutWithRespKey("v1/move/directory", "", "", map[string]interface{}{
		"moveDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveResourcesToDirectory operates on MoveResourcesToDirectory
func (cli *ZSClient) MoveResourcesToDirectory(params param.MoveResourcesToDirectoryParam) (*view.MoveResourcesToDirectoryEventView, error) {
	resp := view.MoveResourcesToDirectoryEventView{}
	if err := cli.PutWithRespKey("v1/move/resources/directory", "", "", map[string]interface{}{
		"moveResourcesToDirectory": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
