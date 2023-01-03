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


// ShopsApiService ShopsApi service
type ShopsApiService service

type ApiCreateShopRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	xAuthorization *string
	xPassword *string
	createShopRequest *CreateShopRequest
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiCreateShopRequest) XAuthorization(xAuthorization string) ApiCreateShopRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiCreateShopRequest) XPassword(xPassword string) ApiCreateShopRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiCreateShopRequest) CreateShopRequest(createShopRequest CreateShopRequest) ApiCreateShopRequest {
	r.createShopRequest = &createShopRequest
	return r
}

func (r ApiCreateShopRequest) Execute() (*CreateShop200Response, *http.Response, error) {
	return r.ApiService.CreateShopExecute(r)
}

/*
CreateShop Create shop

Creates a new game shop and deploys a shop contract on behalf of the authenticating game's primary wallet. The deployed shop contract allows you to create fixed price rates for players to buy specific items from any item collection or ERC1155 contract. Additionally, a shop allows you to create shop offers for some set of item(s) to another set of item(s) or any mix of currency. Shops completely support gasless player transactions.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateShopRequest
*/
func (a *ShopsApiService) CreateShop(ctx context.Context) ApiCreateShopRequest {
	return ApiCreateShopRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateShop200Response
func (a *ShopsApiService) CreateShopExecute(r ApiCreateShopRequest) (*CreateShop200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateShop200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.CreateShop")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
	}
	if r.createShopRequest == nil {
		return localVarReturnValue, nil, reportError("createShopRequest is required and must be specified")
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
	localVarPostBody = r.createShopRequest
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

type ApiGetShopOfferRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	shopId string
	shopOfferId string
}

func (r ApiGetShopOfferRequest) Execute() (*ShopOffer, *http.Response, error) {
	return r.ApiService.GetShopOfferExecute(r)
}

/*
GetShopOffer Get shop offer

Returns a shop offer object for the provided shopOfferId.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shopId Any shop id within the MetaFab ecosystem.
 @param shopOfferId Any offer id for the shop. Zero, or a positive integer.
 @return ApiGetShopOfferRequest
*/
func (a *ShopsApiService) GetShopOffer(ctx context.Context, shopId string, shopOfferId string) ApiGetShopOfferRequest {
	return ApiGetShopOfferRequest{
		ApiService: a,
		ctx: ctx,
		shopId: shopId,
		shopOfferId: shopOfferId,
	}
}

// Execute executes the request
//  @return ShopOffer
func (a *ShopsApiService) GetShopOfferExecute(r ApiGetShopOfferRequest) (*ShopOffer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ShopOffer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.GetShopOffer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops/{shopId}/items/{shopOfferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shopId"+"}", url.PathEscape(parameterToString(r.shopId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shopOfferId"+"}", url.PathEscape(parameterToString(r.shopOfferId, "")), -1)

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

type ApiGetShopOffersRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	shopId string
}

func (r ApiGetShopOffersRequest) Execute() ([]ShopOffer, *http.Response, error) {
	return r.ApiService.GetShopOffersExecute(r)
}

/*
GetShopOffers Get shop offers

Returns all shop offers as an array of shop offer objects.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shopId Any shop id within the MetaFab ecosystem.
 @return ApiGetShopOffersRequest
*/
func (a *ShopsApiService) GetShopOffers(ctx context.Context, shopId string) ApiGetShopOffersRequest {
	return ApiGetShopOffersRequest{
		ApiService: a,
		ctx: ctx,
		shopId: shopId,
	}
}

// Execute executes the request
//  @return []ShopOffer
func (a *ShopsApiService) GetShopOffersExecute(r ApiGetShopOffersRequest) ([]ShopOffer, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []ShopOffer
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.GetShopOffers")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops/{shopId}/offers"
	localVarPath = strings.Replace(localVarPath, "{"+"shopId"+"}", url.PathEscape(parameterToString(r.shopId, "")), -1)

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

type ApiGetShopsRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	xGameKey *string
}

// The &#x60;publishedKey&#x60; of a specific game. This can be shared or included in game clients, websites, etc.
func (r ApiGetShopsRequest) XGameKey(xGameKey string) ApiGetShopsRequest {
	r.xGameKey = &xGameKey
	return r
}

