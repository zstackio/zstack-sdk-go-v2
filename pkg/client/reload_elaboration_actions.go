// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// ReloadElaboration 操作ReloadElaboration
func (cli *ZSClient) ReloadElaboration(uuid string, params param.ReloadElaborationParam) (*view.ReloadElaborationEventView, error) {
	resp := view.ReloadElaborationEventView{}
	if err := cli.Put("v1/errorcode/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

