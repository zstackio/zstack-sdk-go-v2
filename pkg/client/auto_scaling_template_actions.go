// Copyright (c) ZStack.io, Inc.

package client

import (
	"fmt"

	"github.com/zstackio/zstack-sdk-go-v2/pkg/param"
	"github.com/zstackio/zstack-sdk-go-v2/pkg/view"
)

var _ = param.BaseParam{} // avoid unused import
var _ = view.MapView{} // avoid unused import

// DeleteAutoScalingTemplate deletes AutoScalingTemplate
func (cli *ZSClient) DeleteAutoScalingTemplate(uuid string, deleteMode param.DeleteMode) error {
	return cli.Delete("v1/autoscaling/template", uuid, string(deleteMode))
}

// AttachAutoScalingTemplateToGroup operates on AutoScalingTemplateToGroup
func (cli *ZSClient) AttachAutoScalingTemplateToGroup(uuid string, groupUuid string, params param.AttachAutoScalingTemplateToGroupParam) (*view.AutoScalingGroupInventoryView, error) {
	resp := view.AutoScalingGroupInventoryView{}
	if err := cli.Post(fmt.Sprintf("v1/autoscaling/template/%s/groups/%s", uuid, groupUuid), params, &resp); err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachAutoScalingTemplateFromGroup operates on AutoScalingTemplateFromGroup
func (cli *ZSClient) DetachAutoScalingTemplateFromGroup(templateUuid string, groupUuid string, deleteMode param.DeleteMode) error {
	return cli.DeleteWithSpec("v1/autoscaling/template", templateUuid, fmt.Sprintf("groups/%s", groupUuid), fmt.Sprintf("deleteMode=%s", deleteMode), nil)
}
