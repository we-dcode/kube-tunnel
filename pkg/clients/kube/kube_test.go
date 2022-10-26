package kube_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/we-dcode/kube-tunnel/pkg/clients/kube"
	"github.com/we-dcode/kube-tunnel/pkg/models"
	"testing"
)

func TestGetServiceWithSinglePortOnDefaultNamespace(t *testing.T) {

	client := kube.MustNew("", "")

	context, err := client.GetServiceContext("kubetunnel-svc")

	assert.NoError(t, err)
	assert.NotNil(t, context)
	assert.Len(t, context.Ports, 1)
}

func TestGetServiceWithMultiplePortsOnDefaultNamespace(t *testing.T) {

	client := kube.MustNew("", "")

	context, err := client.GetServiceContext("kubetunnel-multi-svc")

	assert.NoError(t, err)
	assert.NotNil(t, context)
	assert.Len(t, context.Ports, 2)
}

func TestGetServiceFromExplicitNamespaceWithMultipleLables(t *testing.T) {

	client := kube.MustNew("", "kubetunnel-explicit")

	context, err := client.GetServiceContext("kubetunnel-svc")

	assert.NoError(t, err)
	assert.NotNil(t, context)
	assert.Len(t, context.Ports, 1)
	assert.Len(t, context.LabelSelector, 2)
}

func TestKubePortForward(t *testing.T) {

	client := kube.MustNew("/Users/maordavidov/dcode/gitlab-cicd-kubeconfig.yaml", "kubetunnel")

	listeningPort, err := client.PortForward("kubetunnel-nginx", "7000")

	assert.NoError(t, err)
	assert.NotEqual(t, -1, listeningPort)
}

func TestCreateKubeTunnelResource(t *testing.T) {

	client := kube.MustNew("/Users/maordavidov/dcode/gitlab-cicd-kubeconfig.yaml", "kubetunnel")

	err := client.CreateKubeTunnelResource(models.KubeTunnelResourceSpec{
		Ports:       models.Ports{Values: []string{"80", "8081"}},
		ServiceName: "nginx",
		PodSelectorLabels: map[string]string{
			"app": "nginx",
		},
	})

	assert.NoError(t, err)
}

func TestGetPodLabelsByServiceName(t *testing.T) {

	client := kube.MustNew("/Users/maordavidov/dcode/gitlab-cicd-kubeconfig.yaml", "kubetunnel")

	context, err := client.GetServiceContext("nginx")

	assert.NoError(t, err)
	assert.NotNil(t, context)

	labels, err := client.GetPodLabelsByLabelSelector("kubetunnel", context.LabelSelector)
	assert.NoError(t, err)
	assert.NotNil(t, labels)
}
