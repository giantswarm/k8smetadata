package label

// AppOperatorWatching is the label added to configmaps by app-operator when
// it is watching for values changes.
const AppOperatorWatching = "app-operator.giantswarm.io/watching"

// ClusterAppsOperatorWatching is the label which should be added to Cluster CRs to indicate
// that cluster-apps-operator should watch it.
const ClusterAppsOperatorWatching = "cluster-apps-operator.giantswarm.io/watching"
