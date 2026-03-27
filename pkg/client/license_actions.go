// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteLicense deletes License
func (cli *ZSClient) DeleteLicense(ctx context.Context, uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete(ctx, "v1/licenses/mn", uuid, string(deleteMode))
}
// UpdateLicense updates License
func (cli *ZSClient) UpdateLicense(ctx context.Context, managementNodeUuid string, params param.UpdateLicenseParam) (*view.LicenseInventoryView, error) {
	resp := view.LicenseInventoryView{}
	if err := cli.PutWithRespKey(ctx, "v1/licenses/mn", managementNodeUuid, "", map[string]interface{}{
		"updateLicense": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
