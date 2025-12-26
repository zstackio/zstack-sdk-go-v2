// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// QueryDRSAdvice queries DRSAdvice list
func (cli *ZSClient) QueryDRSAdvice(params *param.QueryParam) ([]view.DRSAdviceInventoryView, error) {
	var resp []view.DRSAdviceInventoryView
	return resp, cli.List("v1/clusters/drs/advice", params, &resp)
}
