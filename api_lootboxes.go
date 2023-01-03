/*
MetaFab API

 Complete MetaFab API references and guides can be found at: https://trymetafab.com

API version: 1.4.1
Contact: metafabproject@gmail.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package metafab

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// LootboxesApiService LootboxesApi service
type LootboxesApiService service

type ApiCreateLootboxManagerRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	xAuthorization *string
	xPassword *string
	createLootboxManagerRequest *CreateLootboxManagerRequest
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiCreateLootboxManagerRequest) XAuthorization(xAuthorization string) ApiCreateLootboxManagerRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiCreateLootboxManagerRequest) XPassword(xPassword string) ApiCreateLootboxManagerRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiCreateLootboxManagerRequest) CreateLootboxManagerRequest(createLootboxManagerRequest CreateLootboxManagerRequest) ApiCreateLootboxManagerRequest {
	r.createLootboxManagerRequest = &createLootboxManagerRequest
	return r
}

func (r ApiCreateLootboxManagerRequest) Execute() (*CreateLootboxManager200Response, *http.Response, error) {
	return r.ApiService.CreateLootboxManagerExecute(r)
}

/*
CreateLootboxManager Create lootbox manager

Creates a new game lootbox manager and deploys a lootbox manager contract on behalf of the authenticating game's primary wallet. The deployed lootbox manager contract allows you to create lootbox behavior for existing items. For example, you can define item id(s) from one of your item collections as the requirement(s) to open a "lootbox". The required item(s) would be burned from the interacting player's wallet and the player would receive item(s) from a weighted randomized set of possible items the lootbox can contain.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateLootboxManagerRequest
*/
func (a *LootboxesApiService) CreateLootboxManager(ctx context.Context) ApiCreateLootboxManagerRequest {
	return ApiCreateLootboxManagerRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateLootboxManager200Response
func (a *LootboxesApiService) CreateLootboxManagerExecute(r ApiCreateLootboxManagerRequest) (*CreateLootboxManager200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateLootboxManager200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.CreateLootboxManager")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
	}
	if r.createLootboxManagerRequest == nil {
		return localVarReturnValue, nil, reportError("createLootboxManagerRequest is required and must be specified")
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
	localVarHeaderParams["X-Authorization"] = parameterToString(*r.xAuthorization, "")
	localVarHeaderParams["X-Password"] = parameterToString(*r.xPassword, "")
	// body params
	localVarPostBody = r.createLootboxManagerRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiGetLootboxManagerLootboxRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	lootboxManagerId string
	lootboxManagerLootboxId string
}

func (r ApiGetLootboxManagerLootboxRequest) Execute() (*LootboxManagerLootbox, *http.Response, error) {
	return r.ApiService.GetLootboxManagerLootboxExecute(r)
}

/*
GetLootboxManagerLootbox Get lootbox manager lootbox

Returns a lootbox manager lootbox object for the provided lootboxManagerLootboxId.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lootboxManagerId Any lootbox manager id within the MetaFab ecosystem.
 @param lootboxManagerLootboxId Any lootbox manager lootbox id within the MetaFab ecosystem.
 @return ApiGetLootboxManagerLootboxRequest
*/
func (a *LootboxesApiService) GetLootboxManagerLootbox(ctx context.Context, lootboxManagerId string, lootboxManagerLootboxId string) ApiGetLootboxManagerLootboxRequest {
	return ApiGetLootboxManagerLootboxRequest{
		ApiService: a,
		ctx: ctx,
		lootboxManagerId: lootboxManagerId,
		lootboxManagerLootboxId: lootboxManagerLootboxId,
	}
}

// Execute executes the request
//  @return LootboxManagerLootbox
func (a *LootboxesApiService) GetLootboxManagerLootboxExecute(r ApiGetLootboxManagerLootboxRequest) (*LootboxManagerLootbox, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LootboxManagerLootbox
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.GetLootboxManagerLootbox")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId}"
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerId"+"}", url.PathEscape(parameterToString(r.lootboxManagerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerLootboxId"+"}", url.PathEscape(parameterToString(r.lootboxManagerLootboxId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiGetLootboxManagerLootboxesRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	lootboxManagerId string
}

func (r ApiGetLootboxManagerLootboxesRequest) Execute() ([]LootboxManagerLootbox, *http.Response, error) {
	return r.ApiService.GetLootboxManagerLootboxesExecute(r)
}

/*
GetLootboxManagerLootboxes Get lootbox manager lootboxes

Returns all lootbox manager lootboxes as an array of lootbox objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lootboxManagerId Any lootbox manager id within the MetaFab ecosystem.
 @return ApiGetLootboxManagerLootboxesRequest
*/
func (a *LootboxesApiService) GetLootboxManagerLootboxes(ctx context.Context, lootboxManagerId string) ApiGetLootboxManagerLootboxesRequest {
	return ApiGetLootboxManagerLootboxesRequest{
		ApiService: a,
		ctx: ctx,
		lootboxManagerId: lootboxManagerId,
	}
}

// Execute executes the request
//  @return []LootboxManagerLootbox
func (a *LootboxesApiService) GetLootboxManagerLootboxesExecute(r ApiGetLootboxManagerLootboxesRequest) ([]LootboxManagerLootbox, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []LootboxManagerLootbox
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.GetLootboxManagerLootboxes")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers/{lootboxManagerId}/lootboxes"
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerId"+"}", url.PathEscape(parameterToString(r.lootboxManagerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiGetLootboxManagersRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	xGameKey *string
}

// The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc.
func (r ApiGetLootboxManagersRequest) XGameKey(xGameKey string) ApiGetLootboxManagersRequest {
	r.xGameKey = &xGameKey
	return r
}

func (r ApiGetLootboxManagersRequest) Execute() ([]GetLootboxManagers200ResponseInner, *http.Response, error) {
	return r.ApiService.GetLootboxManagersExecute(r)
}

/*
GetLootboxManagers Get lootbox managers

Returns an array of active lootbox managers for the game associated with the provided `X-Game-Key`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetLootboxManagersRequest
*/
func (a *LootboxesApiService) GetLootboxManagers(ctx context.Context) ApiGetLootboxManagersRequest {
	return ApiGetLootboxManagersRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetLootboxManagers200ResponseInner
func (a *LootboxesApiService) GetLootboxManagersExecute(r ApiGetLootboxManagersRequest) ([]GetLootboxManagers200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetLootboxManagers200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.GetLootboxManagers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xGameKey == nil {
		return localVarReturnValue, nil, reportError("xGameKey is required and must be specified")
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
	localVarHeaderParams["X-Game-Key"] = parameterToString(*r.xGameKey, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiOpenLootboxManagerLootboxRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	lootboxManagerId string
	lootboxManagerLootboxId string
	xAuthorization *string
	xPassword *string
}

// The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player.
func (r ApiOpenLootboxManagerLootboxRequest) XAuthorization(xAuthorization string) ApiOpenLootboxManagerLootboxRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
func (r ApiOpenLootboxManagerLootboxRequest) XPassword(xPassword string) ApiOpenLootboxManagerLootboxRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiOpenLootboxManagerLootboxRequest) Execute() ([]TransactionModel, *http.Response, error) {
	return r.ApiService.OpenLootboxManagerLootboxExecute(r)
}

/*
OpenLootboxManagerLootbox Open lootbox manager lootbox

Opens a lootbox manager lootbox. The required input item(s) are burned from the wallet or player wallet opening the lootbox. The given output item(s) are given to the wallet or player wallet opening the lootbox.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lootboxManagerId Any lootbox manager id within the MetaFab ecosystem.
 @param lootboxManagerLootboxId Any lootbox manager lootbox id within the MetaFab ecosystem.
 @return ApiOpenLootboxManagerLootboxRequest
*/
func (a *LootboxesApiService) OpenLootboxManagerLootbox(ctx context.Context, lootboxManagerId string, lootboxManagerLootboxId string) ApiOpenLootboxManagerLootboxRequest {
	return ApiOpenLootboxManagerLootboxRequest{
		ApiService: a,
		ctx: ctx,
		lootboxManagerId: lootboxManagerId,
		lootboxManagerLootboxId: lootboxManagerLootboxId,
	}
}

// Execute executes the request
//  @return []TransactionModel
func (a *LootboxesApiService) OpenLootboxManagerLootboxExecute(r ApiOpenLootboxManagerLootboxRequest) ([]TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.OpenLootboxManagerLootbox")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId}/opens"
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerId"+"}", url.PathEscape(parameterToString(r.lootboxManagerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerLootboxId"+"}", url.PathEscape(parameterToString(r.lootboxManagerLootboxId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
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
	localVarHeaderParams["X-Authorization"] = parameterToString(*r.xAuthorization, "")
	localVarHeaderParams["X-Password"] = parameterToString(*r.xPassword, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiRemoveLootboxManagerLootboxRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	lootboxManagerId string
	lootboxManagerLootboxId string
	xAuthorization *string
	xPassword *string
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiRemoveLootboxManagerLootboxRequest) XAuthorization(xAuthorization string) ApiRemoveLootboxManagerLootboxRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiRemoveLootboxManagerLootboxRequest) XPassword(xPassword string) ApiRemoveLootboxManagerLootboxRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiRemoveLootboxManagerLootboxRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.RemoveLootboxManagerLootboxExecute(r)
}

/*
RemoveLootboxManagerLootbox Remove lootbox manager lootbox

Removes the provided lootbox by lootboxId from the provided lootbox manager. Removed lootboxes can no longer be used.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lootboxManagerId Any lootbox manager id within the MetaFab ecosystem.
 @param lootboxManagerLootboxId Any lootbox manager lootbox id within the MetaFab ecosystem.
 @return ApiRemoveLootboxManagerLootboxRequest
*/
func (a *LootboxesApiService) RemoveLootboxManagerLootbox(ctx context.Context, lootboxManagerId string, lootboxManagerLootboxId string) ApiRemoveLootboxManagerLootboxRequest {
	return ApiRemoveLootboxManagerLootboxRequest{
		ApiService: a,
		ctx: ctx,
		lootboxManagerId: lootboxManagerId,
		lootboxManagerLootboxId: lootboxManagerLootboxId,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *LootboxesApiService) RemoveLootboxManagerLootboxExecute(r ApiRemoveLootboxManagerLootboxRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.RemoveLootboxManagerLootbox")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers/{lootboxManagerId}/lootboxes/{lootboxManagerLootboxId}"
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerId"+"}", url.PathEscape(parameterToString(r.lootboxManagerId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerLootboxId"+"}", url.PathEscape(parameterToString(r.lootboxManagerLootboxId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
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
	localVarHeaderParams["X-Authorization"] = parameterToString(*r.xAuthorization, "")
	localVarHeaderParams["X-Password"] = parameterToString(*r.xPassword, "")
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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

type ApiSetLootboxManagerLootboxRequest struct {
	ctx context.Context
	ApiService *LootboxesApiService
	lootboxManagerId string
	xAuthorization *string
	xPassword *string
	setLootboxManagerLootboxRequest *SetLootboxManagerLootboxRequest
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiSetLootboxManagerLootboxRequest) XAuthorization(xAuthorization string) ApiSetLootboxManagerLootboxRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiSetLootboxManagerLootboxRequest) XPassword(xPassword string) ApiSetLootboxManagerLootboxRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiSetLootboxManagerLootboxRequest) SetLootboxManagerLootboxRequest(setLootboxManagerLootboxRequest SetLootboxManagerLootboxRequest) ApiSetLootboxManagerLootboxRequest {
	r.setLootboxManagerLootboxRequest = &setLootboxManagerLootboxRequest
	return r
}

func (r ApiSetLootboxManagerLootboxRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.SetLootboxManagerLootboxExecute(r)
}

/*
SetLootboxManagerLootbox Set lootbox manager lootbox

Sets a new lootbox manager lootbox or updates an existing one for the provided id. Lootboxes allow item(s) to be burned to receive a random set of possible item(s) based on probability weight.

Lootboxes can require any number of unique types of items and quantities to open a created lootbox type within the system. A common pattern with lootboxes is to create a lootbox item type within an item collection, and require it as the input item type.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param lootboxManagerId Any lootbox manager id within the MetaFab ecosystem.
 @return ApiSetLootboxManagerLootboxRequest
*/
func (a *LootboxesApiService) SetLootboxManagerLootbox(ctx context.Context, lootboxManagerId string) ApiSetLootboxManagerLootboxRequest {
	return ApiSetLootboxManagerLootboxRequest{
		ApiService: a,
		ctx: ctx,
		lootboxManagerId: lootboxManagerId,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *LootboxesApiService) SetLootboxManagerLootboxExecute(r ApiSetLootboxManagerLootboxRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LootboxesApiService.SetLootboxManagerLootbox")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/lootboxManagers/{lootboxManagerId}/lootboxes"
	localVarPath = strings.Replace(localVarPath, "{"+"lootboxManagerId"+"}", url.PathEscape(parameterToString(r.lootboxManagerId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
	}
	if r.setLootboxManagerLootboxRequest == nil {
		return localVarReturnValue, nil, reportError("setLootboxManagerLootboxRequest is required and must be specified")
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
	localVarHeaderParams["X-Authorization"] = parameterToString(*r.xAuthorization, "")
	localVarHeaderParams["X-Password"] = parameterToString(*r.xPassword, "")
	// body params
	localVarPostBody = r.setLootboxManagerLootboxRequest
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
		if localVarHTTPResponse.StatusCode == 400 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v string
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
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
