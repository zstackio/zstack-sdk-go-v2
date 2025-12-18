// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryApplicationDevelopmentService 查询ApplicationDevelopmentService列表
func (cli *ZSClient) QueryApplicationDevelopmentService(params param.QueryParam) ([]view.QueryApplicationDevelopmentServiceView, error) {
	var resp []view.QueryApplicationDevelopmentServiceView
	return resp, cli.List("v1/ai/model-services/app/", &params, &resp)
}

