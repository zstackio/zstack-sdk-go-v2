// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachProvisionNicToBonding operates on ProvisionNicToBonding
func (cli *ZSClient) AttachProvisionNicToBonding(params param.AttachProvisionNicToBondingParam) (*view.AttachProvisionNicToBondingEventView, error) {
	resp := view.AttachProvisionNicToBondingEventView{}
	if err := cli.Post("v1/baremetal2/bm-instances/{uuid}/bm2-bondings/{bondingUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
