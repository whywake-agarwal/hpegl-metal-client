# AddVolume

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | [optional] 
**Description** | **string** |  | [optional] 
**FlavorID** | **string** | This object is used for new volume creation in the host create operation. Since host create already has location ID that field is not provided in this object. | [optional] 
**Capacity** | **int64** | The size of the volume in GiB | [optional] 
**Shareable** | **bool** | Indicates if the volume can be attached to multiple hosts | [optional] 
**StoragePoolID** | **string** | The storage pool is one of those listed by the StoragePools array returned as part of the get /available-resources call that are available to create volumes of the specified flavor and location. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


