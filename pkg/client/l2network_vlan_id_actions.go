// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeL2NetworkVlanId 操作L2NetworkVlanId
func (cli *ZSClient) ChangeL2NetworkVlanId(uuid string, params param.ChangeL2NetworkVlanIdParam) (*view.ChangeL2NetworkVlanIdEventView, error) {
	resp := view.ChangeL2NetworkVlanIdEventView{}
	if err := cli.Put("v1/l2-networks/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

