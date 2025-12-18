// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncEcsVSwitchFromRemote operates on SyncEcsVSwitchFromRemote
func (cli *ZSClient) SyncEcsVSwitchFromRemote(params param.SyncEcsVSwitchFromRemoteParam) (*view.SyncEcsVSwitchFromRemoteEventView, error) {
	resp := view.SyncEcsVSwitchFromRemoteEventView{}
	if err := cli.Post("v1/hybrid/aliyun/vswitch/{dataCenterUuid}/sync", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
