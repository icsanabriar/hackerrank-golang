// Copyright 2020 Iván Camilo Sanabria
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

// checkMagazine return YES if the Harold's note can be formed using the magazine, otherwise prints NO.
func checkMagazine(magazine []string, note []string) string {
	wordMap := make(map[string]int)

	// Count the number of words in the magazine.
	for index := range magazine {
		word := magazine[index]
		if val, ok := wordMap[word]; ok {
			wordMap[word] = val + 1
		} else {
			wordMap[word] = 1
		}
	}

	// Find out every word of the note.
	for index := range note {
		word := note[index]
		if val, ok := wordMap[word]; ok {
			if val > 1 {
				wordMap[word] = val - 1
			} else {
				delete(wordMap, word)
			}
		} else {
			return "No"
		}
	}

	// All words of the note are in the magazine.
	return "Yes"
}
