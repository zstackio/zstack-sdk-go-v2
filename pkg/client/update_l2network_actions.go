// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UpdateL2Network updates L2Network
func (cli *ZSClient) UpdateL2Network(uuid string, params param.UpdateL2NetworkParam) (*view.UpdateL2NetworkEventView, error) {
	resp := view.UpdateL2NetworkEventView{}
	if err := cli.Put("v1/l2-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
