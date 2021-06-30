// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * Quake Project Client API
 *
 * This Quake client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes. 
 *
 * API version: 1.2.1
 * Contact: chuck.hudson@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// NewHost struct for NewHost
type NewHost struct {
	// The host name used when installing the host operating system.  Note that some OS implementations may require the name to be formatted as a fully qualified domain name.
	Name string `json:"Name"`
	Description string `json:"Description,omitempty"`
	// The image service identifier must be identifer of one of the available imaging services provided by the AvailableImage array returned as part of the get /available-resources call. Images are typically described by category (e.g. Linux), flavor (e.g. ubuntu) and version (e.g. 18.04-20190807)
	ServiceID string `json:"ServiceID"`
	// The location ID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center. The location ID must also be one that has sufficient inventory for the selected MachineSize.  See the MachineInventory array as returned by the get /available-resources call to select a location that has machines available with the correct machine size.
	LocationID string `json:"LocationID"`
	// The machine size ID must be one of those listed in the MachineSizes array returned as part of the get /available-resources call.  This array provides the name and detailed description for each machine size.  Use the MachineInventory array provided by the get /available-resources call to find a location with an adequate inventory of machines with the desired machine size.
	MachineSizeID string `json:"MachineSizeID"`
	// The machine ID is normally not used.
	MachineID string `json:"MachineID,omitempty"`
	// These IDs must correspond to the IDs for SSH keys already created within the context of the project.  These SSH keys and those included as SSHAuthorizedKeys will be used together to provision SSH keys when the machine is imaged.
	SSHKeyIDs []string `json:"SSHKeyIDs,omitempty"`
	// The list of IDs corresponding to the networks that will be provisioned to the host. These networks must be among those listed in the Networks array returned by the get /available-resources call.  This list must include the ID for any network identified by the HostUseEnum as Required.
	NetworkIDs []string `json:"NetworkIDs"`
	// The list of IDs corresponding to existing, unattached volumes that should be attached to the new host.  The volume must be one of those listed in the Volumes array returned by the get /available-resources call.  In addition, the volume must be in the visible state and not a part of any current VolumeAttachment (see get volume-attachments)
	VolumeIDs []string `json:"VolumeIDs,omitempty"`
	// New volumes may be created and connected to the Host when the host is provisioned. The information provided here to create a host is the same as required when doing a post /volumes call
	NewVolumes []AddVolume `json:"NewVolumes,omitempty"`
	// User-provided data to be attached to the image configuration data.
	UserData string `json:"UserData,omitempty"`
	// User-provided data to represent the identity of the host within an application environment. For example, this could be set to represent the Kubernetes node ID if the host is provisioned as a Kubernetes node.
	NodeID string `json:"NodeID,omitempty"`
}
