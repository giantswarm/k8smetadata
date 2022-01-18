package annotation

// Metadata annotation stores an app metadata URL from the appCatalog's index.yaml.
const AppMetadata = "application.giantswarm.io/metadata"

// AppName annotation is used by the chart status watcher to find the
// app CR for each chart CR.
const AppName = "chart-operator.giantswarm.io/app-name"

// AppNamespace annotation is used by the chart status watcher to find the
// app CR for chart CRs which are always in the giantswarm namespace.
const AppNamespace = "chart-operator.giantswarm.io/app-namespace"

// AppOperatorCordonReason is the name of the annotation that indicates
// the reason of why operators should not apply any update to this app CR.
const AppOperatorCordonReason = "app-operator.giantswarm.io/cordon-reason"

// AppOperatorCordonUntil is the name of the annotation that indicates
// the expiration date for this cordon rule.
const AppOperatorCordonUntil = "app-operator.giantswarm.io/cordon-until"

// AppOperatorLatestConfigMapVersion is the highest resource version among the configmaps
// app CRs depends on.
const AppOperatorLatestConfigMapVersion = "app-operator.giantswarm.io/latest-configmap-version"

// AppOperatorLatestSecretVersion is the highest resource version among the secret
// app CRs depends on.
const AppOperatorLatestSecretVersion = "app-operator.giantswarm.io/latest-secret-version"

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
