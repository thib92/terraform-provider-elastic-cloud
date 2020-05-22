Elastic Cloud Terraform Provider
==================

This is the repository for the Elastic Cloud Terraform Provider, which allows one to use Terraform with Elastic's SaaS offering, Elastic Cloud.
Learn more about Elastic Cloud at [https://www.elastic.co/cloud/](https://www.elastic.co/cloud/)

For general information about Terraform, visit the [official website](https://www.terraform.io) and the [GitHub project page](https://github.com/hashicorp/terraform).


Support, Bugs, Feature Requests
-------------------------------

This Terraform Provider is built by the community. Please report any support request, bug or feature request using the Issues section of this repository.


Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) >= 0.12.x
-	[Go](https://golang.org/doc/install) >= 1.12


Using the provider
----------------------

To use a released provider in your Terraform environment, run [`terraform init`](https://www.terraform.io/docs/commands/init.html) and Terraform will automatically install the provider. To specify a particular provider version when installing released providers, see the [`Terraform documentation on provider versioning`](https://www.terraform.io/docs/configuration/providers.html#version-provider-versions).

To instead use a custom-built provider in your Terraform environment (e.g. the provider binary from the build instructions above), follow the instructions to [install it as a plugin](https://www.terraform.io/docs/plugins/basics.html#installing-a-plugin). After placing it into your plugins directory, run `terraform init` to initialize it.

For either installation method, documentation about the provider specific configuration options can be found on the [provider's website](https://www.terraform.io/docs/providers/).


Building The Provider
---------------------

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the Go `install` command: 
```sh
$ go install
```


Adding Dependencies
---------------------

This provider uses [Go modules](https://github.com/golang/go/wiki/Modules).
Please see the Go documentation for the most up to date information about using Go modules.

To add a new dependency `github.com/author/dependency` to your Terraform provider:

```
go get github.com/author/dependency
go mod tidy
```

Then commit the changes to `go.mod` and `go.sum`.


Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (see [Requirements](#requirements) above).

To compile the provider, run `go install`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.


Running the acceptance test
---------------------------

Acceptance tests are not yet available for this module.
