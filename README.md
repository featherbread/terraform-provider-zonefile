# terraform-provider-zonefile

The zonefile provider parses the contents of DNS [**zone files**][zone file],
as defined by [RFC 1035][RFC 1035], and provides structured, well-typed data
for use throughout your `.tf` files. You can find it as `ahamlinman/zonefile`
on both the [OpenTofu][opentofu] and [HashiCorp Terraform][terraform]
registries.

[zone file]: https://en.wikipedia.org/wiki/Zone_file
[RFC 1035]: https://datatracker.ietf.org/doc/html/rfc1035
[opentofu]: https://search.opentofu.org/provider/ahamlinman/zonefile/latest
[terraform]: https://registry.terraform.io/providers/ahamlinman/zonefile/latest

My own code for the provider (committed to this repository) is released under
the terms of the MIT License (see LICENSE.txt). The provider as a whole
incorporates some third-party libraries under the terms of the Mozilla Public
License 2.0. See `release/COPYING.txt` for details.
