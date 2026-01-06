// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ApplyDRSAdvice operates on DRSAdvice
func (cli *ZSClient) ApplyDRSAdvice(uuid string, params param.ApplyDRSAdviceParam) (*view.DRSAdviceInventoryView, error) {
	resp := view.DRSAdviceInventoryView{}
	if err := cli.Put("v1/clusters/drs/advice/{adviceUuid}/actions", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryDRSAdvice queries DRSAdvice list
func (cli *ZSClient) QueryDRSAdvice(params *param.QueryParam) ([]view.DRSAdviceInventoryView, error) {
	var resp []view.DRSAdviceInventoryView
	return resp, cli.List("v1/clusters/drs/advice", params, &resp)
}
