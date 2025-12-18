// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryCephOsdGroup queries CephOsdGroup list
func (cli *ZSClient) QueryCephOsdGroup(params param.QueryParam) ([]view.CephOsdGroupInventoryView, error) {
	var resp []view.CephOsdGroupInventoryView
	return resp, cli.List("v1/primary-storage/ceph/osdgroups", &params, &resp)
}
