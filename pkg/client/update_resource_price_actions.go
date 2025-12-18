// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateResourcePrice updates ResourcePrice
func (cli *ZSClient) UpdateResourcePrice(uuid string, params param.UpdateResourcePriceParam) (*view.UpdateResourcePriceEventView, error) {
	resp := view.UpdateResourcePriceEventView{}
	if err := cli.Put("v1/billings/prices/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
