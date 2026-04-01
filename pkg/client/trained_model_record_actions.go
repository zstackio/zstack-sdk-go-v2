// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTrainedModelRecord queries TrainedModelRecord list
func (cli *ZSClient) QueryTrainedModelRecord(params *param.QueryParam) ([]view.TrainedModelRecordInventoryView, error) {
	var resp []view.TrainedModelRecordInventoryView
	return resp, cli.List("v1/ai/trained-model/records", params, &resp)
}

func (cli *ZSClient) GetTrainedModelRecord(uuid string) (*view.TrainedModelRecordInventoryView, error) {
	var resp view.TrainedModelRecordInventoryView
	if err := cli.Get("v1/ai/trained-model/records", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageTrainedModelRecord Pagination
func (cli *ZSClient) PageTrainedModelRecord(params *param.QueryParam) ([]view.TrainedModelRecordInventoryView, int, error) {
	var trainedModelRecords []view.TrainedModelRecordInventoryView
	total, err := cli.Page("v1/ai/trained-model/records", params, &trainedModelRecords)
	return trainedModelRecords, total, err
}
