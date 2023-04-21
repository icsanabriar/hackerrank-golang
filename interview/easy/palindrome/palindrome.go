// Copyright 2023 Iván Camilo Sanabria
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

// isPalindrome validates if the given string is palindrome or not.
func isPalindrome(s string) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-i-1] {
			return false
		}
	}
	return true
}

// palindromeIndex search the index to remove character to make the given string a palindrome.
func palindromeIndex(s string) int {

	if len(s) > 0 && isPalindrome(s) {
		return -1
	}

	for i := 0; i < len(s)/2; i++ {

		if s[i] != s[len(s)-i-1] {

			if isPalindrome(s[:i] + s[i+1:]) {
				return i
			}

			if isPalindrome(s[:len(s)-i-1] + s[len(s)-i:]) {
				return len(s) - i - 1
			}

			return -1
		}
	}

	return -1
}
