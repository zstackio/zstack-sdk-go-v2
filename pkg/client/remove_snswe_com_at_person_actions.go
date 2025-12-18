// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSNSWeComAtPerson 操作RemoveSNSWeComAtPerson
func (cli *ZSClient) RemoveSNSWeComAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/we-com/{endpointUuid}/at-persons/{userId}", uuid, string(deleteMode))
}

