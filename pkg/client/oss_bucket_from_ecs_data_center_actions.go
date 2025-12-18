// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DetachOssBucketFromEcsDataCenter 操作OssBucketFromEcsDataCenter
func (cli *ZSClient) DetachOssBucketFromEcsDataCenter(uuid string, params param.DetachOssBucketFromEcsDataCenterParam) (*view.DetachOssBucketFromEcsDataCenterEventView, error) {
	resp := view.DetachOssBucketFromEcsDataCenterEventView{}
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket/{ossBucketUuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

