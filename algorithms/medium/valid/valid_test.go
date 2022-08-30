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

import "testing"

// TestSherlockValidStringFirstGivenCase implements the test given as first example on hackerrank.
func TestSherlockValidStringFirstGivenCase(t *testing.T) {

	input := "abc"
	expected := "YES"

	result := isValid(input)

	if result != expected {
		t.Errorf("Sherlock valid string first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestSherlockValidStringSecondGivenCase implements the test given as second example on hackerrank.
func TestSherlockValidStringSecondGivenCase(t *testing.T) {

	input := "abcc"
	expected := "YES"

	result := isValid(input)

	if result != expected {
		t.Errorf("Sherlock valid string second case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestSherlockValidStringThirdGivenCase implements the test given as third example on hackerrank.
func TestSherlockValidStringThirdGivenCase(t *testing.T) {

	input := "abccc"
	expected := "NO"

	result := isValid(input)

	if result != expected {
		t.Errorf("Sherlock valid string third case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestSherlockValidStringFourthGivenCase implements the test given as fourth example on hackerrank.
func TestSherlockValidStringFourthGivenCase(t *testing.T) {

	input := "aabbccddeefghi"
	expected := "NO"

	result := isValid(input)

	if result != expected {
		t.Errorf("Sherlock valid string fourth case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestSherlockValidStringFifthGivenCase implements the test given as fifth example on hackerrank.
func TestSherlockValidStringFifthGivenCase(t *testing.T) {

	input := "abcdefghhgfedecba"
	expected := "YES"

	result := isValid(input)

	if result != expected {
		t.Errorf("Sherlock valid string fifth case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// TestSherlockValidStringHiddenCase implements the test given as hidden example on hackerrank.
func TestSherlockValidStringHiddenCase(t *testing.T) {

	input := "ibfdgaeadiaefgbhbdghhhbgdfgeiccbiehhfcggchgghadhdhagfbahhddgghbdehidbibaeaagaeeigffcebfbaieggabcfbiiedc" +
		"abfihchdfabifahcbhagccbdfifhghcadfiadeeaheeddddiecaicbgigccageicehfdhdgafaddhffadigfhhcaedcedecafeacbdacgfgf" +
		"eeibgaiffdehigebhhehiaahfidibccdcdagifgaihacihadecgifihbebffebdfbchbgigeccahgihbcbcaggebaaafgfedbfgagfediddg" +
		"hdgbgehhhifhgcedechahidcbchebheihaadbbbiaiccededchdagfhccfdefigfibifabeiaccghcegfbcghaefifbachebaacbhbfgfdde" +
		"ceababbacgffbagidebeadfihaefefegbghgddbbgddeehgfbhafbccidebgehifafgbghafacgfdccgifdcbbbidfifhdaibgigebigaede" +
		"aaiadegfefbhacgddhchgcbgcaeaieiegiffchbgbebgbehbbfcebciiagacaiechdigbgbghefcahgbhfibhedaeeiffebdiabcifgccdef" +
		"abccdghehfibfiifdaicfedagahhdcbhbicdgibgcedieihcichadgchgbdcdagaihebbabhibcihicadgadfcihdheefbhffiageddhgaha" +
		"idfdhhdbgciiaciegchiiebfbcbhaeagccfhbfhaddagnfieihghfbaggiffbbfbecgaiiidccdceadbbdfgigibgcgchafccdchgifdeiei" +
		"cbaididhfcfdedbhaadedfageigfdehgcdaecaebebebfcieaecfagfdieaefdiedbcadchabhebgehiidfcgahcdhcdhgchhiiheffiifee" +
		"gcfdgbdeffhgeghdfhbfbifgidcafbfcd"

	expected := "YES"

	result := isValid(input)

	if result != expected {
		t.Errorf("Sherlock valid string hidden case was incorrect, got: %s, want: %s.", result, expected)
	}
}
