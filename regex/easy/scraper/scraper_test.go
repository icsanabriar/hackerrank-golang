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
	"bufio"
	"fmt"
	"log"
	"os"
	"testing"
)

// TestBuildExchangeScraperFirstGivenCase implements the test given as first example on hackerrank.
func TestBuildExchangeScraperFirstGivenCase(t *testing.T) {

	input := `        <div class="question-summary" id="question-summary-80407">
        <div class="statscontainer">
            <div class="statsarrow"></div>
            <div class="stats">
                <div class="vote">
                    <div class="votes">
                        <span class="vote-count-post "><strong>2</strong></span>
                        <div class="viewcount">votes</div>
                    </div>
                </div>
                <div class="status answered">
                    <strong>1</strong>answer
                </div>
            </div>



    <div class="views " title="60 views">
                        60 views
    </div>
        </div>
        <div class="summary">
            <h3><a href="/questions/80407/about-power-supply-of-opertional-amplifier" class="question-hyperlink">about power supply of opertional amplifier</a></h3>
            <div class="excerpt">
                I am constructing an operational amplifier as shown in the following figure. I use a batter as supplier for the OP Amp and set it up as a non-inverting amp circuit. I saw that the output was clipped ...
            </div>

            <div class="tags t-op-amp">
                <a href="/questions/tagged/op-amp" class="post-tag" title="show questions tagged 'op-amp'" rel="tag">op-amp</a>

            </div>
            <div class="started fr">


        <div class="user-info ">
            <div class="user-action-time">


                        asked <span title="2013-08-27 21:49:14Z" class="relativetime">11 hours ago</span>
            </div>
            <div class="user-gravatar32">
                <a href="/users/17060/user1285419"><div class=""><img src="https://www.gravatar.com/avatar/08ee68b20a4eceff26f7eee99b708c08?s=32&d=identicon&r=PG" alt="" width="32" height="32"></div></a>
            </div>
            <div class="user-details">
                <a href="/users/17060/user1285419">user1285419</a><br>
                <span class="reputation-score" title="reputation score" dir="ltr">165</span><span title="5 bronze badges"><span class="badge3"></span><span class="badgecount">5</span></span>
            </div>
        </div>

            </div>
        </div>
    </div>

    <div class="question-summary" id="question-summary-80405">
        <div class="statscontainer">
            <div class="statsarrow"></div>
            <div class="stats">
                <div class="vote">
                    <div class="votes">
                        <span class="vote-count-post "><strong>4</strong></span>
                        <div class="viewcount">votes</div>
                    </div>
                </div>
                <div class="status answered-accepted">
                    <strong>2</strong>answers
                </div>
            </div>



    <div class="views " title="64 views">
                        64 views
    </div>
        </div>
        <div class="summary">
            <h3><a href="/questions/80405/5v-regulator-power-dissipation" class="question-hyperlink">5V Regulator Power Dissipation</a></h3>
            <div class="excerpt">
                I am using a 5V regulator (LP2950) from ON Semiconductor. I am using this for USB power and I'm feeding in 9V from an adapter. USB requires maximum of 500mA right? So the maximum power dissipation in ...
            </div>

            <div class="tags t-voltage-regulator t-surface-mount t-heatsink t-5v t-power-dissipation">
                <a href="/questions/tagged/voltage-regulator" class="post-tag" title="show questions tagged 'voltage-regulator'" rel="tag">voltage-regulator</a> <a href="/questions/tagged/surface-mount" class="post-tag" title="show questions tagged 'surface-mount'" rel="tag">surface-mount</a> <a href="/questions/tagged/heatsink" class="post-tag" title="show questions tagged 'heatsink'" rel="tag">heatsink</a> <a href="/questions/tagged/5v" class="post-tag" title="show questions tagged '5v'" rel="tag">5v</a> <a href="/questions/tagged/power-dissipation" class="post-tag" title="show questions tagged 'power-dissipation'" rel="tag">power-dissipation</a>

            </div>
            <div class="started fr">


        <div class="user-info ">
            <div class="user-action-time">


                        asked <span title="2013-08-27 21:39:31Z" class="relativetime">11 hours ago</span>
            </div>
            <div class="user-gravatar32">
                <a href="/users/10082/david-norman"><div class=""><img src="https://www.gravatar.com/avatar/8b073417e471077280b3fc5ff2eaf1f7?s=32&d=identicon&r=PG" alt="" width="32" height="32"></div></a>
            </div>
            <div class="user-details">
                <a href="/users/10082/david-norman">David Norman</a><br>
                <span class="reputation-score" title="reputation score" dir="ltr">322</span><span title="3 silver badges"><span class="badge2"></span><span class="badgecount">3</span></span><span title="10 bronze badges"><span class="badge3"></span><span class="badgecount">10</span></span>
            </div>
        </div>

            </div>
        </div>
    </div>`

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	expected := "80407;about power supply of opertional amplifier;11 hours ago\n" +
		"80405;5V Regulator Power Dissipation;11 hours ago\n"

	scrape(stdout, input)

	result := readTestFile()

	if result != expected {
		t.Errorf("Build a stack exchange scraper first case was incorrect, got: %s, want: %s.", result, expected)
	}
}

// readTestFile is responsible of reading the output of the program written in the given writer.
func readTestFile() string {
	text := ""

	file, err := os.Open(os.Getenv("OUTPUT_PATH"))
	if err != nil {
		log.Fatal(err)
	}

	defer func(stdout *os.File) {
		err := stdout.Close()
		if err != nil {
			fmt.Println("Error reading output path file.")
		}
	}(file)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text = text + scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return text
}
