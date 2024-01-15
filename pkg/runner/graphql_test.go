/*
Copyright 2023 API Testing Authors.

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

package runner

import (
	"context"
	"net/http"
	"testing"

	"github.com/h2non/gock"
	atest "github.com/linuxsuren/api-testing/pkg/testing"
	"github.com/linuxsuren/api-testing/pkg/util"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestGraphQL(t *testing.T) {
	runner := NewSimpleTestCaseRunner()
	graphqlRunner := NewGraphQLRunner(runner)

	testcase := &atest.TestCase{}
	err := yaml.Unmarshal([]byte(simpleGraphQLRequest), testcase)
	assert.NoError(t, err)

	defer gock.Off()
	gock.New("http://foo").Post("/graphql").
		MatchHeader(util.ContentType, util.JSON).
		Reply(http.StatusOK)
	_, err = graphqlRunner.RunTestCase(testcase, nil, context.TODO())
	assert.NoError(t, err)
}

var simpleGraphQLRequest = `
request:
  api: http://foo/graphql
  body:
    query: |
      query xxx {
       	bookById(id: "book") {
          id
          name
        }
      }
`