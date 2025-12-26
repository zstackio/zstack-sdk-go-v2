// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidatePriceUserConfig operates on ValidatePriceUserConfig
func (cli *ZSClient) ValidatePriceUserConfig(uuid string, params param.ValidatePriceUserConfigParam) (*view.ValidatePriceUserConfigEventView, error) {
	resp := view.ValidatePriceUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
