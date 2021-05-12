package annotation

// ChartOperatorCordonReason is the name of the annotation that indicates
// the reason of why operators should not apply any update to this app CR.
const ChartOperatorCordonReason = "chart-operator.giantswarm.io/cordon-reason"

// ChartOperatorCordonUntil is the name of the annotation that indicates
// the expiration date for this cordon rule.
const ChartOperatorCordonUntil = "chart-operator.giantswarm.io/cordon-until"

// ChartOperatorForceHelmUpgrade is the name of the annotation that controls
// whether force is used when upgrading the Helm release.
const ChartOperatorForceHelmUpgrade = "chart-operator.giantswarm.io/force-helm-upgrade"

// ChartOperatorPrefix is the prefix for annotations that control logic inside chart-operator.
const ChartOperatorPrefix = "chart-operator.giantswarm.io"
