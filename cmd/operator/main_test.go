package main

import (
	"gotest.tools/v3/assert"
	"testing"
)

func TestPatchServiceWithLabel(t *testing.T) {

	kube := connectToKubernetes("kubetunnel") // TODO: change this one

	err := patchServiceWithLabel(kube, "nginx", false)

	assert.NilError(t, err)
}
