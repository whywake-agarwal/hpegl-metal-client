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
import (
	"time"
)
// Host struct for Host
type Host struct {
	// Unique ID for the resource instance as generated by the Metal service
	ID string `json:"ID,omitempty"`
	// Used to determine whether the DB entry has changed since it was last read.  should be returned unchanged on any update operation.
	ETag string `json:"ETag,omitempty"`
	// Common name for the resource instance. Must be 128 or fewer printable characters
	Name string `json:"Name,omitempty"`
	// Time when the resource was created in the database
	Created time.Time `json:"Created,omitempty"`
	// Time when the resource was last modified in the database
	Modified time.Time `json:"Modified,omitempty"`
	Description string `json:"Description,omitempty"`
	// The image service identifier used to image the server. ServiceID is one of those listed by the Images array returned as part of the get /available-resources call.
	ServiceID string `json:"ServiceID,omitempty"`
	// Overall flavor of server image used to image the server
	ServiceFlavor string `json:"ServiceFlavor,omitempty"`
	// Version of the ServiceFlavor used to image the server
	ServiceVersion string `json:"ServiceVersion,omitempty"`
	// The location of the machine assigned to the host.  LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call.
	LocationID string `json:"LocationID,omitempty"`
	// Name of the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call.
	MachineSizeName string `json:"MachineSizeName,omitempty"`
	// UniqueID referring to the machine size used to identify and select the machine assigned to the host.  MachineSizes are described by the MachineSize array returned by the get /available-resources call.
	MachineSizeID string `json:"MachineSizeID,omitempty"`
	// UniqueID referring to the machine on which this host is running.
	MachineID string `json:"MachineID,omitempty"`
	// IDs of SSH Keys used when configuring the Host
	SSHKeyIDs []string `json:"SSHKeyIDs,omitempty"`
	// Specific SSH keys that were when configuring the host.
	SSHAuthorizedKeys []string `json:"SSHAuthorizedKeys,omitempty"`
	// The list of IDs corresponding to the networks that were provisioned to the host. These networks are among those listed in the Networks array returned by the get /available-resources call.
	NetworkIDs []string `json:"NetworkIDs,omitempty"`
	// The host default network ID
	NetworkForDefaultRoute string `json:"NetworkForDefaultRoute,omitempty"`
	// User-provided data attached to the image configuration data when the host was provisioned
	UserData string `json:"UserData,omitempty"`
	// User-provided data to represent the identity of the host within an application environment. For example, this could be set to represent the Kubernetes node ID if the host is provisioned as a Kubernetes node.
	NodeID string `json:"NodeID,omitempty"`
	ISCSIConfig HostIscsiConfig `json:"ISCSIConfig,omitempty"`
	// Details describing host network connections
	Connections []HostConnection `json:"Connections,omitempty"`
	// True if the Host has been deleted.
	Deleted bool `json:"Deleted,omitempty"`
	// Describes if the portal is in active communication to the device
	PortalCommOkay bool `json:"PortalCommOkay,omitempty"`
	PowerStatus HostPowerState `json:"PowerStatus,omitempty"`
	State HostState `json:"State,omitempty"`
	Substate HostSubstate `json:"Substate,omitempty"`
	StateTime time.Time `json:"StateTime,omitempty"`
	SubstateTime time.Time `json:"SubstateTime,omitempty"`
	Progress int64 `json:"Progress,omitempty"`
	Alert bool `json:"Alert,omitempty"`
	AlertInfo []HostAlertInfo `json:"AlertInfo,omitempty"`
}
