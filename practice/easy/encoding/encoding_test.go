// Copyright 2022 Iván Camilo Sanabria
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package main

import (
	"io/ioutil"
	"testing"
)

// TestEncodingManagerFirstGivenCase implements the test given as first example on hackerrank.
func TestEncodingManagerFirstGivenCase(t *testing.T) {
	input := Manager{
		FullName:       "Ivan Camilo Sanabria",
		Position:       "CEO",
		YearsInCompany: 0,
	}

	expected := `{"full_name":"Ivan Camilo Sanabria","position":"CEO"}`

	resultReader, err := EncodeManager(&input)
	if err != nil {
		t.Errorf("Encoding manager first case was incorrect, got an error: %v.", err)
	}

	result, err := ioutil.ReadAll(resultReader)
	if err != nil {
		t.Errorf("Encoding manager first case was incorrect, got an error: %v.", err)
	}

	if string(result) != expected {
		t.Errorf("Encoding manager first case was incorrect, got: %s, want: %s.", string(result), expected)
	}
}
