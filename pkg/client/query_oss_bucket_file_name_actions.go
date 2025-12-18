// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryOssBucketFileName queries OssBucketFileName list
func (cli *ZSClient) QueryOssBucketFileName(params param.QueryParam) ([]view.OssBucketInventoryView, error) {
	var resp []view.OssBucketInventoryView
	return resp, cli.List("v1/hybrid/aliyun/oss-bucket", &params, &resp)
}
