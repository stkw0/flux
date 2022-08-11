// Package feature provides feature flagging capabilities for a Go application.
//
// Flags are defined using the `MakeXXXFlag` functions. These can either be defined
// in code or generated using the feature generator included in this module.
//
// Flags are configured in `flags.yml` at the top of this repository.
// Running `make flags` generates Go code based on this configuration
// to programmatically test flag values in a given request context.
// Boolean flags are the most common case, but integers, floats and
// strings are supported for more complicated experiments.
//
// The `Flagger` interface is the crux of this package.
// It computes a map of feature flag values for a given request context.
// The implementation is application dependent.
// The result can be attached to the context using `Annotate` and then the results
// will be used when checking feature flags.
//
// Flags can be checked using the context after it has been annotated.
// For example, to check a boolean flag that was generated by the `feature` command,
// you might do it this way:
//
//	if feature.MyFeature().Enabled(ctx) {
//	    // my feature is enabled
//	} else {
//	    // my feature is disabled
//	}
package feature
