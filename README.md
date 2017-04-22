# Gofpher [![GoDoc](https://godoc.org/github.com/asteris-llc/gofpher?status.svg)](https://godoc.org/github.com/asteris-llc/gofpher) [![Go Report Card](https://goreportcard.com/badge/github.com/asteris-llc/gofpher)](https://goreportcard.com/report/github.com/asteris-llc/gofpher)

Gofpher is a collection of FP-like utilities for go, useful for rapidly
prototyping an application.

Key features:

- Support for monads and monadic pipelines
- Lazily evaluated lists
- Either monads
- A pipeline execution tool for easily assembling collections of monadic
  computations

<!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-generate-toc again -->
**Table of Contents**

- [Gofpher](#Gofpher)
    - [Usage](#usage)
    - [License](#license)

<!-- markdown-toc end -->

## Usage

Most of the libraries here have some level of interdependence.  In general all
of the monadic constructs should implement `monad.Monad` and other tools may be
written around the behavior of certain monads.  See the tests for real-life
examples of how to use some of these features.

## License

Gofpher is licensed under the Apache 2.0 license. See [LICENSE](LICENSE) for
full details.
