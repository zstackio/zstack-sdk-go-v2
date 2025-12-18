// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckElaborationContent 操作CheckElaborationContent
func (cli *ZSClient) CheckElaborationContent(params param.CheckElaborationContentParam) (*view.CheckElaborationContentView, error) {
	resp := view.CheckElaborationContentView{}
	if err := cli.Post("v1/errorcode/elaborations/check", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

