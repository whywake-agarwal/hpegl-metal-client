// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake Metal Client API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes. 
 *
 * API version: 1.3.5
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// AvailableResources struct for AvailableResources
type AvailableResources struct {
	// Array listing the available host imaging services
	Images []AvailableImage `json:"Images,omitempty"`
	// Array listing the data center locations with available resources
	Locations []LocationInfo `json:"Locations,omitempty"`
	// Array listing the networks available for host connections
	Networks []AvailableNetwork `json:"Networks,omitempty"`
	// Array listing the available machine (server) sizes
	MachineSizes []MachineSize `json:"MachineSizes,omitempty"`
	// Array listing the available volume flavors
	VolumeFlavors []VolumeFlavor `json:"VolumeFlavors,omitempty"`
	// Array listing the existing project volumes that could be attached to a host
	Volumes []VolumeInfo `json:"Volumes,omitempty"`
	// Array listing the number of machines of each size in each location
	MachineInventory []MachineInventory `json:"MachineInventory,omitempty"`
	// Array providing information on the amount of available storage of each flavor in each location
	StorageInventory []StorageInventory `json:"StorageInventory,omitempty"`
	// Array listing pre-defined SSH keys that could be referenced when creating a Host
	SSHKeys []SshKeyEntry `json:"SSHKeys,omitempty"`
}
