// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/kataras/golog"

	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// DeleteMetricTemplate 删除MetricTemplate
func (cli *ZSClient) DeleteMetricTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/zwatch/metrics/httpreceivers/templates/{uuid}", uuid, string(deleteMode))
}

