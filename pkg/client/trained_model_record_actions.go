// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryTrainedModelRecord queries TrainedModelRecord list
func (cli *ZSClient) QueryTrainedModelRecord(params *param.QueryParam) ([]view.TrainedModelRecordInventoryView, error) {
	var resp []view.TrainedModelRecordInventoryView
	return resp, cli.List("v1/ai/trained-model/records", params, &resp)
}
