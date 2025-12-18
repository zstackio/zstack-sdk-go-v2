// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeIPSecConnectionState 操作IPSecConnectionState
func (cli *ZSClient) ChangeIPSecConnectionState(uuid string, params param.ChangeIPSecConnectionStateParam) (*view.ChangeIPSecConnectionStateEventView, error) {
	resp := view.ChangeIPSecConnectionStateEventView{}
	if err := cli.Put("v1/ipsec/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

