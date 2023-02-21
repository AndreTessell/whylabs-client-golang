/*
WhyLabs Songbird

WhyLabs API that enables end-to-end AI observability

API version: 0.1
Contact: support@whylabs.ai
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"reflect"
	"os"
)


// LogApiService LogApi service
type LogApiService service

type ApiLogRequest struct {
	ctx context.Context
	ApiService *LogApiService
	orgId string
	modelId *string
	datasetTimestamp *int64
	segmentTags *[]SegmentTag
	segmentTagsJson *string
	file *os.File
}

// The unique model ID in your company.
func (r ApiLogRequest) ModelId(modelId string) ApiLogRequest {
	r.modelId = &modelId
	return r
}

// The dataset timestamp associated with the entry. Not required. However, this will  override the whylogs dataset timestamp if specified
func (r ApiLogRequest) DatasetTimestamp(datasetTimestamp int64) ApiLogRequest {
	r.datasetTimestamp = &datasetTimestamp
	return r
}

// The segment associated with the log entry. Not required if segment tags are specified in whylogs
func (r ApiLogRequest) SegmentTags(segmentTags []SegmentTag) ApiLogRequest {
	r.segmentTags = &segmentTags
	return r
}

func (r ApiLogRequest) SegmentTagsJson(segmentTagsJson string) ApiLogRequest {
	r.segmentTagsJson = &segmentTagsJson
	return r
}

// The Dataset Profile log entry
func (r ApiLogRequest) File(file *os.File) ApiLogRequest {
	r.file = file
	return r
}

func (r ApiLogRequest) Execute() (*LogResponse, *http.Response, error) {
	return r.ApiService.LogExecute(r)
}

/*
Log Log a dataset profile entry to the backend

This method returns a [LogResponse] object if it succeeds

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId Your company's unique organization ID
 @return ApiLogRequest
*/
func (a *LogApiService) Log(ctx context.Context, orgId string) ApiLogRequest {
	return ApiLogRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
	}
}

// Execute executes the request
//  @return LogResponse
func (a *LogApiService) LogExecute(r ApiLogRequest) (*LogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogApiService.Log")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/organizations/{org_id}/log"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.modelId == nil {
		return localVarReturnValue, nil, reportError("modelId is required and must be specified")
	}

	parameterAddToQuery(localVarQueryParams, "model_id", r.modelId, "")
	if r.datasetTimestamp != nil {
		parameterAddToQuery(localVarQueryParams, "dataset_timestamp", r.datasetTimestamp, "")
	}
	if r.segmentTags != nil {
		t := *r.segmentTags
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToQuery(localVarQueryParams, "segment_tags", s.Index(i), "multi")
			}
		} else {
			parameterAddToQuery(localVarQueryParams, "segment_tags", t, "multi")
		}
	}
	if r.segmentTagsJson != nil {
		parameterAddToQuery(localVarQueryParams, "segment_tags_json", r.segmentTagsJson, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"multipart/form-data"}

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
	var fileLocalVarFormFileName string
	var fileLocalVarFileName     string
	var fileLocalVarFileBytes    []byte

	fileLocalVarFormFileName = "file"


	fileLocalVarFile := r.file

	if fileLocalVarFile != nil {
		fbs, _ := ioutil.ReadAll(fileLocalVarFile)

		fileLocalVarFileBytes = fbs
		fileLocalVarFileName = fileLocalVarFile.Name()
		fileLocalVarFile.Close()
		formFiles = append(formFiles, formFile{fileBytes: fileLocalVarFileBytes, fileName: fileLocalVarFileName, formFileName: fileLocalVarFormFileName})
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
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
			var v LogResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiLogAsyncRequest struct {
	ctx context.Context
	ApiService *LogApiService
	orgId string
	datasetId string
	logAsyncRequest *LogAsyncRequest
}

func (r ApiLogAsyncRequest) LogAsyncRequest(logAsyncRequest LogAsyncRequest) ApiLogAsyncRequest {
	r.logAsyncRequest = &logAsyncRequest
	return r
}

func (r ApiLogAsyncRequest) Execute() (*AsyncLogResponse, *http.Response, error) {
	return r.ApiService.LogAsyncExecute(r)
}

/*
LogAsync Like /log, except this api doesn't take the actual profile content. It returns an upload link that can be used to upload the profile to.

Like /log, except this api doesn't take the actual profile content. It returns an upload link that can be used to upload the profile to.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @param datasetId
 @return ApiLogAsyncRequest
*/
func (a *LogApiService) LogAsync(ctx context.Context, orgId string, datasetId string) ApiLogAsyncRequest {
	return ApiLogAsyncRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		datasetId: datasetId,
	}
}

// Execute executes the request
//  @return AsyncLogResponse
func (a *LogApiService) LogAsyncExecute(r ApiLogAsyncRequest) (*AsyncLogResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *AsyncLogResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogApiService.LogAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/organizations/{org_id}/log/async/{dataset_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"dataset_id"+"}", url.PathEscape(parameterValueToString(r.datasetId, "datasetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.logAsyncRequest == nil {
		return localVarReturnValue, nil, reportError("logAsyncRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.logAsyncRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
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
			var v AsyncLogResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiLogReferenceRequest struct {
	ctx context.Context
	ApiService *LogApiService
	orgId string
	modelId string
	logReferenceRequest *LogReferenceRequest
}

func (r ApiLogReferenceRequest) LogReferenceRequest(logReferenceRequest LogReferenceRequest) ApiLogReferenceRequest {
	r.logReferenceRequest = &logReferenceRequest
	return r
}

func (r ApiLogReferenceRequest) Execute() (*LogReferenceResponse, *http.Response, error) {
	return r.ApiService.LogReferenceExecute(r)
}

/*
LogReference Returns a presigned URL for uploading the reference profile to.

Reference profiles can be used for.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param orgId
 @param modelId
 @return ApiLogReferenceRequest
*/
func (a *LogApiService) LogReference(ctx context.Context, orgId string, modelId string) ApiLogReferenceRequest {
	return ApiLogReferenceRequest{
		ApiService: a,
		ctx: ctx,
		orgId: orgId,
		modelId: modelId,
	}
}

// Execute executes the request
//  @return LogReferenceResponse
func (a *LogApiService) LogReferenceExecute(r ApiLogReferenceRequest) (*LogReferenceResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LogReferenceResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LogApiService.LogReference")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v0/organizations/{org_id}/log/reference/{model_id}"
	localVarPath = strings.Replace(localVarPath, "{"+"org_id"+"}", url.PathEscape(parameterValueToString(r.orgId, "orgId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"model_id"+"}", url.PathEscape(parameterValueToString(r.modelId, "modelId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.logReferenceRequest == nil {
		return localVarReturnValue, nil, reportError("logReferenceRequest is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

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
	// body params
	localVarPostBody = r.logReferenceRequest
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["ApiKeyAuth"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["X-API-Key"] = key
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
			var v LogReferenceResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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