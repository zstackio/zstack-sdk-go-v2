// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteHuaweiIMasterVpc deletes HuaweiIMasterVpc
func (cli *ZSClient) DeleteHuaweiIMasterVpc(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sdn-controller/huawei-imaster/vpcs/{uuid}", uuid, string(deleteMode))
}
