// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

/*
 * Quake Project Client API
 *
 * This Quake client REST API provides access to bare metal as-a-service (BMaaS) within a single project context.  Clients are able to create fully-provisioned hosts, storage volumes, and project-specific private networks in an isolated project environment.  Project-owned resources that can be accessed via this API include... Host, Volume, VolumeAttachment, Network (project private), and SSH Key.    Each API call is done within a single project context.  The specific Project identifier must be provided within the header of each REST call. The server will validate that the provided authentication credentials (JWTs) are valid for the referenced project before any operation is performed.  If a single credential is valid for multiple projects, the client must still reference a single project in the header each API call.  Clients can also access information about available services and resources through the AvailableResources object.  This object provides detailed  information about the OS imaging options, the machine size options, the storage volume options, data center locations, and such that are needed when creating hosts and volumes. 
 *
 * API version: 1.3.1
 * Contact: quake-core@hpe.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package client
// HostSubstate Host substate within HostState
type HostSubstate string

// List of HostSubstate
const (
	HOSTSUBSTATE_EMPTY HostSubstate = ""
	HOSTSUBSTATE_INIT HostSubstate = "Init"
	HOSTSUBSTATE_COMPLETE HostSubstate = "Complete"
	HOSTSUBSTATE_FAILED HostSubstate = "Failed"
	HOSTSUBSTATE_WORKING HostSubstate = "Working"
	HOSTSUBSTATE_NEED_MACHINE HostSubstate = "Need-Machine"
	HOSTSUBSTATE_NEED_I_PS HostSubstate = "Need-IPs"
	HOSTSUBSTATE_INIT_OFF HostSubstate = "Init-Off"
	HOSTSUBSTATE_SETTING_PXE HostSubstate = "Setting-PXE"
	HOSTSUBSTATE_CLEARING_LOG HostSubstate = "Clearing-Log"
	HOSTSUBSTATE_STARTING_DEPLOY HostSubstate = "Starting-Deploy"
	HOSTSUBSTATE_BLOCKED HostSubstate = "Blocked"
	HOSTSUBSTATE_DEPLOY_POWERON HostSubstate = "Deploy-Poweron"
	HOSTSUBSTATE_DEPLOYING HostSubstate = "Deploying"
	HOSTSUBSTATE_CONFIRM_OFF HostSubstate = "Confirm-Off"
	HOSTSUBSTATE_SNAP_LOG HostSubstate = "SnapLog"
	HOSTSUBSTATE_FAIL_SNAP_LOG HostSubstate = "Fail-SnapLog"
	HOSTSUBSTATE_SETTING_DISK_BOOT HostSubstate = "Setting-DiskBoot"
	HOSTSUBSTATE_CONNECT HostSubstate = "Connect"
	HOSTSUBSTATE_CONNECT_POWERON HostSubstate = "Connect-Poweron"
	HOSTSUBSTATE_DISCONNECT HostSubstate = "Disconnect"
	HOSTSUBSTATE_ABORT_DEPLOY HostSubstate = "Abort-Deploy"
	HOSTSUBSTATE_DELETE_POWEROFF HostSubstate = "Delete-Poweroff"
	HOSTSUBSTATE_RELEASING HostSubstate = "Releasing"
)
