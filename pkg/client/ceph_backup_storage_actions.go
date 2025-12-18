// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCephBackupStorage 查询CephBackupStorage列表
func (cli *ZSClient) QueryCephBackupStorage(params param.QueryParam) ([]view.QueryBackupStorageView, error) {
	var resp []view.QueryBackupStorageView
	return resp, cli.List("v1/backup-storage/ceph", &params, &resp)
}

