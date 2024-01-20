// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright 2021 The Sigstore Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// ProposedEntry proposed entry
//
// swagger:discriminator ProposedEntry kind
type ProposedEntry interface {
	runtime.Validatable
	runtime.ContextValidatable

	// kind
	// Required: true
	Kind() string
	SetKind(string)

	// AdditionalProperties in base type shoud be handled just like regular properties
	// At this moment, the base type property is pushed down to the subtype
}

type proposedEntry struct {
	kindField string
}

// Kind gets the kind of this polymorphic type
func (m *proposedEntry) Kind() string {
	return "ProposedEntry"
}

// SetKind sets the kind of this polymorphic type
func (m *proposedEntry) SetKind(val string) {
}

// UnmarshalProposedEntrySlice unmarshals polymorphic slices of ProposedEntry
func UnmarshalProposedEntrySlice(reader io.Reader, consumer runtime.Consumer) ([]ProposedEntry, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []ProposedEntry
	for _, element := range elements {
		obj, err := unmarshalProposedEntry(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalProposedEntry unmarshals polymorphic ProposedEntry
func UnmarshalProposedEntry(reader io.Reader, consumer runtime.Consumer) (ProposedEntry, error) {
	// we need to read this twice, so first into a buffer
	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalProposedEntry(data, consumer)
}

func unmarshalProposedEntry(data []byte, consumer runtime.Consumer) (ProposedEntry, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the kind property.
	var getType struct {
		Kind string `json:"kind"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("kind", "body", getType.Kind); err != nil {
		return nil, err
	}

	// The value of kind is used to determine which type to create and unmarshal the data into
	switch getType.Kind {
	case "ProposedEntry":
		var result proposedEntry
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "alpine":
		var result Alpine
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "cose":
		var result Cose
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "dsse":
		var result DSSE
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "hashedrekord":
		var result Hashedrekord
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "helm":
		var result Helm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "intoto":
		var result Intoto
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "jar":
		var result Jar
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "rekord":
		var result Rekord
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "rfc3161":
		var result Rfc3161
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "rpm":
		var result Rpm
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	case "tuf":
		var result TUF
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil
	}
	return nil, errors.New(422, "invalid kind value: %q", getType.Kind)
}

// Validate validates this proposed entry
func (m *proposedEntry) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this proposed entry based on context it is used
func (m *proposedEntry) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}