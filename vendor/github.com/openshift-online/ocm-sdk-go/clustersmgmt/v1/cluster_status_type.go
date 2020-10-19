/*
Copyright (c) 2020 Red Hat, Inc.

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

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/ocm-sdk-go/clustersmgmt/v1

// ClusterStatusKind is the name of the type used to represent objects
// of type 'cluster_status'.
const ClusterStatusKind = "ClusterStatus"

// ClusterStatusLinkKind is the name of the type used to represent links
// to objects of type 'cluster_status'.
const ClusterStatusLinkKind = "ClusterStatusLink"

// ClusterStatusNilKind is the name of the type used to nil references
// to objects of type 'cluster_status'.
const ClusterStatusNilKind = "ClusterStatusNil"

// ClusterStatus represents the values of the 'cluster_status' type.
//
// Detailed status of a cluster.
type ClusterStatus struct {
	id                    *string
	href                  *string
	link                  bool
	dnsReady              *bool
	description           *string
	provisionErrorCode    *string
	provisionErrorMessage *string
	state                 *ClusterState
}

// Kind returns the name of the type of the object.
func (o *ClusterStatus) Kind() string {
	if o == nil {
		return ClusterStatusNilKind
	}
	if o.link {
		return ClusterStatusLinkKind
	}
	return ClusterStatusKind
}

// ID returns the identifier of the object.
func (o *ClusterStatus) ID() string {
	if o != nil && o.id != nil {
		return *o.id
	}
	return ""
}

// GetID returns the identifier of the object and a flag indicating if the
// identifier has a value.
func (o *ClusterStatus) GetID() (value string, ok bool) {
	ok = o != nil && o.id != nil
	if ok {
		value = *o.id
	}
	return
}

// Link returns true iif this is a link.
func (o *ClusterStatus) Link() bool {
	return o != nil && o.link
}

// HREF returns the link to the object.
func (o *ClusterStatus) HREF() string {
	if o != nil && o.href != nil {
		return *o.href
	}
	return ""
}

// GetHREF returns the link of the object and a flag indicating if the
// link has a value.
func (o *ClusterStatus) GetHREF() (value string, ok bool) {
	ok = o != nil && o.href != nil
	if ok {
		value = *o.href
	}
	return
}

// Empty returns true if the object is empty, i.e. no attribute has a value.
func (o *ClusterStatus) Empty() bool {
	return o == nil || (o.id == nil &&
		o.dnsReady == nil &&
		o.description == nil &&
		o.provisionErrorCode == nil &&
		o.provisionErrorMessage == nil &&
		o.state == nil &&
		true)
}

// DNSReady returns the value of the 'DNS_ready' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// DNSReady from Provisioner
func (o *ClusterStatus) DNSReady() bool {
	if o != nil && o.dnsReady != nil {
		return *o.dnsReady
	}
	return false
}

// GetDNSReady returns the value of the 'DNS_ready' attribute and
// a flag indicating if the attribute has a value.
//
// DNSReady from Provisioner
func (o *ClusterStatus) GetDNSReady() (value bool, ok bool) {
	ok = o != nil && o.dnsReady != nil
	if ok {
		value = *o.dnsReady
	}
	return
}

// Description returns the value of the 'description' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Detailed description of the cluster status.
func (o *ClusterStatus) Description() string {
	if o != nil && o.description != nil {
		return *o.description
	}
	return ""
}

// GetDescription returns the value of the 'description' attribute and
// a flag indicating if the attribute has a value.
//
// Detailed description of the cluster status.
func (o *ClusterStatus) GetDescription() (value string, ok bool) {
	ok = o != nil && o.description != nil
	if ok {
		value = *o.description
	}
	return
}

// ProvisionErrorCode returns the value of the 'provision_error_code' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Provisioning Error Code
func (o *ClusterStatus) ProvisionErrorCode() string {
	if o != nil && o.provisionErrorCode != nil {
		return *o.provisionErrorCode
	}
	return ""
}

// GetProvisionErrorCode returns the value of the 'provision_error_code' attribute and
// a flag indicating if the attribute has a value.
//
// Provisioning Error Code
func (o *ClusterStatus) GetProvisionErrorCode() (value string, ok bool) {
	ok = o != nil && o.provisionErrorCode != nil
	if ok {
		value = *o.provisionErrorCode
	}
	return
}

// ProvisionErrorMessage returns the value of the 'provision_error_message' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// Provisioning Error Message
func (o *ClusterStatus) ProvisionErrorMessage() string {
	if o != nil && o.provisionErrorMessage != nil {
		return *o.provisionErrorMessage
	}
	return ""
}

// GetProvisionErrorMessage returns the value of the 'provision_error_message' attribute and
// a flag indicating if the attribute has a value.
//
// Provisioning Error Message
func (o *ClusterStatus) GetProvisionErrorMessage() (value string, ok bool) {
	ok = o != nil && o.provisionErrorMessage != nil
	if ok {
		value = *o.provisionErrorMessage
	}
	return
}

// State returns the value of the 'state' attribute, or
// the zero value of the type if the attribute doesn't have a value.
//
// The overall state of the cluster.
func (o *ClusterStatus) State() ClusterState {
	if o != nil && o.state != nil {
		return *o.state
	}
	return ClusterState("")
}

// GetState returns the value of the 'state' attribute and
// a flag indicating if the attribute has a value.
//
// The overall state of the cluster.
func (o *ClusterStatus) GetState() (value ClusterState, ok bool) {
	ok = o != nil && o.state != nil
	if ok {
		value = *o.state
	}
	return
}

// ClusterStatusListKind is the name of the type used to represent list of objects of
// type 'cluster_status'.
const ClusterStatusListKind = "ClusterStatusList"

// ClusterStatusListLinkKind is the name of the type used to represent links to list
// of objects of type 'cluster_status'.
const ClusterStatusListLinkKind = "ClusterStatusListLink"

// ClusterStatusNilKind is the name of the type used to nil lists of objects of
// type 'cluster_status'.
const ClusterStatusListNilKind = "ClusterStatusListNil"

// ClusterStatusList is a list of values of the 'cluster_status' type.
type ClusterStatusList struct {
	href  *string
	link  bool
	items []*ClusterStatus
}

// Kind returns the name of the type of the object.
func (l *ClusterStatusList) Kind() string {
	if l == nil {
		return ClusterStatusListNilKind
	}
	if l.link {
		return ClusterStatusListLinkKind
	}
	return ClusterStatusListKind
}

// Link returns true iif this is a link.
func (l *ClusterStatusList) Link() bool {
	return l != nil && l.link
}

// HREF returns the link to the list.
func (l *ClusterStatusList) HREF() string {
	if l != nil && l.href != nil {
		return *l.href
	}
	return ""
}

// GetHREF returns the link of the list and a flag indicating if the
// link has a value.
func (l *ClusterStatusList) GetHREF() (value string, ok bool) {
	ok = l != nil && l.href != nil
	if ok {
		value = *l.href
	}
	return
}

// Len returns the length of the list.
func (l *ClusterStatusList) Len() int {
	if l == nil {
		return 0
	}
	return len(l.items)
}

// Empty returns true if the list is empty.
func (l *ClusterStatusList) Empty() bool {
	return l == nil || len(l.items) == 0
}

// Get returns the item of the list with the given index. If there is no item with
// that index it returns nil.
func (l *ClusterStatusList) Get(i int) *ClusterStatus {
	if l == nil || i < 0 || i >= len(l.items) {
		return nil
	}
	return l.items[i]
}

// Slice returns an slice containing the items of the list. The returned slice is a
// copy of the one used internally, so it can be modified without affecting the
// internal representation.
//
// If you don't need to modify the returned slice consider using the Each or Range
// functions, as they don't need to allocate a new slice.
func (l *ClusterStatusList) Slice() []*ClusterStatus {
	var slice []*ClusterStatus
	if l == nil {
		slice = make([]*ClusterStatus, 0)
	} else {
		slice = make([]*ClusterStatus, len(l.items))
		copy(slice, l.items)
	}
	return slice
}

// Each runs the given function for each item of the list, in order. If the function
// returns false the iteration stops, otherwise it continues till all the elements
// of the list have been processed.
func (l *ClusterStatusList) Each(f func(item *ClusterStatus) bool) {
	if l == nil {
		return
	}
	for _, item := range l.items {
		if !f(item) {
			break
		}
	}
}

// Range runs the given function for each index and item of the list, in order. If
// the function returns false the iteration stops, otherwise it continues till all
// the elements of the list have been processed.
func (l *ClusterStatusList) Range(f func(index int, item *ClusterStatus) bool) {
	if l == nil {
		return
	}
	for index, item := range l.items {
		if !f(index, item) {
			break
		}
	}
}
