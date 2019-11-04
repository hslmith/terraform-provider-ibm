// Code generated by go-swagger; DO NOT EDIT.

package network

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewPostFloatingIpsParams creates a new PostFloatingIpsParams object
// with the default values initialized.
func NewPostFloatingIpsParams() *PostFloatingIpsParams {
	var ()
	return &PostFloatingIpsParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPostFloatingIpsParamsWithTimeout creates a new PostFloatingIpsParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPostFloatingIpsParamsWithTimeout(timeout time.Duration) *PostFloatingIpsParams {
	var ()
	return &PostFloatingIpsParams{

		timeout: timeout,
	}
}

// NewPostFloatingIpsParamsWithContext creates a new PostFloatingIpsParams object
// with the default values initialized, and the ability to set a context for a request
func NewPostFloatingIpsParamsWithContext(ctx context.Context) *PostFloatingIpsParams {
	var ()
	return &PostFloatingIpsParams{

		Context: ctx,
	}
}

// NewPostFloatingIpsParamsWithHTTPClient creates a new PostFloatingIpsParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPostFloatingIpsParamsWithHTTPClient(client *http.Client) *PostFloatingIpsParams {
	var ()
	return &PostFloatingIpsParams{
		HTTPClient: client,
	}
}

/*PostFloatingIpsParams contains all the parameters to send to the API endpoint
for the post floating ips operation typically these are written to a http.Request
*/
type PostFloatingIpsParams struct {

	/*Body*/
	Body PostFloatingIpsBody
	/*Generation
	  The infrastructure generation for the request.

	*/
	Generation int64
	/*Version
	  Requests the version of the API as of a date in the format `YYYY-MM-DD`. Any date up to the current date may be provided. Specify the current date to request the latest version.

	*/
	Version string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the post floating ips params
func (o *PostFloatingIpsParams) WithTimeout(timeout time.Duration) *PostFloatingIpsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the post floating ips params
func (o *PostFloatingIpsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the post floating ips params
func (o *PostFloatingIpsParams) WithContext(ctx context.Context) *PostFloatingIpsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the post floating ips params
func (o *PostFloatingIpsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the post floating ips params
func (o *PostFloatingIpsParams) WithHTTPClient(client *http.Client) *PostFloatingIpsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the post floating ips params
func (o *PostFloatingIpsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the post floating ips params
func (o *PostFloatingIpsParams) WithBody(body PostFloatingIpsBody) *PostFloatingIpsParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the post floating ips params
func (o *PostFloatingIpsParams) SetBody(body PostFloatingIpsBody) {
	o.Body = body
}

// WithGeneration adds the generation to the post floating ips params
func (o *PostFloatingIpsParams) WithGeneration(generation int64) *PostFloatingIpsParams {
	o.SetGeneration(generation)
	return o
}

// SetGeneration adds the generation to the post floating ips params
func (o *PostFloatingIpsParams) SetGeneration(generation int64) {
	o.Generation = generation
}

// WithVersion adds the version to the post floating ips params
func (o *PostFloatingIpsParams) WithVersion(version string) *PostFloatingIpsParams {
	o.SetVersion(version)
	return o
}

// SetVersion adds the version to the post floating ips params
func (o *PostFloatingIpsParams) SetVersion(version string) {
	o.Version = version
}

// WriteToRequest writes these params to a swagger request
func (o *PostFloatingIpsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	// query param generation
	qrGeneration := o.Generation
	qGeneration := swag.FormatInt64(qrGeneration)
	if qGeneration != "" {
		if err := r.SetQueryParam("generation", qGeneration); err != nil {
			return err
		}
	}

	// query param version
	qrVersion := o.Version
	qVersion := qrVersion
	if qVersion != "" {
		if err := r.SetQueryParam("version", qVersion); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}