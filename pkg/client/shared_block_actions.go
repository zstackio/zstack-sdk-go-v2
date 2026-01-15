// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QuerySharedBlock queries SharedBlock list
func (cli *ZSClient) QuerySharedBlock(params *param.QueryParam) ([]view.SharedBlockInventoryView, error) {
	var resp []view.SharedBlockInventoryView
	return resp, cli.List("v1/sharedblock-group/sharedblocks", params, &resp)
}

// PageSharedBlock Pagination
func (cli *ZSClient) PageSharedBlock(params *param.QueryParam) ([]view.SharedBlockInventoryView, int, error) {
	var sharedBlocks []view.SharedBlockInventoryView
	total, err := cli.Page("v1/sharedblock-group/sharedblocks", params, &sharedBlocks)
	return sharedBlocks, total, err
}
// UpdateSharedBlock updates SharedBlock
func (cli *ZSClient) UpdateSharedBlock(sharedBlockGroupUuid string, uuid string, params param.UpdateSharedBlockParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	err := cli.PutWithSpec("v1/primary-storage/sharedblockgroup", sharedBlockGroupUuid, fmt.Sprintf("sharedblocks/%s/actions", uuid), "", params, &resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
