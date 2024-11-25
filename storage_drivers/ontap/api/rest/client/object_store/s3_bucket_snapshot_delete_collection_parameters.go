// Code generated by go-swagger; DO NOT EDIT.

package object_store

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

// NewS3BucketSnapshotDeleteCollectionParams creates a new S3BucketSnapshotDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewS3BucketSnapshotDeleteCollectionParams() *S3BucketSnapshotDeleteCollectionParams {
	return &S3BucketSnapshotDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewS3BucketSnapshotDeleteCollectionParamsWithTimeout creates a new S3BucketSnapshotDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewS3BucketSnapshotDeleteCollectionParamsWithTimeout(timeout time.Duration) *S3BucketSnapshotDeleteCollectionParams {
	return &S3BucketSnapshotDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewS3BucketSnapshotDeleteCollectionParamsWithContext creates a new S3BucketSnapshotDeleteCollectionParams object
// with the ability to set a context for a request.
func NewS3BucketSnapshotDeleteCollectionParamsWithContext(ctx context.Context) *S3BucketSnapshotDeleteCollectionParams {
	return &S3BucketSnapshotDeleteCollectionParams{
		Context: ctx,
	}
}

// NewS3BucketSnapshotDeleteCollectionParamsWithHTTPClient creates a new S3BucketSnapshotDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewS3BucketSnapshotDeleteCollectionParamsWithHTTPClient(client *http.Client) *S3BucketSnapshotDeleteCollectionParams {
	return &S3BucketSnapshotDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
S3BucketSnapshotDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the s3 bucket snapshot delete collection operation.

	Typically these are written to a http.Request.
*/
type S3BucketSnapshotDeleteCollectionParams struct {

	/* BucketUUID.

	   Filter by bucket_uuid
	*/
	BucketUUID *string

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* CreateTime.

	   Filter by create_time
	*/
	CreateTime *string

	/* Info.

	   Info specification
	*/
	Info S3BucketSnapshotDeleteCollectionBody

	/* Name.

	   Filter by name
	*/
	Name *string

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

	/* S3BucketUUID.

	   The unique identifier of the bucket.
	*/
	S3BucketUUID string

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   UUID of the SVM to which this object belongs.
	*/
	SvmUUID string

	/* UUID.

	   Filter by uuid
	*/
	UUID *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the s3 bucket snapshot delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketSnapshotDeleteCollectionParams) WithDefaults() *S3BucketSnapshotDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the s3 bucket snapshot delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *S3BucketSnapshotDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := S3BucketSnapshotDeleteCollectionParams{
		ContinueOnFailure: &continueOnFailureDefault,
		ReturnRecords:     &returnRecordsDefault,
		ReturnTimeout:     &returnTimeoutDefault,
		SerialRecords:     &serialRecordsDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithTimeout(timeout time.Duration) *S3BucketSnapshotDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithContext(ctx context.Context) *S3BucketSnapshotDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithHTTPClient(client *http.Client) *S3BucketSnapshotDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBucketUUID adds the bucketUUID to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithBucketUUID(bucketUUID *string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetBucketUUID(bucketUUID)
	return o
}

// SetBucketUUID adds the bucketUuid to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetBucketUUID(bucketUUID *string) {
	o.BucketUUID = bucketUUID
}

// WithContinueOnFailure adds the continueOnFailure to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *S3BucketSnapshotDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithCreateTime adds the createTime to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithCreateTime(createTime *string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetCreateTime(createTime)
	return o
}

// SetCreateTime adds the createTime to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetCreateTime(createTime *string) {
	o.CreateTime = createTime
}

// WithInfo adds the info to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithInfo(info S3BucketSnapshotDeleteCollectionBody) *S3BucketSnapshotDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetInfo(info S3BucketSnapshotDeleteCollectionBody) {
	o.Info = info
}

// WithName adds the name to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithName(name *string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *S3BucketSnapshotDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *S3BucketSnapshotDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithS3BucketUUID adds the s3BucketUUID to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithS3BucketUUID(s3BucketUUID string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetS3BucketUUID(s3BucketUUID)
	return o
}

// SetS3BucketUUID adds the s3BucketUuid to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetS3BucketUUID(s3BucketUUID string) {
	o.S3BucketUUID = s3BucketUUID
}

// WithSerialRecords adds the serialRecords to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *S3BucketSnapshotDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithSvmName adds the svmName to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithSvmName(svmName *string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithSvmUUID(svmUUID string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetSvmUUID(svmUUID string) {
	o.SvmUUID = svmUUID
}

// WithUUID adds the uuid to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) WithUUID(uuid *string) *S3BucketSnapshotDeleteCollectionParams {
	o.SetUUID(uuid)
	return o
}

// SetUUID adds the uuid to the s3 bucket snapshot delete collection params
func (o *S3BucketSnapshotDeleteCollectionParams) SetUUID(uuid *string) {
	o.UUID = uuid
}

// WriteToRequest writes these params to a swagger request
func (o *S3BucketSnapshotDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.BucketUUID != nil {

		// query param bucket_uuid
		var qrBucketUUID string

		if o.BucketUUID != nil {
			qrBucketUUID = *o.BucketUUID
		}
		qBucketUUID := qrBucketUUID
		if qBucketUUID != "" {

			if err := r.SetQueryParam("bucket_uuid", qBucketUUID); err != nil {
				return err
			}
		}
	}

	if o.ContinueOnFailure != nil {

		// query param continue_on_failure
		var qrContinueOnFailure bool

		if o.ContinueOnFailure != nil {
			qrContinueOnFailure = *o.ContinueOnFailure
		}
		qContinueOnFailure := swag.FormatBool(qrContinueOnFailure)
		if qContinueOnFailure != "" {

			if err := r.SetQueryParam("continue_on_failure", qContinueOnFailure); err != nil {
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
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
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

	// path param s3_bucket.uuid
	if err := r.SetPathParam("s3_bucket.uuid", o.S3BucketUUID); err != nil {
		return err
	}

	if o.SerialRecords != nil {

		// query param serial_records
		var qrSerialRecords bool

		if o.SerialRecords != nil {
			qrSerialRecords = *o.SerialRecords
		}
		qSerialRecords := swag.FormatBool(qrSerialRecords)
		if qSerialRecords != "" {

			if err := r.SetQueryParam("serial_records", qSerialRecords); err != nil {
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

	// path param svm.uuid
	if err := r.SetPathParam("svm.uuid", o.SvmUUID); err != nil {
		return err
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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
