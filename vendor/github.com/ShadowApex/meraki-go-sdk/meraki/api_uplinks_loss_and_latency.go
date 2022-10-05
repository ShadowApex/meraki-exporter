/*
Meraki Dashboard API

The Cisco Meraki Dashboard API is a modern REST API based on the OpenAPI specification.  > Date: 06 July, 2022 > > [Recent Updates](https://meraki.io/whats-new/)  ---  [API Documentation](https://meraki.io/api)  [Community Support](https://meraki.io/community)  [Meraki Homepage](https://www.meraki.com) 

API version: 1.23.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package meraki

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// UplinksLossAndLatencyApiService UplinksLossAndLatencyApi service
type UplinksLossAndLatencyApiService service

type UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest struct {
	ctx context.Context
	ApiService *UplinksLossAndLatencyApiService
	organizationId string
	t0 *string
	t1 *string
	timespan *float32
	uplink *string
	ip *string
}

// The beginning of the timespan for the data. The maximum lookback period is 365 days from today.
func (r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) T0(t0 string) UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest {
	r.t0 = &t0
	return r
}

// The end of the timespan for the data. t1 can be a maximum of 5 minutes after t0. The latest possible time that t1 can be is 2 minutes into the past.
func (r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) T1(t1 string) UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest {
	r.t1 = &t1
	return r
}

// The timespan for which the information will be fetched. If specifying timespan, do not specify parameters t0 and t1. The value must be in seconds and be less than or equal to 5 minutes. The default is 5 minutes.
func (r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) Timespan(timespan float32) UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest {
	r.timespan = &timespan
	return r
}

// Optional filter for a specific WAN uplink. Valid uplinks are wan1, wan2, cellular. Default will return all uplinks.
func (r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) Uplink(uplink string) UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest {
	r.uplink = &uplink
	return r
}

// Optional filter for a specific destination IP. Default will return all destination IPs.
func (r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) Ip(ip string) UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest {
	r.ip = &ip
	return r
}

func (r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) Execute() ([]GetOrganizationDevicesUplinksLossAndLatency200ResponseInner, *http.Response, error) {
	return r.ApiService.GetOrganizationDevicesUplinksLossAndLatencyExecute(r)
}

/*
GetOrganizationDevicesUplinksLossAndLatency Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago

Return the uplink loss and latency for every MX in the organization from at latest 2 minutes ago

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param organizationId
 @return UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest
*/
func (a *UplinksLossAndLatencyApiService) GetOrganizationDevicesUplinksLossAndLatency(ctx context.Context, organizationId string) UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest {
	return UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest{
		ApiService: a,
		ctx: ctx,
		organizationId: organizationId,
	}
}

// Execute executes the request
//  @return []GetOrganizationDevicesUplinksLossAndLatency200ResponseInner
func (a *UplinksLossAndLatencyApiService) GetOrganizationDevicesUplinksLossAndLatencyExecute(r UplinksLossAndLatencyApiGetOrganizationDevicesUplinksLossAndLatencyRequest) ([]GetOrganizationDevicesUplinksLossAndLatency200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetOrganizationDevicesUplinksLossAndLatency200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UplinksLossAndLatencyApiService.GetOrganizationDevicesUplinksLossAndLatency")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/organizations/{organizationId}/devices/uplinksLossAndLatency"
	localVarPath = strings.Replace(localVarPath, "{"+"organizationId"+"}", url.PathEscape(parameterToString(r.organizationId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.t0 != nil {
		localVarQueryParams.Add("t0", parameterToString(*r.t0, ""))
	}
	if r.t1 != nil {
		localVarQueryParams.Add("t1", parameterToString(*r.t1, ""))
	}
	if r.timespan != nil {
		localVarQueryParams.Add("timespan", parameterToString(*r.timespan, ""))
	}
	if r.uplink != nil {
		localVarQueryParams.Add("uplink", parameterToString(*r.uplink, ""))
	}
	if r.ip != nil {
		localVarQueryParams.Add("ip", parameterToString(*r.ip, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["meraki_api_key"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-Cisco-Meraki-API-Key"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
