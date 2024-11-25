// Code generated by go-swagger; DO NOT EDIT.

package storage

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewSnapshotCollectionGetParams creates a new SnapshotCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewSnapshotCollectionGetParams() *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewSnapshotCollectionGetParamsWithTimeout creates a new SnapshotCollectionGetParams object
// with the ability to set a timeout on a request.
func NewSnapshotCollectionGetParamsWithTimeout(timeout time.Duration) *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		timeout: timeout,
	}
}

// NewSnapshotCollectionGetParamsWithContext creates a new SnapshotCollectionGetParams object
// with the ability to set a context for a request.
func NewSnapshotCollectionGetParamsWithContext(ctx context.Context) *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		Context: ctx,
	}
}

// NewSnapshotCollectionGetParamsWithHTTPClient creates a new SnapshotCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewSnapshotCollectionGetParamsWithHTTPClient(client *http.Client) *SnapshotCollectionGetParams {
	return &SnapshotCollectionGetParams{
		HTTPClient: client,
	}
}

/*
SnapshotCollectionGetParams contains all the parameters to send to the API endpoint

	for the snapshot collection get operation.

	Typically these are written to a http.Request.
*/
type SnapshotCollectionGetParams struct {

	/* Comment.

	   Filter by comment
	*/
	Comment *string

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTime *string

	/* DeltaSizeConsumed.

	   Filter by delta.size_consumed
	*/
	DeltaSizeConsumed *int64

	/* DeltaTimeElapsed.

	   Filter by delta.time_elapsed
	*/
	DeltaTimeElapsed *string

	/* ExpiryTime.

	   Filter by expiry_time
	*/
	ExpiryTime *string

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* LogicalSize.

	   Filter by logical_size
	*/
	LogicalSize *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* Name.

	   Filter by name
	*/
	Name *string

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

	/* Owners.

	   Filter by owners
	*/
	Owners *string

	/* ProvenanceVolumeUUID.

	   Filter by provenance_volume.uuid
	*/
	ProvenanceVolumeUUID *string

	/* ReclaimableSpace.

	   Filter by reclaimable_space
	*/
	ReclaimableSpace *int64

	/* ReturnRecords.

	   The default is true for GET calls.  When set to false, only the number of records is returned.

	   Default: true
	*/
	ReturnRecords *bool

	/* ReturnTimeout.

	   The number of seconds to allow the call to execute before returning.  When iterating over a collection, the default is 15 seconds.  ONTAP returns earlier if either max records or the end of the collection is reached.

	   Default: 15
	*/
	ReturnTimeout *int64

	/* Size.

	   Filter by size
	*/
	Size *int64

	/* SnaplockExpired.

	   Filter by snaplock.expired
	*/
	SnaplockExpired *bool

	/* SnaplockExpiryTime.

	   Filter by snaplock.expiry_time
	*/
	SnaplockExpiryTime *string

	/* SnaplockTimeUntilExpiry.

	   Filter by snaplock.time_until_expiry
	*/
	SnaplockTimeUntilExpiry *string

	/* SnapmirrorLabel.

	   Filter by snapmirror_label
	*/
	SnapmirrorLabel *string

	/* State.

	   Filter by state
	*/
	State *string

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	/* VersionUUID.

	   Filter by version_uuid
	*/
	VersionUUID *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeName *string

	/* VolumeUUID.

	   Volume
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the snapshot collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotCollectionGetParams) WithDefaults() *SnapshotCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the snapshot collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *SnapshotCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)
	)

	val := SnapshotCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithTimeout(timeout time.Duration) *SnapshotCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithContext(ctx context.Context) *SnapshotCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithHTTPClient(client *http.Client) *SnapshotCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithComment adds the comment to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithComment(comment *string) *SnapshotCollectionGetParams {
	o.SetComment(comment)
	return o
}

// SetComment adds the comment to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetComment(comment *string) {
	o.Comment = comment
}

// WithCreateTime adds the createTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithCreateTime(createTime *string) *SnapshotCollectionGetParams {
	o.SetCreateTime(createTime)
	return o
}

// SetCreateTime adds the createTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetCreateTime(createTime *string) {
	o.CreateTime = createTime
}

// WithDeltaSizeConsumed adds the deltaSizeConsumed to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithDeltaSizeConsumed(deltaSizeConsumed *int64) *SnapshotCollectionGetParams {
	o.SetDeltaSizeConsumed(deltaSizeConsumed)
	return o
}

// SetDeltaSizeConsumed adds the deltaSizeConsumed to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetDeltaSizeConsumed(deltaSizeConsumed *int64) {
	o.DeltaSizeConsumed = deltaSizeConsumed
}

// WithDeltaTimeElapsed adds the deltaTimeElapsed to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithDeltaTimeElapsed(deltaTimeElapsed *string) *SnapshotCollectionGetParams {
	o.SetDeltaTimeElapsed(deltaTimeElapsed)
	return o
}

// SetDeltaTimeElapsed adds the deltaTimeElapsed to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetDeltaTimeElapsed(deltaTimeElapsed *string) {
	o.DeltaTimeElapsed = deltaTimeElapsed
}

// WithExpiryTime adds the expiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithExpiryTime(expiryTime *string) *SnapshotCollectionGetParams {
	o.SetExpiryTime(expiryTime)
	return o
}

// SetExpiryTime adds the expiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetExpiryTime(expiryTime *string) {
	o.ExpiryTime = expiryTime
}

// WithFields adds the fields to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithFields(fields []string) *SnapshotCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithLogicalSize adds the logicalSize to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithLogicalSize(logicalSize *int64) *SnapshotCollectionGetParams {
	o.SetLogicalSize(logicalSize)
	return o
}

// SetLogicalSize adds the logicalSize to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetLogicalSize(logicalSize *int64) {
	o.LogicalSize = logicalSize
}

// WithMaxRecords adds the maxRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithMaxRecords(maxRecords *int64) *SnapshotCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithName adds the name to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithName(name *string) *SnapshotCollectionGetParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetName(name *string) {
	o.Name = name
}

// WithOrderBy adds the orderBy to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithOrderBy(orderBy []string) *SnapshotCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithOwners adds the owners to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithOwners(owners *string) *SnapshotCollectionGetParams {
	o.SetOwners(owners)
	return o
}

// SetOwners adds the owners to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetOwners(owners *string) {
	o.Owners = owners
}

// WithProvenanceVolumeUUID adds the provenanceVolumeUUID to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithProvenanceVolumeUUID(provenanceVolumeUUID *string) *SnapshotCollectionGetParams {
	o.SetProvenanceVolumeUUID(provenanceVolumeUUID)
	return o
}

// SetProvenanceVolumeUUID adds the provenanceVolumeUuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetProvenanceVolumeUUID(provenanceVolumeUUID *string) {
	o.ProvenanceVolumeUUID = provenanceVolumeUUID
}

// WithReclaimableSpace adds the reclaimableSpace to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithReclaimableSpace(reclaimableSpace *int64) *SnapshotCollectionGetParams {
	o.SetReclaimableSpace(reclaimableSpace)
	return o
}

// SetReclaimableSpace adds the reclaimableSpace to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetReclaimableSpace(reclaimableSpace *int64) {
	o.ReclaimableSpace = reclaimableSpace
}

// WithReturnRecords adds the returnRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithReturnRecords(returnRecords *bool) *SnapshotCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *SnapshotCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSize adds the size to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSize(size *int64) *SnapshotCollectionGetParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSize(size *int64) {
	o.Size = size
}

// WithSnaplockExpired adds the snaplockExpired to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSnaplockExpired(snaplockExpired *bool) *SnapshotCollectionGetParams {
	o.SetSnaplockExpired(snaplockExpired)
	return o
}

// SetSnaplockExpired adds the snaplockExpired to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSnaplockExpired(snaplockExpired *bool) {
	o.SnaplockExpired = snaplockExpired
}

// WithSnaplockExpiryTime adds the snaplockExpiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSnaplockExpiryTime(snaplockExpiryTime *string) *SnapshotCollectionGetParams {
	o.SetSnaplockExpiryTime(snaplockExpiryTime)
	return o
}

// SetSnaplockExpiryTime adds the snaplockExpiryTime to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSnaplockExpiryTime(snaplockExpiryTime *string) {
	o.SnaplockExpiryTime = snaplockExpiryTime
}

// WithSnaplockTimeUntilExpiry adds the snaplockTimeUntilExpiry to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSnaplockTimeUntilExpiry(snaplockTimeUntilExpiry *string) *SnapshotCollectionGetParams {
	o.SetSnaplockTimeUntilExpiry(snaplockTimeUntilExpiry)
	return o
}

// SetSnaplockTimeUntilExpiry adds the snaplockTimeUntilExpiry to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSnaplockTimeUntilExpiry(snaplockTimeUntilExpiry *string) {
	o.SnaplockTimeUntilExpiry = snaplockTimeUntilExpiry
}

// WithSnapmirrorLabel adds the snapmirrorLabel to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSnapmirrorLabel(snapmirrorLabel *string) *SnapshotCollectionGetParams {
	o.SetSnapmirrorLabel(snapmirrorLabel)
	return o
}

// SetSnapmirrorLabel adds the snapmirrorLabel to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSnapmirrorLabel(snapmirrorLabel *string) {
	o.SnapmirrorLabel = snapmirrorLabel
}

// WithState adds the state to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithState(state *string) *SnapshotCollectionGetParams {
	o.SetState(state)
	return o
}

// SetState adds the state to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetState(state *string) {
	o.State = state
}

// WithSvmName adds the svmName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSvmName(svmName *string) *SnapshotCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithSvmUUID(svmUUID *string) *SnapshotCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithUUID(uuid *string) *SnapshotCollectionGetParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WithVersionUUID adds the versionUUID to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithVersionUUID(versionUUID *string) *SnapshotCollectionGetParams {
	o.SetVersionUUID(versionUUID)
	return o
}

// SetVersionUUID adds the versionUuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetVersionUUID(versionUUID *string) {
	o.VersionUUID = versionUUID
}

// WithVolumeName adds the volumeName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithVolumeName(volumeName *string) *SnapshotCollectionGetParams {
	o.SetVolumeName(volumeName)
	return o
}

// SetVolumeName adds the volumeName to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetVolumeName(volumeName *string) {
	o.VolumeName = volumeName
}

// WithVolumeUUID adds the volumeUUID to the snapshot collection get params
func (o *SnapshotCollectionGetParams) WithVolumeUUID(volumeUUID string) *SnapshotCollectionGetParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the snapshot collection get params
func (o *SnapshotCollectionGetParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *SnapshotCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Comment != nil {

		// query param comment
		var qrComment string

		if o.Comment != nil {
			qrComment = *o.Comment
		}
		qComment := qrComment
		if qComment != "" {

			if err := r.SetQueryParam("comment", qComment); err != nil {
				return err
			}
		}
	}

	if o.CreateTime != nil {

		// query param create_time
		var qrCreateTime string

		if o.CreateTime != nil {
			qrCreateTime = *o.CreateTime
		}
		qCreateTime := qrCreateTime
		if qCreateTime != "" {

			if err := r.SetQueryParam("create_time", qCreateTime); err != nil {
				return err
			}
		}
	}

	if o.DeltaSizeConsumed != nil {

		// query param delta.size_consumed
		var qrDeltaSizeConsumed int64

		if o.DeltaSizeConsumed != nil {
			qrDeltaSizeConsumed = *o.DeltaSizeConsumed
		}
		qDeltaSizeConsumed := swag.FormatInt64(qrDeltaSizeConsumed)
		if qDeltaSizeConsumed != "" {

			if err := r.SetQueryParam("delta.size_consumed", qDeltaSizeConsumed); err != nil {
				return err
			}
		}
	}

	if o.DeltaTimeElapsed != nil {

		// query param delta.time_elapsed
		var qrDeltaTimeElapsed string

		if o.DeltaTimeElapsed != nil {
			qrDeltaTimeElapsed = *o.DeltaTimeElapsed
		}
		qDeltaTimeElapsed := qrDeltaTimeElapsed
		if qDeltaTimeElapsed != "" {

			if err := r.SetQueryParam("delta.time_elapsed", qDeltaTimeElapsed); err != nil {
				return err
			}
		}
	}

	if o.ExpiryTime != nil {

		// query param expiry_time
		var qrExpiryTime string

		if o.ExpiryTime != nil {
			qrExpiryTime = *o.ExpiryTime
		}
		qExpiryTime := qrExpiryTime
		if qExpiryTime != "" {

			if err := r.SetQueryParam("expiry_time", qExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.LogicalSize != nil {

		// query param logical_size
		var qrLogicalSize int64

		if o.LogicalSize != nil {
			qrLogicalSize = *o.LogicalSize
		}
		qLogicalSize := swag.FormatInt64(qrLogicalSize)
		if qLogicalSize != "" {

			if err := r.SetQueryParam("logical_size", qLogicalSize); err != nil {
				return err
			}
		}
	}

	if o.MaxRecords != nil {

		// query param max_records
		var qrMaxRecords int64

		if o.MaxRecords != nil {
			qrMaxRecords = *o.MaxRecords
		}
		qMaxRecords := swag.FormatInt64(qrMaxRecords)
		if qMaxRecords != "" {

			if err := r.SetQueryParam("max_records", qMaxRecords); err != nil {
				return err
			}
		}
	}

	if o.Name != nil {

		// query param name
		var qrName string

		if o.Name != nil {
			qrName = *o.Name
		}
		qName := qrName
		if qName != "" {

			if err := r.SetQueryParam("name", qName); err != nil {
				return err
			}
		}
	}

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
		}
	}

	if o.Owners != nil {

		// query param owners
		var qrOwners string

		if o.Owners != nil {
			qrOwners = *o.Owners
		}
		qOwners := qrOwners
		if qOwners != "" {

			if err := r.SetQueryParam("owners", qOwners); err != nil {
				return err
			}
		}
	}

	if o.ProvenanceVolumeUUID != nil {

		// query param provenance_volume.uuid
		var qrProvenanceVolumeUUID string

		if o.ProvenanceVolumeUUID != nil {
			qrProvenanceVolumeUUID = *o.ProvenanceVolumeUUID
		}
		qProvenanceVolumeUUID := qrProvenanceVolumeUUID
		if qProvenanceVolumeUUID != "" {

			if err := r.SetQueryParam("provenance_volume.uuid", qProvenanceVolumeUUID); err != nil {
				return err
			}
		}
	}

	if o.ReclaimableSpace != nil {

		// query param reclaimable_space
		var qrReclaimableSpace int64

		if o.ReclaimableSpace != nil {
			qrReclaimableSpace = *o.ReclaimableSpace
		}
		qReclaimableSpace := swag.FormatInt64(qrReclaimableSpace)
		if qReclaimableSpace != "" {

			if err := r.SetQueryParam("reclaimable_space", qReclaimableSpace); err != nil {
				return err
			}
		}
	}

	if o.ReturnRecords != nil {

		// query param return_records
		var qrReturnRecords bool

		if o.ReturnRecords != nil {
			qrReturnRecords = *o.ReturnRecords
		}
		qReturnRecords := swag.FormatBool(qrReturnRecords)
		if qReturnRecords != "" {

			if err := r.SetQueryParam("return_records", qReturnRecords); err != nil {
				return err
			}
		}
	}

	if o.ReturnTimeout != nil {

		// query param return_timeout
		var qrReturnTimeout int64

		if o.ReturnTimeout != nil {
			qrReturnTimeout = *o.ReturnTimeout
		}
		qReturnTimeout := swag.FormatInt64(qrReturnTimeout)
		if qReturnTimeout != "" {

			if err := r.SetQueryParam("return_timeout", qReturnTimeout); err != nil {
				return err
			}
		}
	}

	if o.Size != nil {

		// query param size
		var qrSize int64

		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt64(qrSize)
		if qSize != "" {

			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}
	}

	if o.SnaplockExpired != nil {

		// query param snaplock.expired
		var qrSnaplockExpired bool

		if o.SnaplockExpired != nil {
			qrSnaplockExpired = *o.SnaplockExpired
		}
		qSnaplockExpired := swag.FormatBool(qrSnaplockExpired)
		if qSnaplockExpired != "" {

			if err := r.SetQueryParam("snaplock.expired", qSnaplockExpired); err != nil {
				return err
			}
		}
	}

	if o.SnaplockExpiryTime != nil {

		// query param snaplock.expiry_time
		var qrSnaplockExpiryTime string

		if o.SnaplockExpiryTime != nil {
			qrSnaplockExpiryTime = *o.SnaplockExpiryTime
		}
		qSnaplockExpiryTime := qrSnaplockExpiryTime
		if qSnaplockExpiryTime != "" {

			if err := r.SetQueryParam("snaplock.expiry_time", qSnaplockExpiryTime); err != nil {
				return err
			}
		}
	}

	if o.SnaplockTimeUntilExpiry != nil {

		// query param snaplock.time_until_expiry
		var qrSnaplockTimeUntilExpiry string

		if o.SnaplockTimeUntilExpiry != nil {
			qrSnaplockTimeUntilExpiry = *o.SnaplockTimeUntilExpiry
		}
		qSnaplockTimeUntilExpiry := qrSnaplockTimeUntilExpiry
		if qSnaplockTimeUntilExpiry != "" {

			if err := r.SetQueryParam("snaplock.time_until_expiry", qSnaplockTimeUntilExpiry); err != nil {
				return err
			}
		}
	}

	if o.SnapmirrorLabel != nil {

		// query param snapmirror_label
		var qrSnapmirrorLabel string

		if o.SnapmirrorLabel != nil {
			qrSnapmirrorLabel = *o.SnapmirrorLabel
		}
		qSnapmirrorLabel := qrSnapmirrorLabel
		if qSnapmirrorLabel != "" {

			if err := r.SetQueryParam("snapmirror_label", qSnapmirrorLabel); err != nil {
				return err
			}
		}
	}

	if o.State != nil {

		// query param state
		var qrState string

		if o.State != nil {
			qrState = *o.State
		}
		qState := qrState
		if qState != "" {

			if err := r.SetQueryParam("state", qState); err != nil {
				return err
			}
		}
	}

	if o.SvmName != nil {

		// query param svm.name
		var qrSvmName string

		if o.SvmName != nil {
			qrSvmName = *o.SvmName
		}
		qSvmName := qrSvmName
		if qSvmName != "" {

			if err := r.SetQueryParam("svm.name", qSvmName); err != nil {
				return err
			}
		}
	}

	if o.SvmUUID != nil {

		// query param svm.uuid
		var qrSvmUUID string

		if o.SvmUUID != nil {
			qrSvmUUID = *o.SvmUUID
		}
		qSvmUUID := qrSvmUUID
		if qSvmUUID != "" {

			if err := r.SetQueryParam("svm.uuid", qSvmUUID); err != nil {
				return err
			}
		}
	}

	if o.UUID != nil {

		// query param uuid
		var qrUUID string

		if o.UUID != nil {
			qrUUID = *o.UUID
		}
		qUUID := qrUUID
		if qUUID != "" {

			if err := r.SetQueryParam("uuid", qUUID); err != nil {
				return err
			}
		}
	}

	if o.VersionUUID != nil {

		// query param version_uuid
		var qrVersionUUID string

		if o.VersionUUID != nil {
			qrVersionUUID = *o.VersionUUID
		}
		qVersionUUID := qrVersionUUID
		if qVersionUUID != "" {

			if err := r.SetQueryParam("version_uuid", qVersionUUID); err != nil {
				return err
			}
		}
	}

	if o.VolumeName != nil {

		// query param volume.name
		var qrVolumeName string

		if o.VolumeName != nil {
			qrVolumeName = *o.VolumeName
		}
		qVolumeName := qrVolumeName
		if qVolumeName != "" {

			if err := r.SetQueryParam("volume.name", qVolumeName); err != nil {
				return err
			}
		}
	}

	// path param volume.uuid
	if err := r.SetPathParam("volume.uuid", o.VolumeUUID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamSnapshotCollectionGet binds the parameter fields
func (o *SnapshotCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
	fieldsIR := o.Fields

	var fieldsIC []string
	for _, fieldsIIR := range fieldsIR { // explode []string

		fieldsIIV := fieldsIIR // string as string
		fieldsIC = append(fieldsIC, fieldsIIV)
	}

	// items.CollectionFormat: "csv"
	fieldsIS := swag.JoinByFormat(fieldsIC, "csv")

	return fieldsIS
}

// bindParamSnapshotCollectionGet binds the parameter order_by
func (o *SnapshotCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
	orderByIR := o.OrderBy

	var orderByIC []string
	for _, orderByIIR := range orderByIR { // explode []string

		orderByIIV := orderByIIR // string as string
		orderByIC = append(orderByIC, orderByIIV)
	}

	// items.CollectionFormat: "csv"
	orderByIS := swag.JoinByFormat(orderByIC, "csv")

	return orderByIS
}
