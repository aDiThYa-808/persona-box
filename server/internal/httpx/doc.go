/*
Contains helper methods to handle http requests and responses.
Exposes the following three methods:
ExtractToken() extracts access token from the request header and removes "bearer " prefix;
WriteJSONError() returns an http response with the provided error message and status code;
WriteJSONSuccess() returns an http success response with the provided content s of type any;
*/
package httpx
