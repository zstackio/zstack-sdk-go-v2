// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetOssBucketNameFromRemote gets OssBucketNameFromRemote by uuid
func (cli *ZSClient) GetOssBucketNameFromRemote(uuid string) (*view.GetOssBucketNameFromRemoteView, error) {
	var resp view.GetOssBucketNameFromRemoteView
	if err := cli.Get("v1/hybrid/oss/{dataCenterUuid}/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
