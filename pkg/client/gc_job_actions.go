// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

// TriggerGCJob operates on TriggerGCJob
func (cli *ZSClient) TriggerGCJob(uuid string, params param.TriggerGCJobParam) (*view.TriggerGCJobEventView, error) {
	resp := view.TriggerGCJobEventView{}
	if err := cli.PutWithRespKey("v1/gc-jobs", uuid, "", map[string]interface{}{
		"triggerGCJob": params.Params,
	}, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGCJob deletes GCJob
func (cli *ZSClient) DeleteGCJob(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/gc-jobs", uuid, string(deleteMode))
}

// QueryGCJob queries GCJob list
func (cli *ZSClient) QueryGCJob(params *param.QueryParam) ([]view.GarbageCollectorInventoryView, error) {
	var resp []view.GarbageCollectorInventoryView
	return resp, cli.List("v1/gc-jobs", params, &resp)
}

func (cli *ZSClient) GetGCJob(uuid string) (*view.GarbageCollectorInventoryView, error) {
	var resp view.GarbageCollectorInventoryView
	if err := cli.Get("v1/gc-jobs", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageGCJob Pagination
func (cli *ZSClient) PageGCJob(params *param.QueryParam) ([]view.GarbageCollectorInventoryView, int, error) {
	var gCJobs []view.GarbageCollectorInventoryView
	total, err := cli.Page("v1/gc-jobs", params, &gCJobs)
	return gCJobs, total, err
}
