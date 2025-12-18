// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// LongJobInventoryView LongJob
type LongJobInventoryView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	ApiId string `json:"apiId,omitempty"`
	JobName string `json:"jobName,omitempty"`
	JobData string `json:"jobData,omitempty"`
	JobResult string `json:"jobResult,omitempty"`
	State string `json:"state,omitempty"`
	TargetResourceUuid string `json:"targetResourceUuid,omitempty"`
	ManagementNodeUuid string `json:"managementNodeUuid,omitempty"`
	CreateDate time.Time `json:"createDate,omitempty"`
	LastOpDate time.Time `json:"lastOpDate,omitempty"`
	ExecuteTime int64 `json:"executeTime,omitempty"`
}

