// Package annotation defines annotation keys used by Giant Swarm within
// Kubernetes resources.
package annotation

// EncryptionRotationInProgress indicates that there is key rotation in progress. The annotation is stored on the encryption-provider-config secret.
const EncryptionRotationInProgress = "encryption.giantswarm.io/rotation-in-progress"

// EncryptionForceRotation if annotation is set on the encryption-provider-config secret, it will force the key rotation immediately regardless of the age of current encryption key. The annotation is stored on the encryption-provider-config secret.
const EncryptionForceRotation = "encryption.giantswarm.io/force-rotation"

// EncryptionEnableRotation enables the key rotation process of encryption key, if the annotation is not set, the encryption-provider-operator will never rotate keys. The annotation is stored on the encryption-provider-config secret.
const EncryptionEnableRotation = "encryption.giantswarm.io/enable-rotation"

// EncryptionRewriteTimestamp is set to every secret in worklaod cluster that was rewritten as result of key rotation, the value will contain timestamp when the rewrite happened.
const EncryptionRewriteTimestamp = "encryption.giantswarm.io/rewrited-at"

// EncryptionLastRotation is set on to encryption-provider-secret to indicate when was the last key rotation, its used to calculate when the next rotation should happen.
const EncryptionLastRotation = "encryption.giantswarm.io/last-rotation"
