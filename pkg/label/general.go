package label

// PreventDeletion can be used on certain resource types to prevent deletion
// (see https://docs.giantswarm.io/advanced/deletion-prevention/).
const PreventDeletion = "giantswarm.io/prevent-deletion"

// ServicePriority indicates the relative priority of a cluster.
// See https://github.com/giantswarm/rfc/tree/main/classify-cluster-priority
// for details.
const ServicePriority = "giantswarm.io/service-priority"

const ServicePriorityHighest = "highest"
const ServicePriorityMedium = "medium"
const ServicePriorityLowest = "lowest"
