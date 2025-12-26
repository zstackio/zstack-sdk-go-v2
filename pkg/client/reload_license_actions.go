// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ReloadLicense operates on ReloadLicense
func (cli *ZSClient) ReloadLicense(uuid string, params param.ReloadLicenseParam) (*view.ReloadLicenseView, error) {
	resp := view.ReloadLicenseView{}
	if err := cli.Put("v1/licenses/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
