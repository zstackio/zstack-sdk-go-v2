// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryDRSAdvice 查询DRSAdvice列表
func (cli *ZSClient) QueryDRSAdvice(params param.QueryParam) ([]view.QueryDRSAdviceView, error) {
	var resp []view.QueryDRSAdviceView
	return resp, cli.List("v1/clusters/drs/advice", &params, &resp)
}

