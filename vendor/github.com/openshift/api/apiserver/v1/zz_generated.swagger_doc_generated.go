package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_APIRequestCount = map[string]string{
	"":       "APIRequestCount tracks requests made to an API. The instance name must be of the form `resource.version.group`, matching the resource.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
	"spec":   "spec defines the characteristics of the resource.",
	"status": "status contains the observed state of the resource.",
}

func (APIRequestCount) SwaggerDoc() map[string]string {
	return map_APIRequestCount
}

var map_APIRequestCountList = map[string]string{
	"": "APIRequestCountList is a list of APIRequestCount resources.\n\nCompatibility level 1: Stable within a major release for a minimum of 12 months or 3 minor releases (whichever is longer).",
}

func (APIRequestCountList) SwaggerDoc() map[string]string {
	return map_APIRequestCountList
}

var map_APIRequestCountSpec = map[string]string{
	"numberOfUsersToReport": "numberOfUsersToReport is the number of users to include in the report. If unspecified or zero, the default is ten.  This is default is subject to change.",
}

func (APIRequestCountSpec) SwaggerDoc() map[string]string {
	return map_APIRequestCountSpec
}

var map_APIRequestCountStatus = map[string]string{
	"conditions":       "conditions contains details of the current status of this API Resource.",
	"removedInRelease": "removedInRelease is when the API will be removed.",
	"requestCount":     "requestCount is a sum of all requestCounts across all current hours, nodes, and users.",
	"currentHour":      "currentHour contains request history for the current hour. This is porcelain to make the API easier to read by humans seeing if they addressed a problem. This field is reset on the hour.",
	"last24h":          "last24h contains request history for the last 24 hours, indexed by the hour, so 12:00AM-12:59 is in index 0, 6am-6:59am is index 6, etc. The index of the current hour is updated live and then duplicated into the requestsLastHour field.",
}

func (APIRequestCountStatus) SwaggerDoc() map[string]string {
	return map_APIRequestCountStatus
}

var map_PerNodeAPIRequestLog = map[string]string{
	"":             "PerNodeAPIRequestLog contains logs of requests to a certain node.",
	"nodeName":     "nodeName where the request are being handled.",
	"requestCount": "requestCount is a sum of all requestCounts across all users, even those outside of the top 10 users.",
	"byUser":       "byUser contains request details by top .spec.numberOfUsersToReport users. Note that because in the case of an apiserver, restart the list of top users is determined on a best-effort basis, the list might be imprecise. In addition, some system users may be explicitly included in the list.",
}

func (PerNodeAPIRequestLog) SwaggerDoc() map[string]string {
	return map_PerNodeAPIRequestLog
}

var map_PerResourceAPIRequestLog = map[string]string{
	"":             "PerResourceAPIRequestLog logs request for various nodes.",
	"byNode":       "byNode contains logs of requests per node.",
	"requestCount": "requestCount is a sum of all requestCounts across nodes.",
}

func (PerResourceAPIRequestLog) SwaggerDoc() map[string]string {
	return map_PerResourceAPIRequestLog
}

var map_PerUserAPIRequestCount = map[string]string{
	"":             "PerUserAPIRequestCount contains logs of a user's requests.",
	"username":     "userName that made the request.",
	"userAgent":    "userAgent that made the request. The same user often has multiple binaries which connect (pods with many containers).  The different binaries will have different userAgents, but the same user.  In addition, we have userAgents with version information embedded and the userName isn't likely to change.",
	"requestCount": "requestCount of requests by the user across all verbs.",
	"byVerb":       "byVerb details by verb.",
}

func (PerUserAPIRequestCount) SwaggerDoc() map[string]string {
	return map_PerUserAPIRequestCount
}

var map_PerVerbAPIRequestCount = map[string]string{
	"":             "PerVerbAPIRequestCount requestCounts requests by API request verb.",
	"verb":         "verb of API request (get, list, create, etc...)",
	"requestCount": "requestCount of requests for verb.",
}

func (PerVerbAPIRequestCount) SwaggerDoc() map[string]string {
	return map_PerVerbAPIRequestCount
}

// AUTO-GENERATED FUNCTIONS END HERE
