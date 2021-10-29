// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake Metal Client API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes.    Note: All URIs are relative to https://<metal_service_url>/rest/v1 
 *
 * API version: 1.3.5
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// NewIpPool struct for NewIpPool
type NewIpPool struct {
	// Name for the IP pool
	Name string `json:"Name,omitempty"`
	// Description for the IP pool
	Description string `json:"Description,omitempty"`
	IPVersion IpVer `json:"IPVersion,omitempty"`
	// Base address of the IP pool
	BaseIP string `json:"BaseIP,omitempty"`
	Netmask Netmask `json:"Netmask,omitempty"`
	// Default route associated with the IP pool
	DefaultRoute string `json:"DefaultRoute,omitempty"`
	Sources []IpSource `json:"Sources,omitempty"`
	// List of DNS servers for the IP pool
	DNS []string `json:"DNS,omitempty"`
	// Optional web-proxy for external internet access should the pool actually be behind a firewall
	Proxy string `json:"Proxy,omitempty"`
	// Addresses or CIDRs for which proxy requests are not made
	NoProxy string `json:"NoProxy,omitempty"`
	// List of NTP servers for the IP pool
	NTP []string `json:"NTP,omitempty"`
}
