// Copyright (c) ZStack.io, Inc.

package client

import (
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/param"
	"github.com/terraform-zstack-modules/zstack-sdk-go/pkg/view"
)

// CreateMetricTemplate creates MetricTemplate
func (cli *ZSClient) CreateMetricTemplate(params param.CreateMetricTemplateParam) (*view.CreateMetricTemplateEventView, error) {
	resp := view.CreateMetricTemplateEventView{}
	if err := cli.Post("v1/zwatch/metrics/httpreceivers/templates", params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}
