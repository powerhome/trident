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

// NewTopMetricsUserCollectionGetParams creates a new TopMetricsUserCollectionGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewTopMetricsUserCollectionGetParams() *TopMetricsUserCollectionGetParams {
	return &TopMetricsUserCollectionGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewTopMetricsUserCollectionGetParamsWithTimeout creates a new TopMetricsUserCollectionGetParams object
// with the ability to set a timeout on a request.
func NewTopMetricsUserCollectionGetParamsWithTimeout(timeout time.Duration) *TopMetricsUserCollectionGetParams {
	return &TopMetricsUserCollectionGetParams{
		timeout: timeout,
	}
}

// NewTopMetricsUserCollectionGetParamsWithContext creates a new TopMetricsUserCollectionGetParams object
// with the ability to set a context for a request.
func NewTopMetricsUserCollectionGetParamsWithContext(ctx context.Context) *TopMetricsUserCollectionGetParams {
	return &TopMetricsUserCollectionGetParams{
		Context: ctx,
	}
}

// NewTopMetricsUserCollectionGetParamsWithHTTPClient creates a new TopMetricsUserCollectionGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewTopMetricsUserCollectionGetParamsWithHTTPClient(client *http.Client) *TopMetricsUserCollectionGetParams {
	return &TopMetricsUserCollectionGetParams{
		HTTPClient: client,
	}
}

/*
TopMetricsUserCollectionGetParams contains all the parameters to send to the API endpoint

	for the top metrics user collection get operation.

	Typically these are written to a http.Request.
*/
type TopMetricsUserCollectionGetParams struct {

	/* Fields.

	   Specify the fields to return.
	*/
	Fields []string

	/* IopsErrorLowerBound.

	   Filter by iops.error.lower_bound
	*/
	IopsErrorLowerBound *int64

	/* IopsErrorUpperBound.

	   Filter by iops.error.upper_bound
	*/
	IopsErrorUpperBound *int64

	/* IopsRead.

	   Filter by iops.read
	*/
	IopsRead *int64

	/* IopsWrite.

	   Filter by iops.write
	*/
	IopsWrite *int64

	/* MaxRecords.

	   Limit the number of records returned.
	*/
	MaxRecords *int64

	/* OrderBy.

	   Order results by specified fields and optional [asc|desc] direction. Default direction is 'asc' for ascending.
	*/
	OrderBy []string

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

	/* SvmName.

	   Filter by svm.name
	*/
	SvmName *string

	/* SvmUUID.

	   Filter by svm.uuid
	*/
	SvmUUID *string

	/* ThroughputErrorLowerBound.

	   Filter by throughput.error.lower_bound
	*/
	ThroughputErrorLowerBound *int64

	/* ThroughputErrorUpperBound.

	   Filter by throughput.error.upper_bound
	*/
	ThroughputErrorUpperBound *int64

	/* ThroughputRead.

	   Filter by throughput.read
	*/
	ThroughputRead *int64

	/* ThroughputWrite.

	   Filter by throughput.write
	*/
	ThroughputWrite *int64

	/* TopMetric.

	   I/O activity type

	   Default: "iops.read"
	*/
	TopMetric *string

	/* UserID.

	   Filter by user_id
	*/
	UserID *string

	/* UserName.

	   Filter by user_name
	*/
	UserName *string

	/* VolumeName.

	   Filter by volume.name
	*/
	VolumeName *string

	/* VolumeUUID.

	   Volume UUID
	*/
	VolumeUUID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the top metrics user collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsUserCollectionGetParams) WithDefaults() *TopMetricsUserCollectionGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the top metrics user collection get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *TopMetricsUserCollectionGetParams) SetDefaults() {
	var (
		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		topMetricDefault = string("iops.read")
	)

	val := TopMetricsUserCollectionGetParams{
		ReturnRecords: &returnRecordsDefault,
		ReturnTimeout: &returnTimeoutDefault,
		TopMetric:     &topMetricDefault,
	}

	val.timeout = o.timeout
	val.Context = o.Context
	val.HTTPClient = o.HTTPClient
	*o = val
}

