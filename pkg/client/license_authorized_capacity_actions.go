// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetLicenseAuthorizedCapacity gets LicenseAuthorizedCapacity by uuid
func (cli *ZSClient) GetLicenseAuthorizedCapacity() (*view.GetLicenseAuthorizedCapacityView, error) {
	var resp view.GetLicenseAuthorizedCapacityView
	if err := cli.GetWithRespKey("v1/license-server/authorized-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
