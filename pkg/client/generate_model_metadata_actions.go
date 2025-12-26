// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// GenerateModelMetadata operates on GenerateModelMetadata
func (cli *ZSClient) GenerateModelMetadata(uuid string, params param.GenerateModelMetadataParam) (*view.GenerateModelMetadataEventView, error) {
	resp := view.GenerateModelMetadataEventView{}
	if err := cli.Put("v1/ai/model/metadata/generate", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