// WithTimeout adds the timeout to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithTimeout(timeout time.Duration) *TopMetricsUserCollectionGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithContext(ctx context.Context) *TopMetricsUserCollectionGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithHTTPClient(client *http.Client) *TopMetricsUserCollectionGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithFields(fields []string) *TopMetricsUserCollectionGetParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetFields(fields []string) {
	o.Fields = fields
}

// WithIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithIopsErrorLowerBound(iopsErrorLowerBound *int64) *TopMetricsUserCollectionGetParams {
	o.SetIopsErrorLowerBound(iopsErrorLowerBound)
	return o
}

// SetIopsErrorLowerBound adds the iopsErrorLowerBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetIopsErrorLowerBound(iopsErrorLowerBound *int64) {
	o.IopsErrorLowerBound = iopsErrorLowerBound
}

// WithIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithIopsErrorUpperBound(iopsErrorUpperBound *int64) *TopMetricsUserCollectionGetParams {
	o.SetIopsErrorUpperBound(iopsErrorUpperBound)
	return o
}

// SetIopsErrorUpperBound adds the iopsErrorUpperBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetIopsErrorUpperBound(iopsErrorUpperBound *int64) {
	o.IopsErrorUpperBound = iopsErrorUpperBound
}

// WithIopsRead adds the iopsRead to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithIopsRead(iopsRead *int64) *TopMetricsUserCollectionGetParams {
	o.SetIopsRead(iopsRead)
	return o
}

// SetIopsRead adds the iopsRead to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetIopsRead(iopsRead *int64) {
	o.IopsRead = iopsRead
}

// WithIopsWrite adds the iopsWrite to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithIopsWrite(iopsWrite *int64) *TopMetricsUserCollectionGetParams {
	o.SetIopsWrite(iopsWrite)
	return o
}

// SetIopsWrite adds the iopsWrite to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetIopsWrite(iopsWrite *int64) {
	o.IopsWrite = iopsWrite
}

// WithMaxRecords adds the maxRecords to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithMaxRecords(maxRecords *int64) *TopMetricsUserCollectionGetParams {
	o.SetMaxRecords(maxRecords)
	return o
}

// SetMaxRecords adds the maxRecords to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetMaxRecords(maxRecords *int64) {
	o.MaxRecords = maxRecords
}

// WithOrderBy adds the orderBy to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithOrderBy(orderBy []string) *TopMetricsUserCollectionGetParams {
	o.SetOrderBy(orderBy)
	return o
}

// SetOrderBy adds the orderBy to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetOrderBy(orderBy []string) {
	o.OrderBy = orderBy
}

// WithReturnRecords adds the returnRecords to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithReturnRecords(returnRecords *bool) *TopMetricsUserCollectionGetParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithReturnTimeout(returnTimeout *int64) *TopMetricsUserCollectionGetParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSvmName adds the svmName to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithSvmName(svmName *string) *TopMetricsUserCollectionGetParams {
	o.SetSvmName(svmName)
	return o
}

// SetSvmName adds the svmName to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetSvmName(svmName *string) {
	o.SvmName = svmName
}

// WithSvmUUID adds the svmUUID to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithSvmUUID(svmUUID *string) *TopMetricsUserCollectionGetParams {
	o.SetSvmUUID(svmUUID)
	return o
}

// SetSvmUUID adds the svmUuid to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetSvmUUID(svmUUID *string) {
	o.SvmUUID = svmUUID
}

// WithThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithThroughputErrorLowerBound(throughputErrorLowerBound *int64) *TopMetricsUserCollectionGetParams {
	o.SetThroughputErrorLowerBound(throughputErrorLowerBound)
	return o
}

// SetThroughputErrorLowerBound adds the throughputErrorLowerBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetThroughputErrorLowerBound(throughputErrorLowerBound *int64) {
	o.ThroughputErrorLowerBound = throughputErrorLowerBound
}

// WithThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithThroughputErrorUpperBound(throughputErrorUpperBound *int64) *TopMetricsUserCollectionGetParams {
	o.SetThroughputErrorUpperBound(throughputErrorUpperBound)
	return o
}

