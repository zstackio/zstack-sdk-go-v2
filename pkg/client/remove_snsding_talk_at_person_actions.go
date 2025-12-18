// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// RemoveSNSDingTalkAtPerson removes SNSDingTalkAtPerson
func (cli *ZSClient) RemoveSNSDingTalkAtPerson(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/application-endpoints/ding-talk/{endpointUuid}/at-persons/{phoneNumber}", uuid, string(deleteMode))
}
