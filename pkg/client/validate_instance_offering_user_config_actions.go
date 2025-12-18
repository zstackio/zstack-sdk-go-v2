// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidateInstanceOfferingUserConfig 操作ValidateInstanceOfferingUserConfig
func (cli *ZSClient) ValidateInstanceOfferingUserConfig(uuid string, params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

