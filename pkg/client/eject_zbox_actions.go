// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// EjectZBox operates on EjectZBox
func (cli *ZSClient) EjectZBox(uuid string, params param.EjectZBoxParam) (*view.EjectZBoxEventView, error) {
	resp := view.EjectZBoxEventView{}
	if err := cli.Put("v1/zbox/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
