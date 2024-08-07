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
package server_test

import (
	"context"
	"errors"
	"testing"

	"github.com/linuxsuren/api-testing/pkg/server"
	"github.com/stretchr/testify/assert"
)

func TestUnimplement(t *testing.T) {
	unimplemented := &server.UnimplementedRunnerServer{}
	_, err := unimplemented.Run(context.TODO(), nil)
	assert.NotNil(t, err)

	_, err = unimplemented.GetVersion(context.Background(), nil)
	assert.NotNil(t, err)

	var reply *server.HelloReply
	assert.Empty(t, reply.GetMessage())
	assert.Empty(t, reply.GetError())
	assert.Empty(t, &server.Empty{})

	var task *server.TestTask
	assert.Empty(t, task.GetData())
	assert.Empty(t, task.GetKind())
	assert.Empty(t, task.GetCaseName())
	assert.Empty(t, task.GetLevel())
	assert.Nil(t, task.GetEnv())

	task = &server.TestTask{
		Data:     "data",
		Kind:     "kind",
		CaseName: "casename",
		Level:    "level",
		Env:      map[string]string{},
	}
	assert.Equal(t, "data", task.GetData())
	assert.Equal(t, "kind", task.GetKind())
	assert.Equal(t, "casename", task.GetCaseName())
	assert.Equal(t, "level", task.GetLevel())
	assert.Equal(t, map[string]string{}, task.GetEnv())
}

func TestServer(t *testing.T) {
	client, _ := server.NewFakeClient(context.Background(), "version", nil)
	ver, err := client.GetVersion(context.Background(), &server.Empty{})
	assert.NotNil(t, ver)
	assert.Equal(t, "version", ver.GetVersion())
	assert.Nil(t, err)

	var testResult *server.TestResult
	testResult, err = client.Run(context.Background(), &server.TestTask{})
	assert.NotNil(t, testResult)
	assert.Nil(t, err)

	var reply *server.HelloReply
	reply, err = client.Sample(context.Background(), &server.Empty{})
	assert.Nil(t, err)
	assert.Empty(t, reply.GetMessage())

	clientWithErr, _ := server.NewFakeClient(context.Background(), "version", errors.New("fake"))
	ver, err = clientWithErr.GetVersion(context.Background(), &server.Empty{})
	assert.NotNil(t, err)
	assert.Nil(t, ver)

	testResult, err = clientWithErr.Run(context.Background(), &server.TestTask{})
	assert.NotNil(t, err)
	assert.Nil(t, testResult)
}
