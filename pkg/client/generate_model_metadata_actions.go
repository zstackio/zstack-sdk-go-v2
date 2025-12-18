// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// GenerateModelMetadata 操作GenerateModelMetadata
func (cli *ZSClient) GenerateModelMetadata(uuid string, params param.GenerateModelMetadataParam) (*view.GenerateModelMetadataEventView, error) {
	resp := view.GenerateModelMetadataEventView{}
	if err := cli.Put("v1/ai/model/metadata/generate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

