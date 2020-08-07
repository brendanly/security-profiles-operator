/*
Copyright 2020 The Kubernetes Authors.

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

package initialize_test

import (
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"testing"

	"github.com/stretchr/testify/require"

	"sigs.k8s.io/seccomp-operator/internal/pkg/initialize"
)

func TestSetupRootless(t *testing.T) {
	tempDir, err := ioutil.TempDir("", "so-test-")
	require.Nil(t, err)
	defer require.Nil(t, os.RemoveAll(tempDir))

	currentUser, err := user.Current()
	require.Nil(t, err)
	uid, err := strconv.Atoi(currentUser.Uid)
	require.Nil(t, err)

	kubeletSeccompRootPath := filepath.Join(tempDir, "a")
	operatorRootPath := filepath.Join(tempDir, "b")
	profilesRootPath := filepath.Join(tempDir, "c")

	for i := 0; i < 2; i++ {
		require.Nil(t, initialize.SetupRootless(
			kubeletSeccompRootPath,
			operatorRootPath,
			profilesRootPath,
			uid,
		))
	}
}