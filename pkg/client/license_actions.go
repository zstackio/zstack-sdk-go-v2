// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(managementNodeUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/licenses/mn", fmt.Sprintf(\"%s/actions\", managementNodeUuid), string(deleteMode))
}
// UpdateLicense updates License
func (cli *ZSClient) UpdateLicense(managementNodeUuid string, params param.UpdateLicenseParam) (*view.LicenseInventoryView, error) {
	var resp view.UpdateLicenseEventView
	err := cli.PutWithSpec("v1/licenses/mn", fmt.Sprintf(\"%s/actions\", managementNodeUuid), params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp.Inventory, nil
}
