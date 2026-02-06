# Changelog

All notable changes to this project will be documented in this file.

Please choose versions by [Semantic Versioning](http://semver.org/).

* MAJOR version when you make incompatible API changes,
* MINOR version when you add functionality in a backwards-compatible manner, and
* PATCH version when you make backwards-compatible bug fixes.

## v1.10.2

- Update Go to 1.25.7
- Update bborbe/errors to v1.5.2
- Update bborbe/math to v1.3.1
- Update testing frameworks (ginkgo v2.28.1, gomega v1.39.1)
- Update numerous indirect dependencies

## v1.10.1

- Update Go to 1.25.5
- Update golang.org/x/crypto to v0.47.0
- Update dependencies

## v1.10.0

- update go and deps

## v1.9.1

- Rename InvalidTypeError to ErrInvalidType to follow Go error naming conventions

## v1.9.0

- Add HasStrings and HasString interfaces for custom string conversion
- Add support for parsing slices of string subtypes via reflection
- Add support for parsing slices of types implementing String() method
- Improve Makefile with .PHONY declarations for all targets
- Enable race detection in test target
- Add go-modtool for go.mod formatting
- Remove deprecated golang.org/x/lint/golint dependency
- Enhance .golangci.yml with 8 additional linters (funlen, gocognit, nestif, maintidx, errname, unparam, bodyclose, forcetypeassert)
- Fix typo in .golangci.yml depguard rule
- Add standard ignore patterns to .gitignore

## v1.8.3

- Upgrade osv-scanner to v2.2.3 with configuration file support
- Enhance Makefile osv-scanner target to support optional .osv-scanner.toml config
- Update development tool dependencies

## v1.8.2

- Update Go version to 1.25.2
- Update CI workflow to use Go 1.25.2
- Update dependencies (golang.org/x/text, golang.org/x/crypto, golang.org/x/net, golang.org/x/sys, golang.org/x/term, golang.org/x/exp)
- Update prometheus client libraries
- Update internal dependencies (bborbe/collection, bborbe/run)
- Add exclude directive for incompatible golang.org/x/tools v0.38.0

## v1.8.1

- Enhance README.md with comprehensive documentation
- Add installation instructions and usage examples
- Add status badges for CI, coverage, and Go Report Card
- Add GoDoc comments to all exported functions
- Improve API documentation for pkg.go.dev

## v1.8.0

- Add ParseASCII function with correct capitalization
- Add golangci-lint configuration with revive linter
- Add GitHub Actions workflows for CI/CD and code review
- Enhance Makefile with lint target
- Update dependencies and test coverage
- Deprecate ParseAscii in favor of ParseASCII

## v1.7.1

- add tests
- go mod update

## v1.7.0

- remove vendor
- go mod update

## v1.6.1

- add license
- go mod update

## v1.6.0

- allow parse subtypes 

## v1.5.0

- add parse int64 array
- go mod update

## v1.4.2

- add InvalidTypeError to ParseString

## v1.4.1

- go mod update
- improve test

## v1.4.0

- add ParseAscii

## v1.3.2

- improve ParseStrings
- go mod update

## v1.3.1

- allow more types for ParseString
- go mod update

## v1.3.0

- add parse with default
- go mod update

## v1.2.0

- allow parse string to bool

## v1.1.0

- allow stringer in parseInt

## v1.0.1

- go mod update

## v1.0.0

- Initial Version
