// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// QueryApplicationDevelopmentService queries ApplicationDevelopmentService list
func (cli *ZSClient) QueryApplicationDevelopmentService(params param.QueryParam) ([]view.ApplicationDevelopmentServiceInventoryView, error) {
	var resp []view.ApplicationDevelopmentServiceInventoryView
	return resp, cli.List("v1/ai/model-services/app/", &params, &resp)
}
