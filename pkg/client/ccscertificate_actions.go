// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteCCSCertificate 删除CCSCertificate
func (cli *ZSClient) DeleteCCSCertificate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/crypto/ccs-certificate/delete/{uuid}", uuid, string(deleteMode))
}

