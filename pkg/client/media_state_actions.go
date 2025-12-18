// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ChangeMediaState 操作MediaState
func (cli *ZSClient) ChangeMediaState(uuid string, params param.ChangeMediaStateParam) (*view.ChangeMediaStateEventView, error) {
	resp := view.ChangeMediaStateEventView{}
	if err := cli.Put("v1/media/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

