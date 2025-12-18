// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteV2VConversionHost deletes V2VConversionHost
func (cli *ZSClient) DeleteV2VConversionHost(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/v2v-conversion-hosts/{uuid}", uuid, string(deleteMode))
}
