// Copyright (c) ZStack.io, Inc.

package view

import "time"

// LongJobInventoryView LongJob
type LongJobInventoryView struct {
	rest string `json:"uuid,omitempty"`
	rest string `json:"name,omitempty"`
	rest string `json:"description,omitempty"`
	rest string `json:"apiId,omitempty"`
	rest string `json:"jobName,omitempty"`
	rest string `json:"jobData,omitempty"`
	rest string `json:"jobResult,omitempty"`
	rest string `json:"state,omitempty"`
	rest string `json:"targetResourceUuid,omitempty"`
	rest string `json:"managementNodeUuid,omitempty"`
	rest time.Time `json:"createDate,omitempty"`
	rest time.Time `json:"lastOpDate,omitempty"`
	rest int64 `json:"executeTime,omitempty"`
}

