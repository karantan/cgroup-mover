//go:build integration
// +build integration

package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findChildProcesses(t *testing.T) {
	want := []int{2, 3}
	got := findChildProcesses("fixtures/grp1/cgroup.procs")
	assert.Equal(t, want, got)
}

func Test_addToCgroup(t *testing.T) {
	defer os.Truncate("fixtures/grp2/cgroup.procs", 0)
	want := "2\n3\n"
	addToCgroup([]int{2, 3}, "fixtures/grp2/cgroup.procs")
	got, _ := os.ReadFile("fixtures/grp2/cgroup.procs")
	assert.Equal(t, want, string(got))
}
