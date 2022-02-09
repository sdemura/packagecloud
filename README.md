## packagecloud

`packagecloud` is a Go CLI for the [packagecloud API](https://packagecloud.io/docs/api).

It is a trivial fork of [mlafeldt/pkgcloud](https://github.com/mlafeldt/pkgcloud),
adding functionality to retry a package push if it already exists.

This fork was created to solve a particular problem, and is not expected to
be actively maintained, although pull requests will be reviewed.

## Installation

```shell
$ go install github.com/edgeworx/packagecloud
```

## Usage

The `PACKAGECLOUD_TOKEN` envar must be set.

To upload a package:

```shell
$ packagecloud push --overwrite user/repo/distro/version ./mypkg_1.2.3_amd64.deb
```

> When the `--overwrite` flag is present, the CLI will yank (delete) the package file
if the API reports that it already exists.

To yank (remove) a package:

```shell
$ packagecloud yank user/repo/distro/version ./mypkg_1.2.3_arm64.deb
```
