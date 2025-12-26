// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// MatchModelServiceTemplateWithModel operates on MatchModelServiceTemplateWithModel
func (cli *ZSClient) MatchModelServiceTemplateWithModel(params param.MatchModelServiceTemplateWithModelParam) (*view.MatchModelServiceTemplateWithModelEventView, error) {
	resp := view.MatchModelServiceTemplateWithModelEventView{}
	if err := cli.Post("v1/ai/model-services/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
