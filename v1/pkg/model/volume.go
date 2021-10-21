// Copyright 2016-2021 Hewlett Packard Enterprise Development LP

package model

// NewVolume is used as request body during Volume create only
type NewVolume struct {
	Name        string
	Description string // Description is some human readable information.
	FlavorID    string // Flavor of volume created
	Capacity    uint64 // Requested size of the volume in MB.
	LocationID  string // ID of the associated pod / data center
}

// AddVolume is used as part of NewHost to allow creation of Volumes
// at the same time a Host is created
type AddVolume struct {
	Name        string
	Description string // Description is some human readable information.
	FlavorID    string // Flavor of volume created
	Capacity    uint64 // Requested size of the volume in MB.
}

// Volume represents an external array-based disk volume.
// Return body for create, get, list, update and as update request body
type Volume struct {
	ResourceBase
	Description string       // Description is some human readable information.
	LocationID  string       // PodID is the ID of the associated pod
	Capacity    uint64       // Capacity is the requested size of the volume in MB.
	FlavorID    string       // Flavor of volume created
	State       VolumeState  // State of the volume, managed by the portal volmon bot.
	Status      VolumeStatus // Status of the volume, updated by the storage controller.
}

// VolumeInfo represents a minimal description of an already existing Volume for
// use by AvailableResources struct.  This is to allow a user to attach an
// existing volume to a new host.
type VolumeInfo struct {
	ID          string
	Name        string
	Description string
	FlavorID    string
	Capacity    uint64 // Volume size in MB
	LocationID  string
	State       VolumeState  // State of the volume, managed by the portal volmon bot.
	Status      VolumeStatus // Status of the volume, updated by the storage controller.
}

// VolumeState all posible volume states stored in the  model.
type VolumeState string

// VolumeState enum values
const (
	VolumeStateNew        VolumeState = "new"        // Volume is created in the portal/database.
	VolumeStateAllocating VolumeState = "allocating" // Volume is being allocated on the array.
	VolumeStateAllocated  VolumeState = "allocated"  // Volume is allocated on the array and ready for an attachment.
	VolumeStateVisible    VolumeState = "visible"    // Volume is visible from the array to an initiator.
	VolumeStateDeleting   VolumeState = "deleting"   // Volume is being deleted on the array.
	VolumeStateDeleted    VolumeState = "deleted"    // Volume is no longer on the array.
	VolumeStateFailed     VolumeState = "failed"     // Volume is not useable and needs to be deleted.
)

// VolumeStatus status of the volume stored in the model.
type VolumeStatus string

// VolumeStatus enum values
const (
	VolumeStatusOK       VolumeStatus = "ok"       // VolumeStatusOK the array detects no issues with the volume.
	VolumeStatusError    VolumeStatus = "error"    // VolumeStatusError the volume is inaccessible.
	VolumeStatusDegraged VolumeStatus = "degraded" // VolumeStatusDegraged the volume is degraded but still accessible.
)
