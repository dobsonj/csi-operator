// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/openshift/api/operator/v1"
)

// ForwardPluginApplyConfiguration represents an declarative configuration of the ForwardPlugin type for use
// with apply.
type ForwardPluginApplyConfiguration struct {
	Upstreams        []string                              `json:"upstreams,omitempty"`
	Policy           *v1.ForwardingPolicy                  `json:"policy,omitempty"`
	TransportConfig  *DNSTransportConfigApplyConfiguration `json:"transportConfig,omitempty"`
	ProtocolStrategy *v1.ProtocolStrategy                  `json:"protocolStrategy,omitempty"`
}

// ForwardPluginApplyConfiguration constructs an declarative configuration of the ForwardPlugin type for use with
// apply.
func ForwardPlugin() *ForwardPluginApplyConfiguration {
	return &ForwardPluginApplyConfiguration{}
}

// WithUpstreams adds the given value to the Upstreams field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the Upstreams field.
func (b *ForwardPluginApplyConfiguration) WithUpstreams(values ...string) *ForwardPluginApplyConfiguration {
	for i := range values {
		b.Upstreams = append(b.Upstreams, values[i])
	}
	return b
}

// WithPolicy sets the Policy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Policy field is set to the value of the last call.
func (b *ForwardPluginApplyConfiguration) WithPolicy(value v1.ForwardingPolicy) *ForwardPluginApplyConfiguration {
	b.Policy = &value
	return b
}

// WithTransportConfig sets the TransportConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the TransportConfig field is set to the value of the last call.
func (b *ForwardPluginApplyConfiguration) WithTransportConfig(value *DNSTransportConfigApplyConfiguration) *ForwardPluginApplyConfiguration {
	b.TransportConfig = value
	return b
}

// WithProtocolStrategy sets the ProtocolStrategy field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ProtocolStrategy field is set to the value of the last call.
func (b *ForwardPluginApplyConfiguration) WithProtocolStrategy(value v1.ProtocolStrategy) *ForwardPluginApplyConfiguration {
	b.ProtocolStrategy = &value
	return b
}