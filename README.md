# toggle

Toggle is a simple feature flagging API

## contributions

Pull requests welcomed. ❤️

## license

MIT licensed. See [licenses](./licenses) for more details.

## setup

This project uses [Taskfile](https://taskfile.dev) to manage build steps.

To install the Go binaries necessary for development, install Taskfile and then run: `task setup`

## building the Docker image

To build the Docker image for the Toggle API server locally run: `task build-docker`

To run the image there's a convenience function: `task run-docker`

An example of how to connect to the newly created server: `task run -- client create-scope`

##

## design

### overview

The name _Toggle_ was chosen because it provides a semi-opinionated method for managing feature flags. Current home-rolled feature flag systems often end up with a murky binary enable/disable pattern for features that can lead to confusion. ("What does it mean to enable the `DisableFeatureX` flag?")

Toggle tries to side step this issue by having an opinion; things are _toggled_ (think "light switch") to be either _on_ or _off_. This leads to shorter feature names (`FeatureX` instead of `EnableFeatureX`) with a cleaner understanding of what toggling it on means ("`FeatureX` is enabled"). This however necessitates that any behavior to be toggled must be described in terms of being on or off.

One design consideration is that toggle is for _feature flagging and is not a general configuration management tool_. It does not support switches for non-boolean values (eg. `ApiTimeoutInSeconds`) and is unlikely to in the future. For these cases a general purpose configuration tool should be used instead. Ignoring non-binary use cases allows Toggle to be very performant even in high-traffic environments that may make thousands of lookups per second.

### toggle sets

A _toggle set_ is a collection of multiple feature _toggles_ that are logically grouped and allows for turning them all on or off as a single switch.

An example use case would be the ability to toggle on a complete set of features that represents optional capabilities that an enterprise customer would have access to.

Imagine the scenario where sales is attempting to upsell a new customer by providing a preview of the features available to them at an enterprise tier. In this case it would be appropriate to toggle (on a per tenant basis) the feature set for enterprise.

### default values

It is _highly_ recommended that a toggle and toggle sets default value is set to `OFF`. This is because gRPC does not bother to encode and send a field whose value matches its default value; the field is not transmitted on the wire and is instead only populated (inflated) during the decoding phase using the defined protobuf schema. This means there is a substantial savings on both network transfer latency and the message encoding/decoding for requests that are primarily sending or receiving default values.

For example, suppose you have the following toggle set:

```
\- enterprise (default = off)
  \- single_sign_on (default = off)
  \- custom_domains (default = off)
  \- support_dashboard (default = off)
```

For each case where a customer is non-enterprise (the default) there will be no toggle set payload sent in the response. Instead, the receiving client will automatically inflate the response with the correct default values saving considerable time and network bandwidth.

This technique is also utilized when serializing/deserializing the toggle set to persistent storage allowing for even large numbers of toggles to be efficiently stored and retrieved.

## examples

Gives examples for the overall operational lifecycle of a feature flag from creation to use.

This cycle is roughly:

1. Create a new scope
2. Create a feature flag
3. Assign a value to the feature flag
4. (Optional) List all of the flags available. This is mainly useful for building UIs to manage the flag set.

### create a new scope

#### defining scope

The most complicated part of feature flagging is defining the situation under which a toggle's value applies. (Commonly called "scope".)

Toggle allows the operator to define their own scoping strategy in order to service a wider range of use cases.

Common scoping strategies supported include:

-   Scoping by the environment (eg. `development`, `staging`, `production`) that the flag applies to.
-   Scoping by tenant (eg. `customer_a`, `customer_b`, `customer_c`)
-   Scoping by mode (eg. `enterprise`, `experimental`, etc.)

### create a new feature flag

### assign a value to the flag

### list all flags

## features to support

-   [ ] Docker image creation and publishing
-   [ ] Setting and getting scope
-   [ ] Creating and updating toggle sets and toggles
-   [ ] Leader election and multiple instances of API server
-   [ ] Toggle local caching and performance improvements
-   [ ] Kubernetes deployment manifests
-   [ ] Toggle API CRDs and Kubernetes operator
