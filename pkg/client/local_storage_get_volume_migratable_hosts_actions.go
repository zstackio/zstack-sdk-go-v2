// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// LocalStorageGetVolumeMigratableHosts 操作LocalStorageGetVolumeMigratableHosts
func (cli *ZSClient) LocalStorageGetVolumeMigratableHosts(params param.LocalStorageGetVolumeMigratableHostsParam) (*view.LocalStorageGetVolumeMigratableView, error) {
	var resp view.LocalStorageGetVolumeMigratableView
	if err := cli.Get("v1/volumes/{volumeUuid}/migration-target-hosts", "", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

