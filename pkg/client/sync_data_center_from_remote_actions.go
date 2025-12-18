// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncDataCenterFromRemote operates on SyncDataCenterFromRemote
func (cli *ZSClient) SyncDataCenterFromRemote(params param.SyncDataCenterFromRemoteParam) (*view.SyncDataCenterFromRemoteEventView, error) {
	var resp view.SyncDataCenterFromRemoteEventView
	if err := cli.Get("v1/hybrid/data-center/{uuid}/sync", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
