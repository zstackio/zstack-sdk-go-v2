// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// BindModelToService operates on BindModelToService
func (cli *ZSClient) BindModelToService(params param.BindModelToServiceParam) (*view.BindModelToServiceEventView, error) {
	resp := view.BindModelToServiceEventView{}
	if err := cli.Post("v1/ai/models/{modelUuid}/model-services/{modelServiceUuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
