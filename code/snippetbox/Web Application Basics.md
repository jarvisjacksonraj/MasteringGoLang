# Web Application Basics

## Three absolute essentials:

1.  Handlers

        Responsible for executing your application logic and for writing HTTP response headers and bodies.

2.  Router (or servemux in Go terminology)

        Stores a mapping between the URL patterns for your application and the corresponding handlers. One servemux for your application containing all your routes.

3.  web server

        Go can establish a web server and listen for incoming requests as part of application itself. No need for an external third-party server like Nginx or Apache.

## Building Block of Application

    1. handler function writes a byte slice containing "<ANY-TEXT>" as the response body.

    2. http.ResponseWriter parameter provides methods for assembling a HTTP response and sending it to the user

    3. *http.Request parameter is a pointer to a struct which holds information about the current request (like the HTTP method and the URL being requested)

    4. http.NewServeMux() function initialize a new servemux, then registration of the function as the handler for the URL patterns. [Ex: "/" pattern]

    5. http.ListenAndServe() function will start a new web server. It take two parameters: the TCP network address to listen on [EX ":4000"] and the servemux [Ex: mux].

    6. The TCP network address passed to http.ListenAndServe() should be in the format "host:port". If host is omitted [it will be like ":4000"].

    7. when http.ListenAndServe() returns an error, log.Fatal() can be used to log the error message and exit.

    8. Note that any error returned by http.ListenAndServe() is always non-nil.

    9. When code runs, it should start a web server listening on port <4000> of local machine.
       To run Application:
                          go run main.go

    10. Each time the server receives a new HTTP request it will pass the request on to the servemux and — in turn — the servemux will check the URL path and dispatch the request to the matching handler.

    11. Go’s servemux treats the URL pattern "/" like a catch-all.
