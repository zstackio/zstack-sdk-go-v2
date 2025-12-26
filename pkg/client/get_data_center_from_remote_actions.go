// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GetDataCenterFromRemote gets DataCenterFromRemote by uuid
func (cli *ZSClient) GetDataCenterFromRemote(uuid string) (*view.GetDataCenterFromRemoteView, error) {
	var resp view.GetDataCenterFromRemoteView
	if err := cli.Get("v1/hybrid/data-center/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
