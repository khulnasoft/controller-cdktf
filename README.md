# controller-cdktf

This repo contains the generated code for [cdktf](https://github.com/hashicorp/terraform-cdk) in Go. The implementation is being developed in [khulnasoft/controller].

This package is only used internally at Khulnasoft - the generated code is public for ease of use and avoid performance issue with large amount of generated content being tracked in a Git repository.

## Usage

### Adding a new provider or module to CDKTF

Follow https://github.com/khulnasoft/go/cdktf-provider-gen#usage

```bash
make <target>
```

### Upgrading CDKTF

Review the [changelog](https://developer.hashicorp.com/terraform/cdktf/release#upgrade-guides) of the target release.
Watch out for breaking changes and adjust the upgrade plan if neccessary.

Bumpd `CDKTF_VERSION` in `Makefile`:

```diff
-CDKTF_VERSION=0.16.3
+CDKTF_VERSION=0.19.2
```

Re-generate all providers and modules:

```bash
make -j4
```

## FAQ

### Why not use the pre-built providers?

We would like to use specific versions of provides, hence we are not using pre-built providers, such as [cdktf/cdktf-provider-google-go](https://github.com/cdktf/cdktf-provider-google-go).

### Why multiple modules instead of one?

Generated code combined from all providers and modules exceed the limit of `go get`, and this cannot be changed.

[khulnasoft/controller]: https://github.com/khulnasoft/controller