// SetThroughputErrorUpperBound adds the throughputErrorUpperBound to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetThroughputErrorUpperBound(throughputErrorUpperBound *int64) {
	o.ThroughputErrorUpperBound = throughputErrorUpperBound
}

// WithThroughputRead adds the throughputRead to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithThroughputRead(throughputRead *int64) *TopMetricsUserCollectionGetParams {
	o.SetThroughputRead(throughputRead)
	return o
}

// SetThroughputRead adds the throughputRead to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetThroughputRead(throughputRead *int64) {
	o.ThroughputRead = throughputRead
}

// WithThroughputWrite adds the throughputWrite to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithThroughputWrite(throughputWrite *int64) *TopMetricsUserCollectionGetParams {
	o.SetThroughputWrite(throughputWrite)
	return o
}

// SetThroughputWrite adds the throughputWrite to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetThroughputWrite(throughputWrite *int64) {
	o.ThroughputWrite = throughputWrite
}

// WithTopMetric adds the topMetric to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithTopMetric(topMetric *string) *TopMetricsUserCollectionGetParams {
	o.SetTopMetric(topMetric)
	return o
}

// SetTopMetric adds the topMetric to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetTopMetric(topMetric *string) {
	o.TopMetric = topMetric
}

// WithUserID adds the userID to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithUserID(userID *string) *TopMetricsUserCollectionGetParams {
	o.SetUserID(userID)
	return o
}

// SetUserID adds the userId to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetUserID(userID *string) {
	o.UserID = userID
}

// WithUserName adds the userName to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithUserName(userName *string) *TopMetricsUserCollectionGetParams {
	o.SetUserName(userName)
	return o
}

// SetUserName adds the userName to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetUserName(userName *string) {
	o.UserName = userName
}

// WithVolumeName adds the volumeName to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithVolumeName(volumeName *string) *TopMetricsUserCollectionGetParams {
	o.SetVolumeName(volumeName)
	return o
}

// SetVolumeName adds the volumeName to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetVolumeName(volumeName *string) {
	o.VolumeName = volumeName
}

// WithVolumeUUID adds the volumeUUID to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) WithVolumeUUID(volumeUUID string) *TopMetricsUserCollectionGetParams {
	o.SetVolumeUUID(volumeUUID)
	return o
}

// SetVolumeUUID adds the volumeUuid to the top metrics user collection get params
func (o *TopMetricsUserCollectionGetParams) SetVolumeUUID(volumeUUID string) {
	o.VolumeUUID = volumeUUID
}

