// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// LocalStorageGetVolumeMigratableHosts operates on LocalStorageGetVolumeMigratableHosts
func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(params param.LocalStorageGetVolumeMigratableHostsParam) (*view.LocalStorageGetVolumeMigratableView, error) {
	var resp view.LocalStorageGetVolumeMigratableView
	if err := cli.Get("v1/volumes/{volumeUuid}/migration-target-hosts", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
