// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/view"
)

// BatchCreateIAM2VirtualIDFromConfigFile operates on BatchCreateIAM2VirtualIDFromConfigFile
func (cli *ZSClient) BatchCreateIAM2VirtualIDFromConfigFile(params param.BatchCreateIAM2VirtualIDFromConfigFileParam) (*view.BatchCreateIAM2VirtualIDFromConfigFileEventView, error) {
	resp := view.BatchCreateIAM2VirtualIDFromConfigFileEventView{}
	if err := cli.Post("v1/iam2/virtual-ids/from-file", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
