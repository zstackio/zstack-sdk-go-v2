// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTrainedModelRecord queries TrainedModelRecord list
func (cli *ZSClient) QueryTrainedModelRecord(params param.QueryParam) ([]view.TrainedModelRecordInventoryView, error) {
	var resp []view.TrainedModelRecordInventoryView
	return resp, cli.List("v1/ai/trained-model/records", &params, &resp)
}
