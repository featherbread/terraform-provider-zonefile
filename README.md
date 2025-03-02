# terraform-provider-zonefile

The zonefile provider parses the contents of DNS [**zone files**][zone file], as
defined by [RFC 1035][RFC 1035], and provides structured, well-typed data for
use in Terraform and compatible forks like OpenTofu.

[zone file]: https://en.wikipedia.org/wiki/Zone_file
[RFC 1035]: https://datatracker.ietf.org/doc/html/rfc1035

As of 1 March 2025, the provider is available for HashiCorp Terraform users
from the Terraform Registry, and a request has been made to publish the
provider in the OpenTofu Registry. You can also build and install the plugin
into a [local mirror directory][mirror] for testing.

[mirror]: https://developer.hashicorp.com/terraform/cli/config/config-file#implied-local-mirror-directories

My own code for the provider (committed to this repository) is released under
the terms of the MIT License (see LICENSE.txt). The provider as a whole
incorporates some third-party libraries under the terms of the Mozilla Public
License 2.0. See `release/COPYING.txt` for details.
