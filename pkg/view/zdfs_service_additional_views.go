// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// ZdfsServiceView ZdfsService
type ZdfsServiceView struct {
	Uuid string `json:"uuid,omitempty"`
	SlaveCount int `json:"slaveCount,omitempty"`
	SentinelCount int `json:"sentinelCount,omitempty"`
	MetaServerStatus *string `json:"metaServerStatus,omitempty"`
	MetaServers []MetaServerServiceView `json:"metaServers,omitempty"`
}

