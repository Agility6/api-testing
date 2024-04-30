/*
Copyright 2024 API Testing Authors.

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
package service

import (
	"fmt"
	"net/http"
)

// WebRunnerHandler accepts the HTTP requests and run the testSuite
func WebRunnerHandler(w http.ResponseWriter, r *http.Request,
	params map[string]string) {
	testSuite := params["suite"]
	testCase := params["case"]

	fmt.Println(testSuite, testCase)
	fmt.Println(r.URL.Query())
}
