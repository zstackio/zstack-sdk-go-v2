// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// UpdateOssBucket updates OssBucket
func (cli *ZSClient) UpdateOssBucket(uuid string, params param.UpdateOssBucketParam) (*view.UpdateOssBucketEventView, error) {
	resp := view.UpdateOssBucketEventView{}
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket/{uuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