// WriteToRequest writes these params to a swagger request
func (o *TopMetricsUserCollectionGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// binding items for fields
		joinedFields := o.bindParamFields(reg)

		// query array param fields
		if err := r.SetQueryParam("fields", joinedFields...); err != nil {
			return err
		}
	}

	if o.IopsErrorLowerBound != nil {

		// query param iops.error.lower_bound
		var qrIopsErrorLowerBound int64

		if o.IopsErrorLowerBound != nil {
			qrIopsErrorLowerBound = *o.IopsErrorLowerBound
		}
		qIopsErrorLowerBound := swag.FormatInt64(qrIopsErrorLowerBound)
		if qIopsErrorLowerBound != "" {

			if err := r.SetQueryParam("iops.error.lower_bound", qIopsErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.IopsErrorUpperBound != nil {

		// query param iops.error.upper_bound
		var qrIopsErrorUpperBound int64

		if o.IopsErrorUpperBound != nil {
			qrIopsErrorUpperBound = *o.IopsErrorUpperBound
		}
		qIopsErrorUpperBound := swag.FormatInt64(qrIopsErrorUpperBound)
		if qIopsErrorUpperBound != "" {

			if err := r.SetQueryParam("iops.error.upper_bound", qIopsErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.IopsRead != nil {

		// query param iops.read
		var qrIopsRead int64

		if o.IopsRead != nil {
			qrIopsRead = *o.IopsRead
		}
		qIopsRead := swag.FormatInt64(qrIopsRead)
		if qIopsRead != "" {

			if err := r.SetQueryParam("iops.read", qIopsRead); err != nil {
				return err
			}
		}
	}

	if o.IopsWrite != nil {

		// query param iops.write
		var qrIopsWrite int64

		if o.IopsWrite != nil {
			qrIopsWrite = *o.IopsWrite
		}
		qIopsWrite := swag.FormatInt64(qrIopsWrite)
		if qIopsWrite != "" {

			if err := r.SetQueryParam("iops.write", qIopsWrite); err != nil {
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

	if o.OrderBy != nil {

		// binding items for order_by
		joinedOrderBy := o.bindParamOrderBy(reg)

		// query array param order_by
		if err := r.SetQueryParam("order_by", joinedOrderBy...); err != nil {
			return err
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

	if o.ThroughputErrorLowerBound != nil {

		// query param throughput.error.lower_bound
		var qrThroughputErrorLowerBound int64

		if o.ThroughputErrorLowerBound != nil {
			qrThroughputErrorLowerBound = *o.ThroughputErrorLowerBound
		}
		qThroughputErrorLowerBound := swag.FormatInt64(qrThroughputErrorLowerBound)
		if qThroughputErrorLowerBound != "" {

			if err := r.SetQueryParam("throughput.error.lower_bound", qThroughputErrorLowerBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputErrorUpperBound != nil {

		// query param throughput.error.upper_bound
		var qrThroughputErrorUpperBound int64

		if o.ThroughputErrorUpperBound != nil {
			qrThroughputErrorUpperBound = *o.ThroughputErrorUpperBound
		}
		qThroughputErrorUpperBound := swag.FormatInt64(qrThroughputErrorUpperBound)
		if qThroughputErrorUpperBound != "" {

			if err := r.SetQueryParam("throughput.error.upper_bound", qThroughputErrorUpperBound); err != nil {
				return err
			}
		}
	}

	if o.ThroughputRead != nil {

		// query param throughput.read
		var qrThroughputRead int64

		if o.ThroughputRead != nil {
			qrThroughputRead = *o.ThroughputRead
		}
		qThroughputRead := swag.FormatInt64(qrThroughputRead)
		if qThroughputRead != "" {

			if err := r.SetQueryParam("throughput.read", qThroughputRead); err != nil {
				return err
			}
		}
	}

	if o.ThroughputWrite != nil {

		// query param throughput.write
		var qrThroughputWrite int64

		if o.ThroughputWrite != nil {
			qrThroughputWrite = *o.ThroughputWrite
		}
		qThroughputWrite := swag.FormatInt64(qrThroughputWrite)
		if qThroughputWrite != "" {

			if err := r.SetQueryParam("throughput.write", qThroughputWrite); err != nil {
				return err
			}
		}
	}

	if o.TopMetric != nil {

		// query param top_metric
		var qrTopMetric string

		if o.TopMetric != nil {
			qrTopMetric = *o.TopMetric
		}
		qTopMetric := qrTopMetric
		if qTopMetric != "" {

			if err := r.SetQueryParam("top_metric", qTopMetric); err != nil {
				return err
			}
		}
	}

	if o.UserID != nil {

		// query param user_id
		var qrUserID string

		if o.UserID != nil {
			qrUserID = *o.UserID
		}
		qUserID := qrUserID
		if qUserID != "" {

			if err := r.SetQueryParam("user_id", qUserID); err != nil {
				return err
			}
		}
	}

	if o.UserName != nil {

		// query param user_name
		var qrUserName string

		if o.UserName != nil {
			qrUserName = *o.UserName
		}
		qUserName := qrUserName
		if qUserName != "" {

			if err := r.SetQueryParam("user_name", qUserName); err != nil {
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

// bindParamTopMetricsUserCollectionGet binds the parameter fields
func (o *TopMetricsUserCollectionGetParams) bindParamFields(formats strfmt.Registry) []string {
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

// bindParamTopMetricsUserCollectionGet binds the parameter order_by
func (o *TopMetricsUserCollectionGetParams) bindParamOrderBy(formats strfmt.Registry) []string {
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
