package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
)

// InputV2 input v2
// swagger:model InputV2
type InputV2 struct {

	// Unique, non-empty document identifier.
	ID string `json:"id,omitempty"`

	// text
	Text string `json:"text,omitempty"`
}

// Validate validates this input v2
func (m *InputV2) Validate(formats strfmt.Registry) error {
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}