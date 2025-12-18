// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CheckIAM2VirtualIDConfigFile 操作CheckIAM2VirtualIDConfigFile
func (cli *ZSClient) CheckIAM2VirtualIDConfigFile(uuid string, params param.CheckIAM2VirtualIDConfigFileParam) (*view.CheckIAM2VirtualIDConfigFileView, error) {
	resp := view.CheckIAM2VirtualIDConfigFileView{}
	if err := cli.Put("v1/iam2/virtual-ids/from-file", uuid, params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

