// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateCasClient 创建CasClient
func (cli *ZSClient) CreateCasClient(params param.CreateCasClientParam) (*view.CreateCasClientEventView, error) {
	resp := view.CreateCasClientEventView{}
	if err := cli.Post("v1/create/cas/client", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

