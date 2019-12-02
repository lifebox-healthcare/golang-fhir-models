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

// SupplyDelivery is documented here http://hl7.org/fhir/StructureDefinition/SupplyDelivery
type SupplyDelivery struct {
	Id                *string                     `json:"id,omitempty"`
	Meta              *Meta                       `json:"meta,omitempty"`
	ImplicitRules     *string                     `json:"implicitRules,omitempty"`
	Language          *string                     `json:"language,omitempty"`
	Text              *Narrative                  `json:"text,omitempty"`
	Extension         []Extension                 `json:"extension,omitempty"`
	ModifierExtension []Extension                 `json:"modifierExtension,omitempty"`
	Identifier        []Identifier                `json:"identifier,omitempty"`
	BasedOn           []Reference                 `json:"basedOn,omitempty"`
	PartOf            []Reference                 `json:"partOf,omitempty"`
	Status            *SupplyDeliveryStatus       `json:"status,omitempty"`
	Patient           *Reference                  `json:"patient,omitempty"`
	Type              *CodeableConcept            `json:"type,omitempty"`
	SuppliedItem      *SupplyDeliverySuppliedItem `json:"suppliedItem,omitempty"`
	Supplier          *Reference                  `json:"supplier,omitempty"`
	Destination       *Reference                  `json:"destination,omitempty"`
	Receiver          []Reference                 `json:"receiver,omitempty"`
}
type SupplyDeliverySuppliedItem struct {
	Id                *string     `json:"id,omitempty"`
	Extension         []Extension `json:"extension,omitempty"`
	ModifierExtension []Extension `json:"modifierExtension,omitempty"`
	Quantity          *Quantity   `json:"quantity,omitempty"`
}
type OtherSupplyDelivery SupplyDelivery

// MarshalJSON marshals the given SupplyDelivery as JSON into a byte slice
func (r SupplyDelivery) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		OtherSupplyDelivery
		ResourceType string `json:"resourceType"`
	}{
		OtherSupplyDelivery: OtherSupplyDelivery(r),
		ResourceType:        "SupplyDelivery",
	})
}

// UnmarshalSupplyDelivery unmarshals a SupplyDelivery.
func UnmarshalSupplyDelivery(b []byte) (SupplyDelivery, error) {
	var supplyDelivery SupplyDelivery
	if err := json.Unmarshal(b, &supplyDelivery); err != nil {
		return supplyDelivery, err
	}
	return supplyDelivery, nil
}