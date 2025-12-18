// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// UnsubscribeSNSTopic 操作UnsubscribeSNSTopic
func (cli *ZSClient) UnsubscribeSNSTopic(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/sns/topics/{topicUuid}/endpoints/{endpointUuid}", uuid, string(deleteMode))
}

