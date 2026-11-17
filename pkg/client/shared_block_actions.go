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

func (cli *ZSClient) GetSharedBlock(uuid string) (*view.SharedBlockInventoryView, error) {
	var resp view.SharedBlockInventoryView
	if err := cli.Get("v1/sharedblock-group/sharedblocks", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
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

// AddSharedBlockToSharedBlockGroup adds SharedBlockToSharedBlockGroup
func (cli *ZSClient) AddSharedBlockToSharedBlockGroup(uuid string, params param.AddSharedBlockToSharedBlockGroupParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/primary-storage/sharedblockgroup/%s/sharedblocks", uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSharedBlockCandidate gets SharedBlockCandidate by uuid
func (cli *ZSClient) GetSharedBlockCandidate() (*view.GetSharedBlockCandidateView, error) {
	var resp view.GetSharedBlockCandidateView
	if err := cli.GetWithRespKey("v1/primary-storage/sharedblockgroup/sharedblock-candidates", "", "", nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefreshSharedblockDeviceCapacity operates on SharedblockDeviceCapacity
func (cli *ZSClient) RefreshSharedblockDeviceCapacity(sharedBlockGroupUuid string, uuid string, params param.RefreshSharedblockDeviceCapacityParam) (*view.SharedBlockGroupPrimaryStorageInventoryView, error) {
	resp := view.SharedBlockGroupPrimaryStorageInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/primary-storage/sharedblockgroup/%s/sharedblocks/%s", sharedBlockGroupUuid, uuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
