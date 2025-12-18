// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeEipState changes EipState
func (cli *ZSClient) ChangeEipState(uuid string, params param.ChangeEipStateParam) (*view.ChangeEipStateEventView, error) {
	resp := view.ChangeEipStateEventView{}
	if err := cli.Put("v1/eips/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
