// (C) Copyright 2021-2023 Hewlett Packard Enterprise Development LP

/*
 * HPE GreenLake for bare metal API
 *
 * This Metal Client REST API provides access to bare metal as-a-service (BMaaS) within a single project context. Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.   Project-owned resources that can be accessed via this API include - Host,  Volume, VolumeAttachment, Network (project private), and SSH Key. Each API  call is done within a single project context. The specific Project identifier  must be provided within the header of for each REST call. The server will  validate that the provided authentication credentials (JWTs) are valid  for  the referenced project before any operation is performed. If a single credential is valid for multiple projects, the client must still reference a single project  in the header for each API call.    Clients can also access information about available services and resources through the AvailableResources object. This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, and data center locations which are needed when creating hosts and volumes.  
 *
 * API version: v1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// AddVolume struct for AddVolume
type AddVolume struct {
	Name string `json:"Name,omitempty"`
	Description string `json:"Description,omitempty"`
	// This object is used for new volume creation in the host create operation. Since host create already has location ID that field is not provided in this object.
	FlavorID string `json:"FlavorID,omitempty"`
	// The size of the volume in GiB
	Capacity int64 `json:"Capacity,omitempty"`
	// Indicates if the volume can be attached to multiple hosts
	Shareable bool `json:"Shareable,omitempty"`
	// The storage pool is one of those listed by the StoragePools array returned as part of the get /available-resources call that are available to create volumes of the specified flavor and location.
	StoragePoolID string `json:"StoragePoolID,omitempty"`
}
