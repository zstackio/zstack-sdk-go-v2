// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryOssBucketFileName queries OssBucketFileName list
func (cli *ZSClient) QueryOssBucketFileName(params *param.QueryParam) ([]view.OssBucketInventoryView, error) {
	var resp []view.OssBucketInventoryView
	return resp, cli.List("v1/hybrid/aliyun/oss-bucket", params, &resp)
}
