package annotation

// FluxKustomizeReconcile is used to manage Flux reconciliation of the resource holding this annotation.
// It can take the following values: enalbed, disabled.
// When set to disabled Flux will no longer apply changes to the resource.
const FluxKustomizeReconcile = "kustomize.toolkit.fluxcd.io/reconcile"
