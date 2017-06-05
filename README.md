<p align="center">
  <h1 align="center">Devise</h1>
  <p align="center">Configuration as Code.</p>
  <p align="center">
    <a href="https://gitter.im/autonomy/devise"><img alt="Gitter" src="https://img.shields.io/gitter/room/autonomy/devise.svg?style=flat-square"></a>
    <a href="https://godoc.org/github.com/autonomy/devise"><img alt="GoDoc" src="http://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square"></a>
  </p>
  <p align="center">
    <a href="https://travis-ci.org/autonomy/devise"><img alt="Travis" src="https://img.shields.io/travis/autonomy/devise.svg?style=flat-square"></a>
    <a href="https://codecov.io/gh/autonomy/devise"><img alt="Codecov" src="https://img.shields.io/codecov/c/github/autonomy/devise.svg?style=flat-square"></a>
    <a href="https://goreportcard.com/report/github.com/autonomy/devise"><img alt="Go Report Card" src="https://goreportcard.com/badge/github.com/autonomy/devise?style=flat-square"></a>
  </p>
  <p align="center">
    <a href="https://github.com/autonomy/devise/releases/latest"><img alt="Release" src="https://img.shields.io/github/release/autonomy/devise.svg?style=flat-square"></a>
    <a href="https://github.com/autonomy/devise/releases/latest"><img alt="GitHub (pre-)release" src="https://img.shields.io/github/release/autonomy/devise/all.svg?style=flat-square"></a>
  </p>
</p>

---

**Devise** is a tool for standardized, programmatic, and bi-directional configuration management.

The key features of Devise are:
-   **Configuration as code**: Devise encapsulates a wide range of APIs into a single powerfully simple API that applications can consume via [gRPC](http://www.grpc.io/). By enabling applications to configure themselves programmatically, Devise facilitates configuration as code by providing key-value pair lookups, or fully rendered templates at the request of the application.
-   **Universal RPC Framework**: Utilizing [protobuf](https://developers.google.com/protocol-buffers/) and [gRPC](http://www.grpc.io/), applications written in any of the [supported laguages](http://www.grpc.io/docs/reference/) can use Devise.

Getting Started
---------------
By default Devise uses ports `8080` and `50000` for the UI and backend respectively. The UI provides an interface for administration while the backend exposes an API that can be consumed directly inside client applications. To get started:
```sh
$ docker run \
    --rm \
    --detach \
    --publish 8080:8080 \
    --publish 50000:50000 \
    --name devise \
    autonomy/devise:latest serve
```
For an example on how a client might use Devise, see the [example client](https://github.com/autonomy/devise/tree/master/examples/client).
> **Note:** The default storage for Devise is an in-memory datastore. It is intended for development. Production deployments should make use of a production quality datastore.

Devloping Devise
----------------
The build of Devise depends on [conform](https://github.com/autonomy/conform):
```sh
$ go get -u github.com/autonomy/conform
```
> **Note:** Conform leverages [multi-stage builds](https://docs.docker.com/engine/userguide/eng-image/multistage-build/). Docker 17.05.0 or greater is required.

To build the image, run:
```sh
$ conform enforce image
```

### License
[![license](https://img.shields.io/github/license/autonomy/devise.svg?style=flat-square)](https://github.com/autonomy/devise/blob/master/LICENSE)
