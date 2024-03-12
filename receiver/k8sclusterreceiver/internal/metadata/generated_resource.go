// Code generated by mdatagen. DO NOT EDIT.

package metadata

import (
	"go.opentelemetry.io/collector/pdata/pcommon"
)

// ResourceBuilder is a helper struct to build resources predefined in metadata.yaml.
// The ResourceBuilder is not thread-safe and must not to be used in multiple goroutines.
type ResourceBuilder struct {
	config ResourceAttributesConfig
	res    pcommon.Resource
}

// NewResourceBuilder creates a new ResourceBuilder. This method should be called on the start of the application.
func NewResourceBuilder(rac ResourceAttributesConfig) *ResourceBuilder {
	return &ResourceBuilder{
		config: rac,
		res:    pcommon.NewResource(),
	}
}

// SetContainerID sets provided value as "container.id" attribute.
func (rb *ResourceBuilder) SetContainerID(val string) {
	if rb.config.ContainerID.Enabled {
		rb.res.Attributes().PutStr("container.id", val)
	}
}

// SetContainerImageName sets provided value as "container.image.name" attribute.
func (rb *ResourceBuilder) SetContainerImageName(val string) {
	if rb.config.ContainerImageName.Enabled {
		rb.res.Attributes().PutStr("container.image.name", val)
	}
}

// SetContainerImageTag sets provided value as "container.image.tag" attribute.
func (rb *ResourceBuilder) SetContainerImageTag(val string) {
	if rb.config.ContainerImageTag.Enabled {
		rb.res.Attributes().PutStr("container.image.tag", val)
	}
}

// SetContainerRuntime sets provided value as "container.runtime" attribute.
func (rb *ResourceBuilder) SetContainerRuntime(val string) {
	if rb.config.ContainerRuntime.Enabled {
		rb.res.Attributes().PutStr("container.runtime", val)
	}
}

// SetContainerRuntimeVersion sets provided value as "container.runtime.version" attribute.
func (rb *ResourceBuilder) SetContainerRuntimeVersion(val string) {
	if rb.config.ContainerRuntimeVersion.Enabled {
		rb.res.Attributes().PutStr("container.runtime.version", val)
	}
}

// SetK8sContainerName sets provided value as "k8s.container.name" attribute.
func (rb *ResourceBuilder) SetK8sContainerName(val string) {
	if rb.config.K8sContainerName.Enabled {
		rb.res.Attributes().PutStr("k8s.container.name", val)
	}
}

// SetK8sCronjobName sets provided value as "k8s.cronjob.name" attribute.
func (rb *ResourceBuilder) SetK8sCronjobName(val string) {
	if rb.config.K8sCronjobName.Enabled {
		rb.res.Attributes().PutStr("k8s.cronjob.name", val)
	}
}

// SetK8sCronjobUID sets provided value as "k8s.cronjob.uid" attribute.
func (rb *ResourceBuilder) SetK8sCronjobUID(val string) {
	if rb.config.K8sCronjobUID.Enabled {
		rb.res.Attributes().PutStr("k8s.cronjob.uid", val)
	}
}

// SetK8sDaemonsetName sets provided value as "k8s.daemonset.name" attribute.
func (rb *ResourceBuilder) SetK8sDaemonsetName(val string) {
	if rb.config.K8sDaemonsetName.Enabled {
		rb.res.Attributes().PutStr("k8s.daemonset.name", val)
	}
}

// SetK8sDaemonsetUID sets provided value as "k8s.daemonset.uid" attribute.
func (rb *ResourceBuilder) SetK8sDaemonsetUID(val string) {
	if rb.config.K8sDaemonsetUID.Enabled {
		rb.res.Attributes().PutStr("k8s.daemonset.uid", val)
	}
}

// SetK8sDeploymentName sets provided value as "k8s.deployment.name" attribute.
func (rb *ResourceBuilder) SetK8sDeploymentName(val string) {
	if rb.config.K8sDeploymentName.Enabled {
		rb.res.Attributes().PutStr("k8s.deployment.name", val)
	}
}

