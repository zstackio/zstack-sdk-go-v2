// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// ValidateDiskOfferingUserConfig operates on ValidateDiskOfferingUserConfig
func (cli *ZSClient) ValidateDiskOfferingUserConfig(uuid string, params param.ValidateDiskOfferingUserConfigParam) (*view.ValidateDiskOfferingUserConfigEventView, error) {
	resp := view.ValidateDiskOfferingUserConfigEventView{}
	if err := cli.Put("v1/billings/accounts/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
