// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// RequestLicenseCapacity operates on RequestLicenseCapacity
func (cli *ZSClient) RequestLicenseCapacity(params param.RequestLicenseCapacityParam) (*view.RequestLicenseCapacityEventView, error) {
	resp := view.RequestLicenseCapacityEventView{}
	if err := cli.Post("v1/license-server/capacity-application", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
