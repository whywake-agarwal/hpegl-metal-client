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
import (
	"time"
)
// Network struct for Network
type Network struct {
	// Unique ID for the resource instance as generated by the Metal service
	ID string `json:"ID,omitempty"`
	// Used to determine whether the DB entry has changed since it was last read.  should be returned unchanged on any update operation.
	ETag string `json:"ETag,omitempty"`
	// Common name for the resource instance
	Name string `json:"Name,omitempty"`
	// Time when the resource was created in the database
	Created time.Time `json:"Created,omitempty"`
	// Time when the resource was last modified in the database
	Modified time.Time `json:"Modified,omitempty"`
	// The location ID is the data center location of the resource.  The LocationID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center.
	LocationID string `json:"LocationID,omitempty"`
	Description string `json:"Description,omitempty"`
	Kind NetworkKind `json:"Kind,omitempty"`
	HostUse NetworkHostUse `json:"HostUse,omitempty"`
	IPPoolID string `json:"IPPoolID,omitempty"`
}
