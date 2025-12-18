// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryPortMirrorSession 查询PortMirrorSession列表
func (cli *ZSClient) QueryPortMirrorSession(params param.QueryParam) ([]view.QueryPortMirrorSessionView, error) {
	var resp []view.QueryPortMirrorSessionView
	return resp, cli.List("v1/port-mirrors/sessions", &params, &resp)
}

