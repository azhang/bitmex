/* 
 * BitMEX API
 *
 * ## REST API for the BitMEX Trading Platform  [View Changelog](/app/apiChangelog)    #### Getting Started   ##### Fetching Data  All REST endpoints are documented below. You can try out any query right from this interface.  Most table queries accept `count`, `start`, and `reverse` params. Set `reverse=true` to get rows newest-first.  Additional documentation regarding filters, timestamps, and authentication is available in [the main API documentation](https://www.bitmex.com/app/restAPI).  *All* table data is available via the [Websocket](/app/wsAPI). We highly recommend using the socket if you want to have the quickest possible data without being subject to ratelimits.  ##### Return Types  By default, all data is returned as JSON. Send `?_format=csv` to get CSV data or `?_format=xml` to get XML data.  ##### Trade Data Queries  *This is only a small subset of what is available, to get you started.*  Fill in the parameters and click the `Try it out!` button to try any of these queries.  * [Pricing Data](#!/Quote/Quote_get)  * [Trade Data](#!/Trade/Trade_get)  * [OrderBook Data](#!/OrderBook/OrderBook_getL2)  * [Settlement Data](#!/Settlement/Settlement_get)  * [Exchange Statistics](#!/Stats/Stats_history)  Every function of the BitMEX.com platform is exposed here and documented. Many more functions are available.  ##### Swagger Specification  [⇩ Download Swagger JSON](swagger.json)    ## All API Endpoints  Click to expand a section. 
 *
 * OpenAPI spec version: 1.2.0
 * Contact: support@bitmex.com
 * Generated by: https://github.com/swagger-api/swagger-codegen.git
 */

package swagger

import (
	"net/url"
	"net/http"
	"strings"
	"golang.org/x/net/context"
	"time"
	"encoding/json"
)

// Linger please
var (
	_ context.Context
)

type ExecutionApiService service

/* ExecutionApiService Get all raw executions for your account.
 This returns all raw transactions, which includes order opening and cancelation, and order status changes. It can be quite noisy. More focused information is available at &#x60;/execution/tradeHistory&#x60;.  You may also use the &#x60;filter&#x60; param to target your query. Specify an array as a filter value, such as &#x60;{\&quot;execType\&quot;: [\&quot;Settlement\&quot;, \&quot;Trade\&quot;]}&#x60; to filter on multiple values.  See [the FIX Spec](http://www.onixs.biz/fix-dictionary/5.0.SP2/msgType_8_8.html) for explanations of these fields. 
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "symbol" (string) Instrument symbol. Send a bare series (e.g. XBU) to get data for the nearest expiring contract in that series.  You can also send a timeframe, e.g. &#x60;XBU:monthly&#x60;. Timeframes are &#x60;daily&#x60;, &#x60;weekly&#x60;, &#x60;monthly&#x60;, &#x60;quarterly&#x60;, and &#x60;biquarterly&#x60;.
     @param "filter" (string) Generic table filter. Send JSON key/value pairs, such as &#x60;{\&quot;key\&quot;: \&quot;value\&quot;}&#x60;. You can key on individual fields, and do more advanced querying on timestamps. See the [Timestamp Docs](https://www.bitmex.com/app/restAPI#timestamp-filters) for more details.
     @param "columns" (string) Array of column names to fetch. If omitted, will return all columns.  Note that this method will always return item keys, even when not specified, so you may receive more columns that you expect.
     @param "count" (float32) Number of results to fetch.
     @param "start" (float32) Starting point for results.
     @param "reverse" (bool) If true, will sort results newest first.
     @param "startTime" (time.Time) Starting date filter for results.
     @param "endTime" (time.Time) Ending date filter for results.
 @return []Execution*/
