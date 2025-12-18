// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ValidatePriceUserConfig 操作ValidatePriceUserConfig
func (cli *ZSClient) ValidatePriceUserConfig(uuid string, params param.ValidatePriceUserConfigParam) (*view.ValidatePriceUserConfigEventView, error) {
	resp := view.ValidatePriceUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

