// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// RemoveSNSWeComAtPerson removes SNSWeComAtPerson
func (cli *ZSClient) RemoveSNSWeComAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/we-com/{endpointUuid}/at-persons/{userId}", uuid, string(deleteMode))
}
