// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// SyncConnectionAccessPointFromRemote operates on SyncConnectionAccessPointFromRemote
func (cli *ZSClient) SyncConnectionAccessPointFromRemote(uuid string, params param.SyncConnectionAccessPointFromRemoteParam) (*view.SyncConnectionAccessPointFromRemoteEventView, error) {
	resp := view.SyncConnectionAccessPointFromRemoteEventView{}
	if err := cli.Put("v1/hybrid/aliyun/access-point/{dataCenterUuid}/sync", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
