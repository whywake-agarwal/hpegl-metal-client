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
// Netmask the model 'Netmask'
type Netmask string

// List of Netmask
const (
	NETMASK__8 Netmask = "/8"
	NETMASK__9 Netmask = "/9"
	NETMASK__10 Netmask = "/10"
	NETMASK__11 Netmask = "/11"
	NETMASK__12 Netmask = "/12"
	NETMASK__13 Netmask = "/13"
	NETMASK__14 Netmask = "/14"
	NETMASK__15 Netmask = "/15"
	NETMASK__16 Netmask = "/16"
	NETMASK__17 Netmask = "/17"
	NETMASK__18 Netmask = "/18"
	NETMASK__19 Netmask = "/19"
	NETMASK__20 Netmask = "/20"
	NETMASK__21 Netmask = "/21"
	NETMASK__22 Netmask = "/22"
	NETMASK__23 Netmask = "/23"
	NETMASK__24 Netmask = "/24"
	NETMASK__25 Netmask = "/25"
	NETMASK__26 Netmask = "/26"
	NETMASK__27 Netmask = "/27"
	NETMASK__28 Netmask = "/28"
	NETMASK__29 Netmask = "/29"
	NETMASK__30 Netmask = "/30"
	NETMASK__31 Netmask = "/31"
	NETMASK__104 Netmask = "/104"
	NETMASK__105 Netmask = "/105"
	NETMASK__106 Netmask = "/106"
	NETMASK__107 Netmask = "/107"
	NETMASK__108 Netmask = "/108"
	NETMASK__109 Netmask = "/109"
	NETMASK__110 Netmask = "/110"
	NETMASK__111 Netmask = "/111"
	NETMASK__112 Netmask = "/112"
	NETMASK__113 Netmask = "/113"
	NETMASK__114 Netmask = "/114"
	NETMASK__115 Netmask = "/115"
	NETMASK__116 Netmask = "/116"
	NETMASK__117 Netmask = "/117"
	NETMASK__118 Netmask = "/118"
	NETMASK__119 Netmask = "/119"
	NETMASK__120 Netmask = "/120"
	NETMASK__121 Netmask = "/121"
	NETMASK__122 Netmask = "/122"
	NETMASK__123 Netmask = "/123"
	NETMASK__124 Netmask = "/124"
	NETMASK__125 Netmask = "/125"
	NETMASK__126 Netmask = "/126"
	NETMASK__127 Netmask = "/127"
)
