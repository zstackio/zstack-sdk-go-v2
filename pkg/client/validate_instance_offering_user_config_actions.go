// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateInstanceOfferingUserConfig operates on ValidateInstanceOfferingUserConfig
func (cli *ZSClient) ValidateInstanceOfferingUserConfig(uuid string, params param.ValidateInstanceOfferingUserConfigParam) (*view.ValidateInstanceOfferingUserConfigEventView, error) {
	resp := view.ValidateInstanceOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
