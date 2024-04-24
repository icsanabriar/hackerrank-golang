// Copyright 2020 Iván Camilo Sanabria
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

import "testing"

// TestRansomNoteFirstGivenCase implements the test given as first example on hackerrank.
func TestRansomNoteFirstGivenCase(t *testing.T) {
	magazine := []string{"give", "me", "one", "grand", "today", "night"}
	note := []string{"give", "one", "grand", "today"}
	expected := "Yes"

	result := checkMagazine(magazine, note)

	if result != expected {
		t.Errorf("Ransom note first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestRansomNoteSecondGivenCase implements the test given as second example on hackerrank.
func TestRansomNoteSecondGivenCase(t *testing.T) {
	magazine := []string{"two", "times", "three", "is", "not", "four"}
	note := []string{"two", "times", "two", "is", "four"}
	expected := "No"

	result := checkMagazine(magazine, note)

	if result != expected {
		t.Errorf("Ransom note second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestRansomNoteThirdGivenCase implements the test given as third example on hackerrank.
func TestRansomNoteThirdGivenCase(t *testing.T) {
	magazine := []string{"ive", "got", "a", "lovely", "bunch", "of", "coconuts"}
	note := []string{"ive", "got", "some", "coconuts"}
	expected := "No"

	result := checkMagazine(magazine, note)

	if result != expected {
		t.Errorf("Ransom note third case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestRansomNoteHiddenCase implements the test given as hidden example on hackerrank.
func TestRansomNoteHiddenCase(t *testing.T) {
	magazine := []string{"avtq", "ekpvq", "z", "rdvzf", "m", "zu", "bof", "pfkzl", "ekpvq", "pfkzl", "bof", "zu", "ekpvq", "ekpvq", "ekpvq", "ekpvq", "z"}
	note := []string{"m", "z", "z", "avtq", "zu", "bof", "pfkzl", "pfkzl", "pfkzl", "rdvzf", "rdvzf", "avtq", "ekpvq", "rdvzf", "avtq"}
	expected := "No"

	result := checkMagazine(magazine, note)

	if result != expected {
		t.Errorf("Ransom note hidden case was incorrect, got: %s, want: %s.", result, expected)
	}
}
