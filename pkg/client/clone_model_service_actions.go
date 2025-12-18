// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CloneModelService operates on ModelService
func (cli *ZSClient) CloneModelService(params param.CloneModelServiceParam) (*view.CloneModelServiceEventView, error) {
	resp := view.CloneModelServiceEventView{}
	if err := cli.Post("v1/ai/model-services/{uuid}", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
