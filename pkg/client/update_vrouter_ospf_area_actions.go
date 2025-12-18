// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateVRouterOspfArea updates VRouterOspfArea
func (cli *ZSClient) UpdateVRouterOspfArea(uuid string, params param.UpdateVRouterOspfAreaParam) (*view.UpdateVRouterOspfAreaEventView, error) {
	resp := view.UpdateVRouterOspfAreaEventView{}
	if err := cli.Put("v1/routerArea/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
