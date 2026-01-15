// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryCephOsdGroup queries CephOsdGroup list
func (cli *ZSClient) QueryCephOsdGroup(params *param.QueryParam) ([]view.CephOsdGroupInventoryView, error) {
	var resp []view.CephOsdGroupInventoryView
	return resp, cli.List("v1/primary-storage/ceph/osdgroups", params, &resp)
}

// PageCephOsdGroup Pagination
func (cli *ZSClient) PageCephOsdGroup(params *param.QueryParam) ([]view.CephOsdGroupInventoryView, int, error) {
	var cephOsdGroups []view.CephOsdGroupInventoryView
	total, err := cli.Page("v1/primary-storage/ceph/osdgroups", params, &cephOsdGroups)
	return cephOsdGroups, total, err
}
