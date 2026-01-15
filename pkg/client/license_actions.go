// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/licenses/mn", uuid, string(deleteMode))
}
// UpdateLicense updates License
func (cli *ZSClient) UpdateLicense(managementNodeUuid string, params param.UpdateLicenseParam) (*view.LicenseInventoryView, error) {
	resp := view.LicenseInventoryView{}
	if err := cli.Put("v1/licenses/mn", managementNodeUuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
