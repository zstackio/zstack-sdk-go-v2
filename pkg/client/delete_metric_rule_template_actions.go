// Copyright (c) ZStack.io, Inc.

package client

import (
	"dev.zstack.io/ye.zou/zstack-go-sdk/pkg/param"
)

// DeleteMetricRuleTemplate deletes MetricRuleTemplate
func (cli *ZSClient) DeleteMetricRuleTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/monitortemplates/metricrules/{uuid}", uuid, string(deleteMode))
}
