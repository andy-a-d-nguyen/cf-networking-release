---
title: Example Apps Overview
expires_at: never
tags: [cf-networking-release]
---

<!-- vim-markdown-toc GFM -->

* [Example Apps Overview](#example-apps-overview)
  * [Use case: you want an app to do network debugging](#use-case-you-want-an-app-to-do-network-debugging)
  * [Use case: you want to try out container-to-container networking](#use-case-you-want-to-try-out-container-to-container-networking)
  * [Use case: you want to try out container-to-container networking with service discovery](#use-case-you-want-to-try-out-container-to-container-networking-with-service-discovery)
  * [Use case: you want to use an a8registry](#use-case-you-want-to-use-an-a8registry)
  * [Use case: you want to test an invalid HTTP response](#use-case-you-want-to-test-an-invalid-http-response)

<!-- vim-markdown-toc -->
# Example Apps Overview

This doc provides an overview of the example apps provided and the uses cases they provide.

## Use case: you want an app to do network debugging
**App:** [Proxy](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/proxy)

**Description**: Proxy has endpoints for using dig and ping, for showing stats, uploading and downloading an arbitrary number of bytes, and more. It is a good app for general debugging purposes. 


## Use case: you want to try out container-to-container networking
**App:** [Cats and Dogs](https://github.com/cloudfoundry-attic/cf-networking-examples/blob/master/docs/c2c-no-service-discovery.md)

**Description:** This example demonstrates container-to-container networking via HTTP and UDP between a frontend and backend app.


## Use case: you want to try out container-to-container networking with service discovery
**App:** [Cats and Dogs with Service Discovery](https://github.com/cloudfoundry-attic/cf-networking-examples/blob/master/docs/c2c-with-service-discovery.md)

**Description:** This example demonstrates container-to-container networking via HTTP and UDP between a frontend and backend app with service discovery.


## Use case: you want to use an a8registry
**App:** [tick](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/tick)

**Description:** This is a simple app that registers itself with an a8registry on a regular interval.

## Use case: you want to test an invalid HTTP response
**App:** [raw-response](https://github.com/cloudfoundry/cf-networking-release/tree/develop/src/example-apps/raw-response)

**Description:** This is a simple app is  designed to make writing arbitrary
  HTTP responses out from a CloudFoundry app, bypassing any of Golang's built-in
  net/http or net/httputil behaviors.
