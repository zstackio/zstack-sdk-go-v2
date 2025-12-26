// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// AttachOssBucketToEcsDataCenter operates on OssBucketToEcsDataCenter
func (cli *ZSClient) AttachOssBucketToEcsDataCenter(uuid string, params param.AttachOssBucketToEcsDataCenterParam) (*view.AttachOssBucketToEcsDataCenterEventView, error) {
	resp := view.AttachOssBucketToEcsDataCenterEventView{}
	if err := cli.Put("v1/hybrid/aliyun/oss-bucket/{ossBucketUuid}/attach", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
