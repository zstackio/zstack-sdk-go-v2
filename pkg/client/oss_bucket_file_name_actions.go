// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryOssBucketFileName 查询OssBucketFileName列表
func (cli *ZSClient) QueryOssBucketFileName(params param.QueryParam) ([]view.QueryOssBucketFileNameView, error) {
	var resp []view.QueryOssBucketFileNameView
	return resp, cli.List("v1/hybrid/aliyun/oss-bucket", &params, &resp)
}

