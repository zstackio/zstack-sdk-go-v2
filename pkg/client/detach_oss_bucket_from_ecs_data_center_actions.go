// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// DetachOssBucketFromEcsDataCenter operates on OssBucketFromEcsDataCenter
func (cli *ZSClient) DetachOssBucketFromEcsDataCenter(uuid string, params param.DetachOssBucketFromEcsDataCenterParam) (*view.DetachOssBucketFromEcsDataCenterEventView, error) {
	resp := view.DetachOssBucketFromEcsDataCenterEventView{}
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket/{ossBucketUuid}/detach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