func (r ApiGetShopsRequest) Execute() ([]GetShops200ResponseInner, *http.Response, error) {
	return r.ApiService.GetShopsExecute(r)
}

/*
GetShops Get shops

Returns an array of active shops for the game associated with the provided `X-Game-Key`.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetShopsRequest
*/
func (a *ShopsApiService) GetShops(ctx context.Context) ApiGetShopsRequest {
	return ApiGetShopsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return []GetShops200ResponseInner
func (a *ShopsApiService) GetShopsExecute(r ApiGetShopsRequest) ([]GetShops200ResponseInner, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []GetShops200ResponseInner
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.GetShops")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops"

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

type ApiRemoveShopOfferRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	shopId string
	shopOfferId string
	xAuthorization *string
	xPassword *string
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiRemoveShopOfferRequest) XAuthorization(xAuthorization string) ApiRemoveShopOfferRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiRemoveShopOfferRequest) XPassword(xPassword string) ApiRemoveShopOfferRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiRemoveShopOfferRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.RemoveShopOfferExecute(r)
}

/*
RemoveShopOffer Remove shop offer

Removes the provided offer by offerId from the provided shop. Removed offers can no longer be used.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shopId Any shop id within the MetaFab ecosystem.
 @param shopOfferId Any offer id for the shop. Zero, or a positive integer.
 @return ApiRemoveShopOfferRequest
*/
func (a *ShopsApiService) RemoveShopOffer(ctx context.Context, shopId string, shopOfferId string) ApiRemoveShopOfferRequest {
	return ApiRemoveShopOfferRequest{
		ApiService: a,
		ctx: ctx,
		shopId: shopId,
		shopOfferId: shopOfferId,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *ShopsApiService) RemoveShopOfferExecute(r ApiRemoveShopOfferRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.RemoveShopOffer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops/{shopId}/offers/{shopOfferId}"
	localVarPath = strings.Replace(localVarPath, "{"+"shopId"+"}", url.PathEscape(parameterToString(r.shopId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shopOfferId"+"}", url.PathEscape(parameterToString(r.shopOfferId, "")), -1)

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

type ApiSetShopOfferRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	shopId string
	xAuthorization *string
	xPassword *string
	setShopOfferRequest *SetShopOfferRequest
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiSetShopOfferRequest) XAuthorization(xAuthorization string) ApiSetShopOfferRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiSetShopOfferRequest) XPassword(xPassword string) ApiSetShopOfferRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiSetShopOfferRequest) SetShopOfferRequest(setShopOfferRequest SetShopOfferRequest) ApiSetShopOfferRequest {
	r.setShopOfferRequest = &setShopOfferRequest
	return r
}

func (r ApiSetShopOfferRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.SetShopOfferExecute(r)
}

/*
SetShopOffer Set shop offer

Sets a new shop offer or updates an existing one for the provided id. Shop offers allow currency to item, item to currency or item to item exchanges.

All request fields besides `id` are optional. Any optional fields omitted will not be used for the offer. This allows you to create many different combinations of offers. For example, you can create an offer that may require 3 unique item ids of specified quantities from a given item collection and gives the user 1 new unique item id in exchange.

Another example, you may want to make a shop offer from one ERC20 token to another. This is also possible - simple set the input and output currency fields and leave the others blank.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shopId Any shop id within the MetaFab ecosystem.
 @return ApiSetShopOfferRequest
*/
func (a *ShopsApiService) SetShopOffer(ctx context.Context, shopId string) ApiSetShopOfferRequest {
	return ApiSetShopOfferRequest{
		ApiService: a,
		ctx: ctx,
		shopId: shopId,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *ShopsApiService) SetShopOfferExecute(r ApiSetShopOfferRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.SetShopOffer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops/{shopId}/offers"
	localVarPath = strings.Replace(localVarPath, "{"+"shopId"+"}", url.PathEscape(parameterToString(r.shopId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
	}
	if r.setShopOfferRequest == nil {
		return localVarReturnValue, nil, reportError("setShopOfferRequest is required and must be specified")
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
	localVarPostBody = r.setShopOfferRequest
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

type ApiUseShopOfferRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	shopId string
	shopOfferId string
	xAuthorization *string
	xPassword *string
}

// The &#x60;secretKey&#x60; of a specific game or the &#x60;accessToken&#x60; of a specific player.
func (r ApiUseShopOfferRequest) XAuthorization(xAuthorization string) ApiUseShopOfferRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game or player. Required to decrypt and perform blockchain transactions with the game or player primary wallet.
func (r ApiUseShopOfferRequest) XPassword(xPassword string) ApiUseShopOfferRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiUseShopOfferRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.UseShopOfferExecute(r)
}

/*
UseShopOffer Use shop offer

Uses a shop offer. The required (input) item(s) and/or currency are removed from the wallet or player wallet using the offer. The given (output) item(s) and/or currency are given to the wallet or player wallet using the offer.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shopId Any shop id within the MetaFab ecosystem.
 @param shopOfferId Any offer id for the shop. Zero, or a positive integer.
 @return ApiUseShopOfferRequest
*/
func (a *ShopsApiService) UseShopOffer(ctx context.Context, shopId string, shopOfferId string) ApiUseShopOfferRequest {
	return ApiUseShopOfferRequest{
		ApiService: a,
		ctx: ctx,
		shopId: shopId,
		shopOfferId: shopOfferId,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *ShopsApiService) UseShopOfferExecute(r ApiUseShopOfferRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.UseShopOffer")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops/{shopId}/offers/{shopOfferId}/uses"
	localVarPath = strings.Replace(localVarPath, "{"+"shopId"+"}", url.PathEscape(parameterToString(r.shopId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"shopOfferId"+"}", url.PathEscape(parameterToString(r.shopOfferId, "")), -1)

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

type ApiWithdrawFromShopRequest struct {
	ctx context.Context
	ApiService *ShopsApiService
	shopId string
	xAuthorization *string
	xPassword *string
	withdrawFromShopRequest *WithdrawFromShopRequest
}

// The &#x60;secretKey&#x60; of the authenticating game.
func (r ApiWithdrawFromShopRequest) XAuthorization(xAuthorization string) ApiWithdrawFromShopRequest {
	r.xAuthorization = &xAuthorization
	return r
}

// The password of the authenticating game. Required to decrypt and perform blockchain transactions with the game primary wallet.
func (r ApiWithdrawFromShopRequest) XPassword(xPassword string) ApiWithdrawFromShopRequest {
	r.xPassword = &xPassword
	return r
}

func (r ApiWithdrawFromShopRequest) WithdrawFromShopRequest(withdrawFromShopRequest WithdrawFromShopRequest) ApiWithdrawFromShopRequest {
	r.withdrawFromShopRequest = &withdrawFromShopRequest
	return r
}

func (r ApiWithdrawFromShopRequest) Execute() (*TransactionModel, *http.Response, error) {
	return r.ApiService.WithdrawFromShopExecute(r)
}

/*
WithdrawFromShop Withdraw from shop

Withdraws native token, currency or items from a shop. Whenever a shop offer has input requirements, the native tokens, currencies or items for the requirements of that offer are deposited into the shop contract when the offer is used. These can be withdrawn to any other address.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param shopId Any shop id within the MetaFab ecosystem.
 @return ApiWithdrawFromShopRequest
*/
func (a *ShopsApiService) WithdrawFromShop(ctx context.Context, shopId string) ApiWithdrawFromShopRequest {
	return ApiWithdrawFromShopRequest{
		ApiService: a,
		ctx: ctx,
		shopId: shopId,
	}
}

// Execute executes the request
//  @return TransactionModel
func (a *ShopsApiService) WithdrawFromShopExecute(r ApiWithdrawFromShopRequest) (*TransactionModel, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *TransactionModel
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShopsApiService.WithdrawFromShop")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/shops/{shopId}/withdrawals"
	localVarPath = strings.Replace(localVarPath, "{"+"shopId"+"}", url.PathEscape(parameterToString(r.shopId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.xAuthorization == nil {
		return localVarReturnValue, nil, reportError("xAuthorization is required and must be specified")
	}
	if r.xPassword == nil {
		return localVarReturnValue, nil, reportError("xPassword is required and must be specified")
	}
	if r.withdrawFromShopRequest == nil {
		return localVarReturnValue, nil, reportError("withdrawFromShopRequest is required and must be specified")
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
	localVarPostBody = r.withdrawFromShopRequest
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
