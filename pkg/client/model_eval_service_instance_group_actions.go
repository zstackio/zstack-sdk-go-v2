// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// QueryModelEvalServiceInstanceGroup queries ModelEvalServiceInstanceGroup list
func (cli *ZSClient) QueryModelEvalServiceInstanceGroup(params *param.QueryParam) ([]view.ModelServiceInstanceGroupInventoryView, error) {
	var resp []view.ModelServiceInstanceGroupInventoryView
	return resp, cli.List("v1/ai/model-eval-services/instances/groups/", params, &resp)
}

func (cli *ZSClient) GetModelEvalServiceInstanceGroup(uuid string) (*view.ModelServiceInstanceGroupInventoryView, error) {
	var resp view.ModelServiceInstanceGroupInventoryView
	if err := cli.Get("v1/ai/model-eval-services/instances/groups/", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
