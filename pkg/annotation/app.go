package annotation

// AppOperatorCordonReason is the name of the annotation that indicates
// the reason of why operators should not apply any update to this app CR.
const AppOperatorCordonReason = "cordon-reason"

// AppOperatorCordonUntil is the name of the annotation that indicates
// the expiration date for this cordon rule.
const AppOperatorCordonUntil = "cordon-until"

// AppOperatorPaused annotation when present prevents app-operator from
// reconciling the resource.
const AppOperatorPaused = "app-operator.giantswarm.io/paused"

// AppOperatorWebhookURL is the URL that chart-operator reports chart updates.
const AppOperatorWebhookURL = "app-operator.giantswarm.io/webhook-url"

// AppOwners annotation is defined in Chart.yaml and added to AppCatalogEntry CRs.
// It is used when an app is owned by multiple teams.
const AppOwners = "application.giantswarm.io/owners"

// AppTeam annotation is defined in Chart.yaml and added to AppCatalogEntry CRs.
// It is used when an app is owned by a single team.
const AppTeam = "application.giantswarm.io/team"

// CordonReason is the name of the annotation that indicates
// the reason of why operators should not apply any update to this app CR.
const ChartOperatorCordonReason = "cordon-reason"

// CordonUntil is the name of the annotation that indicates
// the expiration date for this cordon rule.
const ChartOperatorCordonUntil = "cordon-until"

// LatestConfigMapVersion is the highest resource version among the configmaps
// app CRs depends on.
const LatestConfigMapVersion = "app-operator.giantswarm.io/giantswarm.io/latest-configmap-version"

// LatestSecretVersion is the highest resource version among the secret
// app CRs depends on.
const LatestSecretVersion = "app-operator.giantswarm.io/latest-secret-version"