func (a *ExecutionApiService) ExecutionGet(ctx context.Context, localVarOptionals map[string]interface{}) ([]Execution, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []Execution
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/execution"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["symbol"], "string", "symbol"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filter"], "string", "filter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["columns"], "string", "columns"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["count"], "float32", "count"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["start"], "float32", "start"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["reverse"], "bool", "reverse"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "time.Time", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "time.Time", "endTime"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["symbol"].(string); localVarOk {
		localVarQueryParams.Add("symbol", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].(string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["columns"].(string); localVarOk {
		localVarQueryParams.Add("columns", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["count"].(float32); localVarOk {
		localVarQueryParams.Add("count", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["start"].(float32); localVarOk {
		localVarQueryParams.Add("start", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["reverse"].(bool); localVarOk {
		localVarQueryParams.Add("reverse", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(time.Time); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(time.Time); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded",}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-nonce"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-signature"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}

/* ExecutionApiService Get all balance-affecting executions. This includes each trade, insurance charge, and settlement.
 * @param ctx context.Context Authentication Context 
 @param optional (nil or map[string]interface{}) with one or more of:
     @param "symbol" (string) Instrument symbol. Send a bare series (e.g. XBU) to get data for the nearest expiring contract in that series.  You can also send a timeframe, e.g. &#x60;XBU:monthly&#x60;. Timeframes are &#x60;daily&#x60;, &#x60;weekly&#x60;, &#x60;monthly&#x60;, &#x60;quarterly&#x60;, and &#x60;biquarterly&#x60;.
     @param "filter" (string) Generic table filter. Send JSON key/value pairs, such as &#x60;{\&quot;key\&quot;: \&quot;value\&quot;}&#x60;. You can key on individual fields, and do more advanced querying on timestamps. See the [Timestamp Docs](https://www.bitmex.com/app/restAPI#timestamp-filters) for more details.
     @param "columns" (string) Array of column names to fetch. If omitted, will return all columns.  Note that this method will always return item keys, even when not specified, so you may receive more columns that you expect.
     @param "count" (float32) Number of results to fetch.
     @param "start" (float32) Starting point for results.
     @param "reverse" (bool) If true, will sort results newest first.
     @param "startTime" (time.Time) Starting date filter for results.
     @param "endTime" (time.Time) Ending date filter for results.
 @return []Execution*/
func (a *ExecutionApiService) ExecutionGetTradeHistory(ctx context.Context, localVarOptionals map[string]interface{}) ([]Execution, *http.Response, error) {
	var (
		localVarHttpMethod = strings.ToUpper("Get")
		localVarPostBody   interface{}
		localVarFileName   string
		localVarFileBytes  []byte
		successPayload     []Execution
	)

	// create path and map variables
	localVarPath := a.client.cfg.BasePath + "/execution/tradeHistory"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if err := typeCheckParameter(localVarOptionals["symbol"], "string", "symbol"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["filter"], "string", "filter"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["columns"], "string", "columns"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["count"], "float32", "count"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["start"], "float32", "start"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["reverse"], "bool", "reverse"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["startTime"], "time.Time", "startTime"); err != nil {
		return successPayload, nil, err
	}
	if err := typeCheckParameter(localVarOptionals["endTime"], "time.Time", "endTime"); err != nil {
		return successPayload, nil, err
	}

	if localVarTempParam, localVarOk := localVarOptionals["symbol"].(string); localVarOk {
		localVarQueryParams.Add("symbol", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["filter"].(string); localVarOk {
		localVarQueryParams.Add("filter", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["columns"].(string); localVarOk {
		localVarQueryParams.Add("columns", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["count"].(float32); localVarOk {
		localVarQueryParams.Add("count", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["start"].(float32); localVarOk {
		localVarQueryParams.Add("start", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["reverse"].(bool); localVarOk {
		localVarQueryParams.Add("reverse", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["startTime"].(time.Time); localVarOk {
		localVarQueryParams.Add("startTime", parameterToString(localVarTempParam, ""))
	}
	if localVarTempParam, localVarOk := localVarOptionals["endTime"].(time.Time); localVarOk {
		localVarQueryParams.Add("endTime", parameterToString(localVarTempParam, ""))
	}
	// to determine the Content-Type header
	localVarHttpContentTypes := []string{"application/json", "application/x-www-form-urlencoded",}

	// set Content-Type header
	localVarHttpContentType := selectHeaderContentType(localVarHttpContentTypes)
	if localVarHttpContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHttpContentType
	}

	// to determine the Accept header
	localVarHttpHeaderAccepts := []string{
		"application/json",
		"application/xml",
		"text/xml",
		"application/javascript",
		"text/javascript",
	}

	// set Accept header
	localVarHttpHeaderAccept := selectHeaderAccept(localVarHttpHeaderAccepts)
	if localVarHttpHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHttpHeaderAccept
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-key"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-nonce"] = key
		}
	}
	if ctx != nil {
		// API Key Authentication
		if auth, ok := ctx.Value(ContextAPIKey).(APIKey); ok {
			var key string
			if auth.Prefix != "" {
				key = auth.Prefix + " " + auth.Key
			} else {
				key = auth.Key
			}
			localVarHeaderParams["api-signature"] = key
		}
	}
	r, err := a.client.prepareRequest(ctx, localVarPath, localVarHttpMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFileName, localVarFileBytes)
	if err != nil {
		return successPayload, nil, err
	}

	localVarHttpResponse, err := a.client.callAPI(r)
	if err != nil || localVarHttpResponse == nil {
		return successPayload, localVarHttpResponse, err
	}
	defer localVarHttpResponse.Body.Close()
	if localVarHttpResponse.StatusCode >= 300 {
		return successPayload, localVarHttpResponse, reportError(localVarHttpResponse.Status)
	}

	if err = json.NewDecoder(localVarHttpResponse.Body).Decode(&successPayload); err != nil {
		return successPayload, localVarHttpResponse, err
	}

	return successPayload, localVarHttpResponse, err
}
