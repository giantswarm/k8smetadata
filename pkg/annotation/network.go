package annotation

// NetworkWildcardCNAMETarget overrides the default wildcard CNAME record target for a workload cluster.
// When set, the wildcard DNS record (*.<cluster>.<baseDomain>) will point to this value instead of the default ingress hostname.
const NetworkWildcardCNAMETarget = "network.giantswarm.io/wildcard-cname-target"
