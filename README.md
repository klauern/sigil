# sigil [![Build Status](https://cloud.drone.io/api/badges/danmx/sigil/status.svg)](https://cloud.drone.io/danmx/sigil)

## Description

> *Sigil* is the hub of the Great Wheel, a city at the center of the Outlands, the most balanced of neutral areas at the center of the planes. Also known as the "City of Doors" for the multitude of portals to other planes of existence and the Cage since those portals are the only way in or out, it is the setting for most of Planescape: Torment.

*Sigil* is an AWS SSM Session manager client inspired by [xen0l's aws-gate](https://github.com/xen0l/aws-gate).

## Build

To build binaries for all platforms (Linux, Mac, Windows) and Docker image run:

```console
make build
```

To run specific build use:

```console
make build-[linux|mac|windows|docker]
```

Binaries are located in:

- Linux: `dist/sigil-linux`
- Mac: `dist/sigil-darwin`
- Windows: `dist/sigil-windows`

### Docker

To only build docker image run:

```console
make build-docker
```

It'll create a docker image tagged `sigil:{version}` where `{version}` corresponds to sigil's current version.

## Example usage

Docker:

```console
docker run --rm -it -v ${HOME}/.sigil:/home/.sigil -v ${HOME}/.aws:/home/.aws danmx/sigil:0.0 list --output-format wide
```

Binary:

```console
sigil-darwin -r eu-west-1 session --type instance-id --target i-xxxxxxxxxxxxxxxxx
```

While using [aws-vault](https://github.com/99designs/aws-vault):

```console
aws-vault exec AWS_PROFILE -- sigil-darwin -r eu-west-1 session --type instance-id --target i-xxxxxxxxxxxxxxxxx
```
