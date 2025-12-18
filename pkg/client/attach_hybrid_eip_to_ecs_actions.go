// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// AttachHybridEipToEcs operates on HybridEipToEcs
func (cli *ZSClient) AttachHybridEipToEcs(params param.AttachHybridEipToEcsParam) (*view.AttachHybridEipToEcsEventView, error) {
	resp := view.AttachHybridEipToEcsEventView{}
	if err := cli.Post("v1/hybrid/eip/{eipUuid}/attach", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
