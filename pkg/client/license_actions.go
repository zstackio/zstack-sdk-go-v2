// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateLicense 更新License
func (cli *ZSClient) UpdateLicense(uuid string, params param.UpdateLicenseParam) (*view.UpdateLicenseEventView, error) {
	resp := view.UpdateLicenseEventView{}
	if err := cli.Put("v1/licenses/mn/{managementNodeUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

