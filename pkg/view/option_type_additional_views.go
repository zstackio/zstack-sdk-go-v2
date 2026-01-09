// Copyright (c) ZStack.io, Inc.

package view

import "time"

var _ = time.Now // avoid unused import

// OptionTypeView OptionType
type OptionTypeView struct {
	Uuid string `json:"uuid,omitempty"`
	Name string `json:"name,omitempty"`
	Code *string `json:"code,omitempty"`
	Category *string `json:"category,omitempty"`
	Required *bool `json:"required,omitempty"`
	Editable *bool `json:"editable,omitempty"`
	Enabled *bool `json:"enabled,omitempty"`
	DisplayOrder *int `json:"displayOrder,omitempty"`
	InputType string `json:"inputType,omitempty"`
	PlaceHolderText *string `json:"placeHolderText,omitempty"`
	DefaultValue *string `json:"defaultValue,omitempty"`
	NoSelection *string `json:"noSelection,omitempty"`
	NoBlank *bool `json:"noBlank,omitempty"`
	SecretField *bool `json:"secretField,omitempty"`
	MinVal *int64 `json:"minVal,omitempty"`
	MaxVal *int64 `json:"maxVal,omitempty"`
	MinLength *int64 `json:"minLength,omitempty"`
	MaxLength *int64 `json:"maxLength,omitempty"`
	FieldContext *string `json:"fieldContext,omitempty"`
	FieldClass *string `json:"fieldClass,omitempty"`
	FieldLabel *string `json:"fieldLabel,omitempty"`
	FieldCode *string `json:"fieldCode,omitempty"`
	FieldName *string `json:"fieldName,omitempty"`
	FieldGetName *string `json:"fieldGetName,omitempty"`
	FieldSetName *string `json:"fieldSetName,omitempty"`
	FieldGetContext *string `json:"fieldGetContext,omitempty"`
	FieldSetContext *string `json:"fieldSetContext,omitempty"`
	FieldGroup *string `json:"fieldGroup,omitempty"`
	FieldGroupI18nCode *string `json:"fieldGroupI18nCode,omitempty"`
	HelpText *string `json:"helpText,omitempty"`
	HelpTextI18nCode *string `json:"helpTextI18nCode,omitempty"`
	OptionSourceType *string `json:"optionSourceType,omitempty"`
	OptionSource *string `json:"optionSource,omitempty"`
	DependsOn *string `json:"dependsOn,omitempty"`
	ShowOnEdit *bool `json:"showOnEdit,omitempty"`
	DisplayValueOnDetails *bool `json:"displayValueOnDetails,omitempty"`
	ShowOnCreate *bool `json:"showOnCreate,omitempty"`
	VerifyPattern *string `json:"verifyPattern,omitempty"`
}

