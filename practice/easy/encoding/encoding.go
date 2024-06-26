// Copyright 2022 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//	https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"bytes"
	"encoding/json"
	"io"
)

// Manager is given struct to write proper json encoder.
type Manager struct {
	FullName       string `json:"full_name,omitempty"`
	Position       string `json:"position,omitempty"`
	Age            int64  `json:"age,omitempty"`
	YearsInCompany int64  `json:"years_in_company,omitempty"`
}

// EncodeManager generate a reader for given manager to encode as json.
func EncodeManager(manager *Manager) (io.Reader, error) {
	result, err := json.Marshal(manager)
	return bytes.NewReader(result), err
}
