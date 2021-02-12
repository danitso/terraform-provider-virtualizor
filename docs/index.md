---
layout: home
title: Introduction
nav_order: 1
---

# Virtualizor Provider

This provider for [Terraform](https://www.terraform.io/) is used for interacting with resources supported by [Virtualizor](https://www.virtualizor.com). The provider needs to be configured with the proper endpoint and key before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```
provider "virtualizor" {
  endpoint = "https://hostname:4083"
  key      = "client-api-key-here"
}
```

## Argument Reference

In addition to [generic provider arguments](https://www.terraform.io/docs/configuration/providers.html) (e.g. `alias` and `version`), the following arguments are supported in the Virtualizor `provider` block:

* `endpoint` - (Required) The API endpoint (`https://hostname:4083`)
* `key` - (Required) The API key
