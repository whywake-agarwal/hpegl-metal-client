// (C) Copyright 2021 Hewlett Packard Enterprise Development LP

//

package client

import (
	_context "context"
	_nethttp "net/http"
)

// VolumeAttachmentsAPI defines the client functions provided for VolumeAttachments.
type VolumeAttachmentsAPI interface {
	/*
	   Add Create a new VolumeAttachment
	   Adds a new VolumeAttachment which enables a machine to use a volume.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param newVolumeAttachment NewVolumeAttachement parameters to create a new VolumeAttachment.
	   @return VolumeAttachment
	*/
	Add(ctx _context.Context, newVolumeAttachment NewVolumeAttachment) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   Delete Delete a VolumeAttachment
	   Deletes the VolumeAttachment with the matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param attachmentId ID of VolumeAttachment to delete
	*/
	Delete(ctx _context.Context, attachmentId string) (*_nethttp.Response, error)
	/*
	   GetByID Retrieve volume attachment by ID
	   Returns a single volume attachment with matching ID
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	    * @param attachmentId ID of volume attachment to return
	   @return VolumeAttachment
	*/
	GetByID(ctx _context.Context, attachmentId string) (VolumeAttachment, *_nethttp.Response, error)
	/*
	   List List all volume attachments in project
	   Returns an array of all VolumeAttachments defined within the project.
	    * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	   @return []VolumeAttachment
	*/
	List(ctx _context.Context) ([]VolumeAttachment, *_nethttp.Response, error)
}
