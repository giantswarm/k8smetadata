# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add `application.giantswarm.io/in-cluster-app` annotation.

## [0.10.0] - 2022-03-16

### Added

- Add `alpha.aws.giantswarm.io/ebs-volume-throughput` annotation.
- Add `alpha.aws.giantswarm.io/ebs-volume-iops` annotation.

## [0.9.4] - 2022-03-04

### Added

- Add `alpha.aws.giantswarm.io/iam-roles-for-service-accounts` annotation.

## [0.9.3] - 2022-02-16

### Added

- Add annotation `app-operator.giantswarm.io/trigger-reconciliation`.

## [0.9.2] - 2022-02-01

### Added

- Add annotation `node.giantswarm.io/cgroupv1`.

## [0.9.1] - 2022-01-21

## [0.9.0] - 2022-01-21

### Added

- Add annotations for the encryption-provider-operator.

## [0.8.0] - 2022-01-18

### Added

- Add `AppName` annotation for use by chart CR watcher in app-operator.

## [0.7.1] - 2021-12-08

### Fixed

- Move `ClusterAppsOperatorWatching` to `label` package instead of `annotation` as intended.

## [0.7.0] - 2021-12-08

### Added

- Add `ClusterAppsOperatorWatching` label as a replacement for `ClusterAppsOperatorVersion` when
  only one version of the operator is running to indicate which clusters should be watched (in
  cases where `cluster-operator` is running in the same cluster).

## [0.6.0] - 2021-11-16

### Added

- Add AWSCNIPrefixDelegation annotation.

## [0.5.0] - 2021-11-11

### Added

- Add annotation to indicate a resource shouldn't be removed (`giantswarm.io/keep`)

## [0.4.0] - 2021-10-04

### Added

- Add annotations for `upgrade-schedule-operator`

## [0.3.0] - 2021-05-12

- Add ChartOperatorForceHelmUpggrade annotation.
- Add ClusterAppsOperatorVersion label.

## [0.2.0] - 2021-05-11

### Added

- Add annotations from https://github.com/giantswarm/app/tree/v4.12.0/pkg/annotation
- Add annotations from https://github.com/giantswarm/app-operator/tree/v4.4.0/pkg/annotation
- Add labels from https://github.com/giantswarm/app/tree/v4.12.0/pkg/label

## [0.1.0] - 2021-04-23

### Added

- Add annotations from https://github.com/giantswarm/apiextensions/tree/v3.22.0/pkg/annotation
- Add labels from https://github.com/giantswarm/apiextensions/tree/v3.22.0/pkg/label

[Unreleased]: https://github.com/giantswarm/k8smetadata/compare/v0.10.0...HEAD
[0.10.0]: https://github.com/giantswarm/k8smetadata/compare/v0.9.4...v0.10.0
[0.9.4]: https://github.com/giantswarm/k8smetadata/compare/v0.9.3...v0.9.4
[0.9.3]: https://github.com/giantswarm/k8smetadata/compare/v0.9.2...v0.9.3
[0.9.2]: https://github.com/giantswarm/k8smetadata/compare/v0.9.1...v0.9.2
[0.9.1]: https://github.com/giantswarm/k8smetadata/compare/v0.9.0...v0.9.1
[0.9.0]: https://github.com/giantswarm/k8smetadata/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/giantswarm/k8smetadata/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/giantswarm/k8smetadata/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/giantswarm/k8smetadata/compare/v0.6.0...v0.7.0
[0.6.0]: https://github.com/giantswarm/k8smetadata/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/k8smetadata/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/k8smetadata/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/k8smetadata/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/k8smetadata/compare/v0.1.0...v0.2.0
[0.1.0]: https://github.com/giantswarm/k8smetadata/releases/tag/v0.1.0
