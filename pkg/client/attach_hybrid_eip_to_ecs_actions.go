// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachHybridEipToEcs operates on HybridEipToEcs
func (cli *ZSClient) AttachHybridEipToEcs(params param.AttachHybridEipToEcsParam) (*view.AttachHybridEipToEcsEventView, error) {
	resp := view.AttachHybridEipToEcsEventView{}
	if err := cli.Post("v1/hybrid/eip/{eipUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
