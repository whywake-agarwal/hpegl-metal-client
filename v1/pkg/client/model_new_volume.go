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
// NewVolume struct for NewVolume
type NewVolume struct {
	Name string `json:"Name"`
	Description string `json:"Description,omitempty"`
	// Adds a new volume to the project.  This object requires the LocationID and is used when a new volume is created independently from the host creation therefore requiring a specified location.
	FlavorID string `json:"FlavorID"`
	// The size of the volume in GiB
	Capacity int64 `json:"Capacity"`
	// Indicates if the volume can be attached to multiple hosts
	Shareable bool `json:"Shareable,omitempty"`
	// The location of the volume (and the storage array) LocationID is one of those listed by the LocationInfo array returned as part of the get /available-resources call. Any volumes must be in the same location as their attached Host.
	LocationID string `json:"LocationID"`
	// The map of service/user specified label name to label value for this volume. Setting service labels is restricted by role.
	Labels map[string]string `json:"Labels,omitempty"`
	// The storage pool is one of those listed by the StoragePools array returned as part of the get /available-resources call that are available to create volumes of the specified flavor and location.
	StoragePoolID string `json:"StoragePoolID,omitempty"`
}
