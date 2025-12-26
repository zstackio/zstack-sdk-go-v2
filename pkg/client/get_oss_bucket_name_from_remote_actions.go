// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetOssBucketNameFromRemote gets OssBucketNameFromRemote by uuid
func (cli *ZSClient) GetOssBucketNameFromRemote(uuid string) (*view.GetOssBucketNameFromRemoteView, error) {
	var resp view.GetOssBucketNameFromRemoteView
	if err := cli.Get("v1/hybrid/oss/{dataCenterUuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
