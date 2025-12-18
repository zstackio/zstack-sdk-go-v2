// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteSSORedirectTemplate 删除SSORedirectTemplate
func (cli *ZSClient) DeleteSSORedirectTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/sso/redirect/template", uuid, string(deleteMode))
}

