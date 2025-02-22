# VolumeUsageEntry

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ProjectID** | **string** | Project ID that contained the host | 
**LocationID** | **string** | The location ID is the data center location of the resource.  The LocationID must be one of those provided by the LocationInfo array returned as part of the get /available-resources call.  The locations are typically described by country, region, and data center. | 
**Allocated** | [**time.Time**](time.Time.md) | Timestamp of when resource (machine or storage) was allocated | 
**Ready** | [**time.Time**](time.Time.md) | Timestamp of when resource (host or volume) was ready for use | 
**Freed** | [**time.Time**](time.Time.md) | Timestamp of when resource (machine or storage) was freed | 
**UsageStart** | [**time.Time**](time.Time.md) | The start of the usage reporting window or when the resource was allocated | 
**UsageEnd** | [**time.Time**](time.Time.md) | The end of the usage reporting window or when the resource was freed | 
**UsageHours** | **int64** | The difference between the UsageEnd and UsageStart rounded up to the UsageHours | 
**Error** | **string** | Description of error that affected the usage reporting | 
**VolumeName** | **string** | Name of the volume | 
**VolumeID** | **string** | Unique ID of the volume | 
**FlavorName** | **string** | Name of the volume flavor used when creating the volume | 
**FlavorID** | **string** | Unique ID of the volume flavor used when creating the volume | 
**Capacity** | **int64** | The size of the volume in MB | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


