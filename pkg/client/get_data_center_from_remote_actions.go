// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GetDataCenterFromRemote gets DataCenterFromRemote by uuid
func (cli *ZSClient) GetDataCenterFromRemote(uuid string) (*view.GetDataCenterFromRemoteView, error) {
	var resp view.GetDataCenterFromRemoteView
	if err := cli.Get("v1/hybrid/data-center/remote", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
