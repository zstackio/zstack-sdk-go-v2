// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateVxlanVtep creates VxlanVtep
func (cli *ZSClient) CreateVxlanVtep(params param.CreateVxlanVtepParam) (*view.CreateVxlanVtepEventView, error) {
	resp := view.CreateVxlanVtepEventView{}
	if err := cli.Post("v1/l2-networks/vxlan/vteps", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
