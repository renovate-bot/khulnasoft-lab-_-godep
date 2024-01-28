// Copyright 2023 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package worker

import (
	"context"
	"testing"

	"github.com/khulnasoft-lab/godep/internal/config"
	"github.com/khulnasoft-lab/godep/internal/postgres"
)

func TestExcluded(t *testing.T) {
	ctx := context.Background()
	defer postgres.ResetTestDB(testDB, t)

	// This test is intending to test PopulateExcluded not just IsExcluded.
	if err := PopulateExcluded(ctx, &config.Config{DynamicExcludeLocation: "testdata/excluded.txt"}, testDB); err != nil {
		t.Fatal(err)
	}

	for _, test := range []struct {
		path string
		want bool
	}{
		{"github.com/golang/go", true},
		{"github.com/golang/godep", false},
	} {
		got, err := testDB.IsExcluded(ctx, test.path)
		if err != nil {
			t.Fatal(err)
		}
		if got != test.want {
			t.Errorf("%q: got %t, want %t", test.path, got, test.want)
		}
	}
}
