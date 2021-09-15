/*
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"

	log "github.com/sirupsen/logrus"
)

const scooter = "🛴"

var repeat = 1
var port = "8080"

func main() {

	log.SetFormatter(&log.TextFormatter{
		DisableColors: true,
		FullTimestamp: true,
	})

	debugChecker()

	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}

	http.HandleFunc("/", scooterHandler)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func scooterHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s endpoint hit", scooter)
	fmt.Fprint(w, repScooter())
}

func repScooter() string {
	if os.Getenv("REPEAT_SCOOTER") != "" {
		r, _ := strconv.Atoi(os.Getenv("REPEAT_SCOOTER"))
		repeat = r
	}

	return strings.Repeat(scooter, repeat)
}

func debugChecker() {
	debugFlag, _ := strconv.ParseBool(os.Getenv("DEBUG_SCOOTER"))
	if debugFlag {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.WarnLevel)
	}
}
