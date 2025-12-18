// Copyright (c) ZStack.io, Inc.

package param

// UpdateImageDetailParam UpdateImage detail param
type UpdateImageDetailParam struct {
	Uuid string `json:"uuid" validate:"required"`
	Name string `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	GuestOsType string `json:"guestOsType,omitempty"`
	MediaType string `json:"mediaType,omitempty"`
	Format string `json:"format,omitempty"`
	System bool `json:"system,omitempty"`
	Platform string `json:"platform,omitempty"`
	Architecture string `json:"architecture,omitempty"`
	Virtio bool `json:"virtio,omitempty"`
}

// UpdateImageParam UpdateImage request param
type UpdateImageParam struct {
	BaseParam
	Params UpdateImageDetailParam `json:"params"`
}