// SetK8sDeploymentUID sets provided value as "k8s.deployment.uid" attribute.
func (rb *ResourceBuilder) SetK8sDeploymentUID(val string) {
	if rb.config.K8sDeploymentUID.Enabled {
		rb.res.Attributes().PutStr("k8s.deployment.uid", val)
	}
}

// SetK8sHpaName sets provided value as "k8s.hpa.name" attribute.
func (rb *ResourceBuilder) SetK8sHpaName(val string) {
	if rb.config.K8sHpaName.Enabled {
		rb.res.Attributes().PutStr("k8s.hpa.name", val)
	}
}

// SetK8sHpaUID sets provided value as "k8s.hpa.uid" attribute.
func (rb *ResourceBuilder) SetK8sHpaUID(val string) {
	if rb.config.K8sHpaUID.Enabled {
		rb.res.Attributes().PutStr("k8s.hpa.uid", val)
	}
}

// SetK8sJobName sets provided value as "k8s.job.name" attribute.
func (rb *ResourceBuilder) SetK8sJobName(val string) {
	if rb.config.K8sJobName.Enabled {
		rb.res.Attributes().PutStr("k8s.job.name", val)
	}
}

// SetK8sJobUID sets provided value as "k8s.job.uid" attribute.
func (rb *ResourceBuilder) SetK8sJobUID(val string) {
	if rb.config.K8sJobUID.Enabled {
		rb.res.Attributes().PutStr("k8s.job.uid", val)
	}
}

// SetK8sKubeletVersion sets provided value as "k8s.kubelet.version" attribute.
func (rb *ResourceBuilder) SetK8sKubeletVersion(val string) {
	if rb.config.K8sKubeletVersion.Enabled {
		rb.res.Attributes().PutStr("k8s.kubelet.version", val)
	}
}

// SetK8sNamespaceName sets provided value as "k8s.namespace.name" attribute.
func (rb *ResourceBuilder) SetK8sNamespaceName(val string) {
	if rb.config.K8sNamespaceName.Enabled {
		rb.res.Attributes().PutStr("k8s.namespace.name", val)
	}
}

// SetK8sNamespaceUID sets provided value as "k8s.namespace.uid" attribute.
func (rb *ResourceBuilder) SetK8sNamespaceUID(val string) {
	if rb.config.K8sNamespaceUID.Enabled {
		rb.res.Attributes().PutStr("k8s.namespace.uid", val)
	}
}

// SetK8sNodeName sets provided value as "k8s.node.name" attribute.
func (rb *ResourceBuilder) SetK8sNodeName(val string) {
	if rb.config.K8sNodeName.Enabled {
		rb.res.Attributes().PutStr("k8s.node.name", val)
	}
}

// SetK8sNodeUID sets provided value as "k8s.node.uid" attribute.
func (rb *ResourceBuilder) SetK8sNodeUID(val string) {
	if rb.config.K8sNodeUID.Enabled {
		rb.res.Attributes().PutStr("k8s.node.uid", val)
	}
}

// SetK8sPodName sets provided value as "k8s.pod.name" attribute.
func (rb *ResourceBuilder) SetK8sPodName(val string) {
	if rb.config.K8sPodName.Enabled {
		rb.res.Attributes().PutStr("k8s.pod.name", val)
	}
}

// SetK8sPodQosClass sets provided value as "k8s.pod.qos_class" attribute.
func (rb *ResourceBuilder) SetK8sPodQosClass(val string) {
	if rb.config.K8sPodQosClass.Enabled {
		rb.res.Attributes().PutStr("k8s.pod.qos_class", val)
	}
}

// SetK8sPodUID sets provided value as "k8s.pod.uid" attribute.
func (rb *ResourceBuilder) SetK8sPodUID(val string) {
	if rb.config.K8sPodUID.Enabled {
		rb.res.Attributes().PutStr("k8s.pod.uid", val)
	}
}

