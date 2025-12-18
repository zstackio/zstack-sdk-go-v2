// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeAccessKeyState 操作AccessKeyState
func (cli *ZSClient) ChangeAccessKeyState(uuid string, params param.ChangeAccessKeyStateParam) (*view.ChangeAccessKeyStateEventView, error) {
	resp := view.ChangeAccessKeyStateEventView{}
	if err := cli.Put("v1/accesskeys/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

