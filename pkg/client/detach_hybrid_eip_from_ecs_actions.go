// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachHybridEipFromEcs operates on HybridEipFromEcs
func (cli *ZSClient) DetachHybridEipFromEcs(params param.DetachHybridEipFromEcsParam) (*view.DetachHybridEipFromEcsEventView, error) {
	resp := view.DetachHybridEipFromEcsEventView{}
	if err := cli.Post("v1/hybrid/eip/{eipUuid}/detach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
