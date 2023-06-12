package annotation

//Keep the YAML documentation format as it is used to render in the CRD public documentation. You can use Markdown in the documentation field.

// support:
//   - crd: awsclusters.infrastructure.giantswarm.io
//     apiversion: v1alpha3
//     release: Since 18.0.0
//
// documentation:
//
//	This annotation allows specifying a CIDR to be used by cilium during cluster upgrades from v17 to v18.
const CiliumPodCidr = "cilium.giantswarm.io/pod-cidr"

// Not documented as it's not meant for customers.
// This annotation allows disabling Cilium's kube proxy implementation.
const CiliumForceDisableKubeProxyAnnotation = "cilium.giantswarm.io/force-disable-cilium-kube-proxy"

// support:
//   - crd: clusters.cluster.x-k8s.io
//     apiversion: v1beta1
//     release: Since 19.1.0
//
// documentation:
//
//	This annotation allows specifying what ipam mode to use for cilium. Supported values are "kubernetes", the default
//	and highly recommended mode and "eni". If annotation is missing, "kubernetes" is assumed.
//	Please note it is possible to set this annotation only when cluster has release < v19.0.0 and editing of the value
//	is blocked as soon as v19 upgrade begins.
const CiliumIpamModeAnnotation = "cilium.giantswarm.io/ipam-mode"

const CiliumIpamModeKubernetes = "kubernetes"
const CiliumIpamModeENI = "eni"
