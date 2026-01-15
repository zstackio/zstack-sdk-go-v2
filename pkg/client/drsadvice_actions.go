// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// ApplyDRSAdvice operates on DRSAdvice
func (cli *ZSClient) ApplyDRSAdvice(adviceUuid string, params param.ApplyDRSAdviceParam) (*view.DRSAdviceInventoryView, error) {
	resp := view.DRSAdviceInventoryView{}
	if err := cli.Put("v1/clusters/drs/advice", adviceUuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
// QueryDRSAdvice queries DRSAdvice list
func (cli *ZSClient) QueryDRSAdvice(params *param.QueryParam) ([]view.DRSAdviceInventoryView, error) {
	var resp []view.DRSAdviceInventoryView
	return resp, cli.List("v1/clusters/drs/advice", params, &resp)
}

func (cli *ZSClient) GetDRSAdvice(uuid string) (*view.DRSAdviceInventoryView, error) {
	var resp view.DRSAdviceInventoryView
	if err := cli.Get("v1/clusters/drs/advice", uuid, nil, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// PageDRSAdvice Pagination
func (cli *ZSClient) PageDRSAdvice(params *param.QueryParam) ([]view.DRSAdviceInventoryView, int, error) {
	var dRSAdvices []view.DRSAdviceInventoryView
	total, err := cli.Page("v1/clusters/drs/advice", params, &dRSAdvices)
	return dRSAdvices, total, err
}
