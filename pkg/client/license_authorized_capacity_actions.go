// Copyright (c) ZStack.io, Inc.

package client

import (
	"context"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// GetLicenseAuthorizedCapacity gets LicenseAuthorizedCapacity by uuid
func (cli *ZSClient) GetLicenseAuthorizedCapacity(ctx context.Context) (*view.LicenseAuthorizedCapacityInventoryView, error) {
	var resp view.LicenseAuthorizedCapacityInventoryView
	if err := cli.GetWithRespKey(ctx, "v1/license-server/authorized-capacity", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
