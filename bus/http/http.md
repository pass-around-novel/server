# Pass Around Novel HTTP Bus

## Standard Endpoints

A standard endpoint is called by sending a HTTP POST request to the endpoint name.
The parameters are sent as a JSON array in the request body and the response is sent as a JSON array in the response body.

## notify()

Notify will likely block the HTTP request until there is a notification ready to be read or it is getting close to the timeout.
Therefore, this bus, unlike the others, may return an array of length 0.

## get_image()

Instead of the normal method, get_image() works by sending a HTTP GET request to `/get_image/{id}`.

## Errors

When an error occurs, the error code is returned as the response body and a HTTP 400 error is returned.
During normal operation, HTTP 200 is always returned.
