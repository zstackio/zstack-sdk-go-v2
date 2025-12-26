// Copyright (c) ZStack.io, Inc.

package view

// GetUploadImageJobDetailsView GetUploadImageJobDetails
type GetUploadImageJobDetailsView struct {
	ExistingJobDetails []JobDetailsView `json:"existingJobDetails,omitempty"`
	Success bool `json:"success,omitempty"`
}

