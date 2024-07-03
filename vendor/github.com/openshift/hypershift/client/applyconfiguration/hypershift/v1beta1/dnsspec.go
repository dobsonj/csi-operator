/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1beta1

// DNSSpecApplyConfiguration represents an declarative configuration of the DNSSpec type for use
// with apply.
type DNSSpecApplyConfiguration struct {
	BaseDomain       *string `json:"baseDomain,omitempty"`
	BaseDomainPrefix *string `json:"baseDomainPrefix,omitempty"`
	PublicZoneID     *string `json:"publicZoneID,omitempty"`
	PrivateZoneID    *string `json:"privateZoneID,omitempty"`
}

// DNSSpecApplyConfiguration constructs an declarative configuration of the DNSSpec type for use with
// apply.
func DNSSpec() *DNSSpecApplyConfiguration {
	return &DNSSpecApplyConfiguration{}
}

// WithBaseDomain sets the BaseDomain field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseDomain field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithBaseDomain(value string) *DNSSpecApplyConfiguration {
	b.BaseDomain = &value
	return b
}

// WithBaseDomainPrefix sets the BaseDomainPrefix field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BaseDomainPrefix field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithBaseDomainPrefix(value string) *DNSSpecApplyConfiguration {
	b.BaseDomainPrefix = &value
	return b
}

// WithPublicZoneID sets the PublicZoneID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PublicZoneID field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithPublicZoneID(value string) *DNSSpecApplyConfiguration {
	b.PublicZoneID = &value
	return b
}

// WithPrivateZoneID sets the PrivateZoneID field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the PrivateZoneID field is set to the value of the last call.
func (b *DNSSpecApplyConfiguration) WithPrivateZoneID(value string) *DNSSpecApplyConfiguration {
	b.PrivateZoneID = &value
	return b
}
