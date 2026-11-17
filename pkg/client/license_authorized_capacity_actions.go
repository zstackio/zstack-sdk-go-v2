// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetLicenseAuthorizedCapacity gets LicenseAuthorizedCapacity by uuid
func (cli *ZSClient) GetLicenseAuthorizedCapacity() (*view.LicenseAuthorizedCapacityInventoryView, error) {
	var resp view.LicenseAuthorizedCapacityInventoryView
	if err := cli.GetWithRespKey("v1/license-server/authorized-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncLicenseCapacity operates on LicenseCapacity
func (cli *ZSClient) SyncLicenseCapacity() (*view.SyncLicenseCapacityEventView, error) {
	resp := view.SyncLicenseCapacityEventView{}
	if err := cli.PutWithRespKey("v1/license-server/authorized-capacity/sync", "", "", map[string]interface{}{
		"syncLicenseCapacity": map[string]interface{}{},
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RequestLicenseCapacity operates on RequestLicenseCapacity
func (cli *ZSClient) RequestLicenseCapacity(params param.RequestLicenseCapacityParam) (*view.LicenseAuthorizedCapacityInventoryView, error) {
	resp := view.LicenseAuthorizedCapacityInventoryView{}
	if err := cli.Post("v1/license-server/capacity-application", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
