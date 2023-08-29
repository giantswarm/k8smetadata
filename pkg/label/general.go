package label

// ServicePriority indicates the relative priority of a cluster.
// See https://github.com/giantswarm/rfc/tree/main/classify-cluster-priority
// for details.
const ServicePriority = "giantswarm.io/service-priority"

const ServicePriorityHighest = "highest"
const ServicePriorityMedium = "medium"
const ServicePriorityLowest = "lowest"

// PreventDeletion indicates that a resource cannot be deleted unless the
// label is removed first.
const PreventDeletion = "giantswarm.io/prevent-deletion"
