// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetLicenseAuthorizedCapacity gets LicenseAuthorizedCapacity by uuid
func (cli *ZSClient) GetLicenseAuthorizedCapacity(uuid string) (*view.LicenseAuthorizedCapacityInventoryView, error) {
	var resp view.LicenseAuthorizedCapacityInventoryView
	if err := cli.Get("v1/license-server/authorized-capacity", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
