// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteSSORedirectTemplate deletes SSORedirectTemplate
func (cli *ZSClient) DeleteSSORedirectTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/delete/sso/redirect/template", uuid, string(deleteMode))
}
