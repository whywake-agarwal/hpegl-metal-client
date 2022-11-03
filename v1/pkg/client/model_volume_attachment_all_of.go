// (C) Copyright 2021-2022 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// VolumeAttachmentAllOf struct for VolumeAttachmentAllOf
type VolumeAttachmentAllOf struct {
	// Unique ID of the volume attached to the host
	VolumeID string `json:"VolumeID,omitempty"`
	// Unique ID of the host connected to the volume
	HostID string `json:"HostID,omitempty"`
	// Host IP address for the network connection that connects to the storage array
	HostIPAddress string `json:"HostIPAddress,omitempty"`
	// IQN is the full initiator name used for identification during iSCSI login
	IQN string `json:"IQN,omitempty"`
	// CHAPSecret is the Challenge Authentication Protocol secret to be shared between array and initiator.
	CHAPSecret string `json:"CHAPSecret,omitempty"`
	// CHAPUserName is the CHAP username to use for CHAP authentication
	CHAPUserName string `json:"CHAPUserName,omitempty"`
	// LUN is the Logical Unit Number to be assigned to the volume on export
	LUN int32 `json:"LUN,omitempty"`
	// VolumeTargetIQN is the iQN for the volume, assigned by the array correspnding to the volume
	VolumeTargetIQN string `json:"VolumeTargetIQN,omitempty"`
	// VolumeTargetIPAddress is the IPV4 address of the iSCSI volume export
	VolumeTargetIPAddress string `json:"VolumeTargetIPAddress,omitempty"`
	State VaStateEnum `json:"State,omitempty"`
	// File share specific configuration parameters
	FSConfig *VafsConfig `json:"FSConfig,omitempty"`
}
