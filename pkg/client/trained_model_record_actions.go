// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryTrainedModelRecord 查询TrainedModelRecord列表
func (cli *ZSClient) QueryTrainedModelRecord(params param.QueryParam) ([]view.QueryTrainedModelRecordView, error) {
	var resp []view.QueryTrainedModelRecordView
	return resp, cli.List("v1/ai/trained-model/records", &params, &resp)
}

