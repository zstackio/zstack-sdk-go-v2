// Copyright (c) ZStack.io, Inc.

package view

import (
	"encoding/json"
	"fmt"
	"time"
)

// ZStackTime custom time type for ZStack's time format
type ZStackTime struct {
	time.Time
}

// UnmarshalJSON implements json.Unmarshaler
func (t *ZStackTime) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	if s == "" {
		return nil
	}
	// ZStack time format: "Oct 28, 2025 2:09:26 PM"
	formats := []string{
		"Jan 2, 2006 3:04:05 PM",
		"Jan 2, 2006 15:04:05",
		time.RFC3339,
	}
	for _, format := range formats {
		if parsed, err := time.Parse(format, s); err == nil {
			t.Time = parsed
			return nil
		}
	}
	return fmt.Errorf("cannot parse time: %s", s)
}

// MarshalJSON implements json.Marshaler
func (t ZStackTime) MarshalJSON() ([]byte, error) {
	if t.IsZero() {
		return []byte(`""`), nil
	}
	return json.Marshal(t.Format("Jan 2, 2006 3:04:05 PM"))
}

// BaseInfoView base info view
type BaseInfoView struct {
	UUID        string `json:"uuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

// BaseTimeView time info view
type BaseTimeView struct {
	CreateDate ZStackTime `json:"createDate"`
	LastOpDate ZStackTime `json:"lastOpDate"`
}

// BaseResourceView resource base view
type BaseResourceView struct {
	BaseInfoView
	BaseTimeView
}

// Generic wrapper types for APIs that return simple data types

// MapView wraps map return values
type MapView map[string]interface{}

// ListView wraps list/array return values
type ListView []interface{}

// StringView wraps string return values
type StringView string

// BooleanView wraps boolean return values
type BooleanView bool

// IntView wraps integer return values
type IntView int

// LongView wraps long integer return values
type LongView int64

// SuccessView represents successful operation with no data return
type SuccessView struct {
	Success bool `json:"success"`
}
