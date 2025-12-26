// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateLicense updates License
func (cli *ZSClient) UpdateLicense(uuid string, params param.UpdateLicenseParam) (*view.UpdateLicenseEventView, error) {
	resp := view.UpdateLicenseEventView{}
	if err := cli.Put("v1/licenses/mn/{managementNodeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