// SetK8sReplicasetName sets provided value as "k8s.replicaset.name" attribute.
func (rb *ResourceBuilder) SetK8sReplicasetName(val string) {
	if rb.config.K8sReplicasetName.Enabled {
		rb.res.Attributes().PutStr("k8s.replicaset.name", val)
	}
}

// SetK8sReplicasetUID sets provided value as "k8s.replicaset.uid" attribute.
func (rb *ResourceBuilder) SetK8sReplicasetUID(val string) {
	if rb.config.K8sReplicasetUID.Enabled {
		rb.res.Attributes().PutStr("k8s.replicaset.uid", val)
	}
}

// SetK8sReplicationcontrollerName sets provided value as "k8s.replicationcontroller.name" attribute.
func (rb *ResourceBuilder) SetK8sReplicationcontrollerName(val string) {
	if rb.config.K8sReplicationcontrollerName.Enabled {
		rb.res.Attributes().PutStr("k8s.replicationcontroller.name", val)
	}
}

// SetK8sReplicationcontrollerUID sets provided value as "k8s.replicationcontroller.uid" attribute.
func (rb *ResourceBuilder) SetK8sReplicationcontrollerUID(val string) {
	if rb.config.K8sReplicationcontrollerUID.Enabled {
		rb.res.Attributes().PutStr("k8s.replicationcontroller.uid", val)
	}
}

// SetK8sResourcequotaName sets provided value as "k8s.resourcequota.name" attribute.
func (rb *ResourceBuilder) SetK8sResourcequotaName(val string) {
	if rb.config.K8sResourcequotaName.Enabled {
		rb.res.Attributes().PutStr("k8s.resourcequota.name", val)
	}
}

// SetK8sResourcequotaUID sets provided value as "k8s.resourcequota.uid" attribute.
func (rb *ResourceBuilder) SetK8sResourcequotaUID(val string) {
	if rb.config.K8sResourcequotaUID.Enabled {
		rb.res.Attributes().PutStr("k8s.resourcequota.uid", val)
	}
}

// SetK8sStatefulsetName sets provided value as "k8s.statefulset.name" attribute.
func (rb *ResourceBuilder) SetK8sStatefulsetName(val string) {
	if rb.config.K8sStatefulsetName.Enabled {
		rb.res.Attributes().PutStr("k8s.statefulset.name", val)
	}
}

// SetK8sStatefulsetUID sets provided value as "k8s.statefulset.uid" attribute.
func (rb *ResourceBuilder) SetK8sStatefulsetUID(val string) {
	if rb.config.K8sStatefulsetUID.Enabled {
		rb.res.Attributes().PutStr("k8s.statefulset.uid", val)
	}
}

// SetOpenshiftClusterquotaName sets provided value as "openshift.clusterquota.name" attribute.
func (rb *ResourceBuilder) SetOpenshiftClusterquotaName(val string) {
	if rb.config.OpenshiftClusterquotaName.Enabled {
		rb.res.Attributes().PutStr("openshift.clusterquota.name", val)
	}
}

// SetOpenshiftClusterquotaUID sets provided value as "openshift.clusterquota.uid" attribute.
func (rb *ResourceBuilder) SetOpenshiftClusterquotaUID(val string) {
	if rb.config.OpenshiftClusterquotaUID.Enabled {
		rb.res.Attributes().PutStr("openshift.clusterquota.uid", val)
	}
}

// SetOsDescription sets provided value as "os.description" attribute.
func (rb *ResourceBuilder) SetOsDescription(val string) {
	if rb.config.OsDescription.Enabled {
		rb.res.Attributes().PutStr("os.description", val)
	}
}

// SetOsType sets provided value as "os.type" attribute.
func (rb *ResourceBuilder) SetOsType(val string) {
	if rb.config.OsType.Enabled {
		rb.res.Attributes().PutStr("os.type", val)
	}
}

// Emit returns the built resource and resets the internal builder state.
func (rb *ResourceBuilder) Emit() pcommon.Resource {
	r := rb.res
	rb.res = pcommon.NewResource()
	return r
}
