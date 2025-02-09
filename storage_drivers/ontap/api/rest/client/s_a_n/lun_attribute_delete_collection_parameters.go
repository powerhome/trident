// Code generated by go-swagger; DO NOT EDIT.

package s_a_n

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

// NewLunAttributeDeleteCollectionParams creates a new LunAttributeDeleteCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewLunAttributeDeleteCollectionParams() *LunAttributeDeleteCollectionParams {
	return &LunAttributeDeleteCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewLunAttributeDeleteCollectionParamsWithTimeout creates a new LunAttributeDeleteCollectionParams object
// with the ability to set a timeout on a request.
func NewLunAttributeDeleteCollectionParamsWithTimeout(timeout time.Duration) *LunAttributeDeleteCollectionParams {
	return &LunAttributeDeleteCollectionParams{
		timeout: timeout,
	}
}

// NewLunAttributeDeleteCollectionParamsWithContext creates a new LunAttributeDeleteCollectionParams object
// with the ability to set a context for a request.
func NewLunAttributeDeleteCollectionParamsWithContext(ctx context.Context) *LunAttributeDeleteCollectionParams {
	return &LunAttributeDeleteCollectionParams{
		Context: ctx,
	}
}

// NewLunAttributeDeleteCollectionParamsWithHTTPClient creates a new LunAttributeDeleteCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewLunAttributeDeleteCollectionParamsWithHTTPClient(client *http.Client) *LunAttributeDeleteCollectionParams {
	return &LunAttributeDeleteCollectionParams{
		HTTPClient: client,
	}
}

/*
LunAttributeDeleteCollectionParams contains all the parameters to send to the API endpoint

	for the lun attribute delete collection operation.

	Typically these are written to a http.Request.
*/
type LunAttributeDeleteCollectionParams struct {

	/* ContinueOnFailure.

	   Continue even when the operation fails on one of the records.
	*/
	ContinueOnFailure *bool

	/* Info.

	   Info specification
	*/
	Info LunAttributeDeleteCollectionBody

	/* LunUUID.

	   The unique identifier of the LUN.

	*/
	LunUUID string

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

	/* SerialRecords.

	   Perform the operation on the records synchronously.
	*/
	SerialRecords *bool

	/* Value.

	   Filter by value
	*/
	Value *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the lun attribute delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeDeleteCollectionParams) WithDefaults() *LunAttributeDeleteCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the lun attribute delete collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *LunAttributeDeleteCollectionParams) SetDefaults() {
	var (
		continueOnFailureDefault = bool(false)

		returnRecordsDefault = bool(true)

		returnTimeoutDefault = int64(15)

		serialRecordsDefault = bool(false)
	)

	val := LunAttributeDeleteCollectionParams{
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

// WithTimeout adds the timeout to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithTimeout(timeout time.Duration) *LunAttributeDeleteCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithContext(ctx context.Context) *LunAttributeDeleteCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithHTTPClient(client *http.Client) *LunAttributeDeleteCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithContinueOnFailure adds the continueOnFailure to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithContinueOnFailure(continueOnFailure *bool) *LunAttributeDeleteCollectionParams {
	o.SetContinueOnFailure(continueOnFailure)
	return o
}

// SetContinueOnFailure adds the continueOnFailure to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetContinueOnFailure(continueOnFailure *bool) {
	o.ContinueOnFailure = continueOnFailure
}

// WithInfo adds the info to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithInfo(info LunAttributeDeleteCollectionBody) *LunAttributeDeleteCollectionParams {
	o.SetInfo(info)
	return o
}

// SetInfo adds the info to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetInfo(info LunAttributeDeleteCollectionBody) {
	o.Info = info
}

// WithLunUUID adds the lunUUID to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithLunUUID(lunUUID string) *LunAttributeDeleteCollectionParams {
	o.SetLunUUID(lunUUID)
	return o
}

// SetLunUUID adds the lunUuid to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetLunUUID(lunUUID string) {
	o.LunUUID = lunUUID
}

// WithName adds the name to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithName(name *string) *LunAttributeDeleteCollectionParams {
	o.SetName(name)
	return o
}

// SetName adds the name to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetName(name *string) {
	o.Name = name
}

// WithReturnRecords adds the returnRecords to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithReturnRecords(returnRecords *bool) *LunAttributeDeleteCollectionParams {
	o.SetReturnRecords(returnRecords)
	return o
}

// SetReturnRecords adds the returnRecords to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetReturnRecords(returnRecords *bool) {
	o.ReturnRecords = returnRecords
}

// WithReturnTimeout adds the returnTimeout to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithReturnTimeout(returnTimeout *int64) *LunAttributeDeleteCollectionParams {
	o.SetReturnTimeout(returnTimeout)
	return o
}

// SetReturnTimeout adds the returnTimeout to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetReturnTimeout(returnTimeout *int64) {
	o.ReturnTimeout = returnTimeout
}

// WithSerialRecords adds the serialRecords to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithSerialRecords(serialRecords *bool) *LunAttributeDeleteCollectionParams {
	o.SetSerialRecords(serialRecords)
	return o
}

// SetSerialRecords adds the serialRecords to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetSerialRecords(serialRecords *bool) {
	o.SerialRecords = serialRecords
}

// WithValue adds the value to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) WithValue(value *string) *LunAttributeDeleteCollectionParams {
	o.SetValue(value)
	return o
}

// SetValue adds the value to the lun attribute delete collection params
func (o *LunAttributeDeleteCollectionParams) SetValue(value *string) {
	o.Value = value
}

// WriteToRequest writes these params to a swagger request
func (o *LunAttributeDeleteCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
	if err := r.SetBodyParam(o.Info); err != nil {
		return err
	}

	// path param lun.uuid
	if err := r.SetPathParam("lun.uuid", o.LunUUID); err != nil {
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

	if o.Value != nil {

		// query param value
		var qrValue string

		if o.Value != nil {
			qrValue = *o.Value
		}
		qValue := qrValue
		if qValue != "" {

			if err := r.SetQueryParam("value", qValue); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
