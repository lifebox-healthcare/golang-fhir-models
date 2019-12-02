// Copyright 2019 The Samply Development Community
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

package fhir

import "encoding/json"

// THIS FILE IS GENERATED BY https://github.com/samply/golang-fhir-models
// PLEASE DO NOT EDIT BY HAND

// Bundle is documented here http://hl7.org/fhir/StructureDefinition/Bundle
type Bundle struct {
	Id            *string       `json:"id,omitempty"`
	Meta          *Meta         `json:"meta,omitempty"`
	ImplicitRules *string       `json:"implicitRules,omitempty"`
	Language      *string       `json:"language,omitempty"`
	Identifier    *Identifier   `json:"identifier,omitempty"`
	Type          BundleType    `json:"type"`
	Timestamp     *string       `json:"timestamp,omitempty"`
	Total         *int          `json:"total,omitempty"`
	Link          []BundleLink  `json:"link,omitempty"`
	Entry         []BundleEntry `json:"entry,omitempty"`
	Signature     *Signature    `json:"signature,omitempty"`
}
type BundleLink struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Relation          string      `json:"relation"`
	Url               string      `json:"url"`
}
type BundleEntry struct {
	Id                *string              `json:"id,omitempty"`
	Extension         []Extension          `json:"extension,omitempty"`
	ModifierExtension []Extension          `json:"modifierExtension,omitempty"`
	Link              []BundleLink         `json:"link,omitempty"`
	FullUrl           *string              `json:"fullUrl,omitempty"`
	Resource          json.RawMessage      `json:"resource,omitempty"`
	Search            *BundleEntrySearch   `json:"search,omitempty"`
	Request           *BundleEntryRequest  `json:"request,omitempty"`
	Response          *BundleEntryResponse `json:"response,omitempty"`
}
type BundleEntrySearch struct {
	Id                *string          `json:"id,omitempty"`
	Extension         []Extension      `json:"extension,omitempty"`
	ModifierExtension []Extension      `json:"modifierExtension,omitempty"`
	Mode              *SearchEntryMode `json:"mode,omitempty"`
	Score             *string          `json:"score,omitempty"`
}
type BundleEntryRequest struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Method            HTTPVerb    `json:"method"`
	Url               string      `json:"url"`
	IfNoneMatch       *string     `json:"ifNoneMatch,omitempty"`
	IfModifiedSince   *string     `json:"ifModifiedSince,omitempty"`
	IfMatch           *string     `json:"ifMatch,omitempty"`
	IfNoneExist       *string     `json:"ifNoneExist,omitempty"`
}
type BundleEntryResponse struct {
	Id                *string         `json:"id,omitempty"`
	Extension         []Extension     `json:"extension,omitempty"`
	ModifierExtension []Extension     `json:"modifierExtension,omitempty"`
	Status            string          `json:"status"`
	Location          *string         `json:"location,omitempty"`
	Etag              *string         `json:"etag,omitempty"`
	LastModified      *string         `json:"lastModified,omitempty"`
	Outcome           json.RawMessage `json:"outcome,omitempty"`
}
type OtherBundle Bundle

// MarshalJSON marshals the given Bundle as JSON into a byte slice
func (r Bundle) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherBundle
		ResourceType string `json:"resourceType"`
	}{
		OtherBundle:  OtherBundle(r),
		ResourceType: "Bundle",
	})
}

// UnmarshalBundle unmarshals a Bundle.
func UnmarshalBundle(b []byte) (Bundle, error) {
	var bundle Bundle
	if err := json.Unmarshal(b, &bundle); err != nil {
		return bundle, err
	}
	return bundle, nil
}