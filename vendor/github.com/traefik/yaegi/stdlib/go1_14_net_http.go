// Code generated by 'yaegi extract net/http'. DO NOT EDIT.

// +build go1.14,!go1.15

package stdlib

import (
	"bufio"
	"go/constant"
	"go/token"
	"net"
	"net/http"
	"net/url"
	"os"
	"reflect"
)

func init() {
	Symbols["net/http"] = map[string]reflect.Value{
		// function, constant and variable definitions
		"CanonicalHeaderKey":                  reflect.ValueOf(http.CanonicalHeaderKey),
		"DefaultClient":                       reflect.ValueOf(&http.DefaultClient).Elem(),
		"DefaultMaxHeaderBytes":               reflect.ValueOf(constant.MakeFromLiteral("1048576", token.INT, 0)),
		"DefaultMaxIdleConnsPerHost":          reflect.ValueOf(constant.MakeFromLiteral("2", token.INT, 0)),
		"DefaultServeMux":                     reflect.ValueOf(&http.DefaultServeMux).Elem(),
		"DefaultTransport":                    reflect.ValueOf(&http.DefaultTransport).Elem(),
		"DetectContentType":                   reflect.ValueOf(http.DetectContentType),
		"ErrAbortHandler":                     reflect.ValueOf(&http.ErrAbortHandler).Elem(),
		"ErrBodyNotAllowed":                   reflect.ValueOf(&http.ErrBodyNotAllowed).Elem(),
		"ErrBodyReadAfterClose":               reflect.ValueOf(&http.ErrBodyReadAfterClose).Elem(),
		"ErrContentLength":                    reflect.ValueOf(&http.ErrContentLength).Elem(),
		"ErrHandlerTimeout":                   reflect.ValueOf(&http.ErrHandlerTimeout).Elem(),
		"ErrHeaderTooLong":                    reflect.ValueOf(&http.ErrHeaderTooLong).Elem(),
		"ErrHijacked":                         reflect.ValueOf(&http.ErrHijacked).Elem(),
		"ErrLineTooLong":                      reflect.ValueOf(&http.ErrLineTooLong).Elem(),
		"ErrMissingBoundary":                  reflect.ValueOf(&http.ErrMissingBoundary).Elem(),
		"ErrMissingContentLength":             reflect.ValueOf(&http.ErrMissingContentLength).Elem(),
		"ErrMissingFile":                      reflect.ValueOf(&http.ErrMissingFile).Elem(),
		"ErrNoCookie":                         reflect.ValueOf(&http.ErrNoCookie).Elem(),
		"ErrNoLocation":                       reflect.ValueOf(&http.ErrNoLocation).Elem(),
		"ErrNotMultipart":                     reflect.ValueOf(&http.ErrNotMultipart).Elem(),
		"ErrNotSupported":                     reflect.ValueOf(&http.ErrNotSupported).Elem(),
		"ErrServerClosed":                     reflect.ValueOf(&http.ErrServerClosed).Elem(),
		"ErrShortBody":                        reflect.ValueOf(&http.ErrShortBody).Elem(),
		"ErrSkipAltProtocol":                  reflect.ValueOf(&http.ErrSkipAltProtocol).Elem(),
		"ErrUnexpectedTrailer":                reflect.ValueOf(&http.ErrUnexpectedTrailer).Elem(),
		"ErrUseLastResponse":                  reflect.ValueOf(&http.ErrUseLastResponse).Elem(),
		"ErrWriteAfterFlush":                  reflect.ValueOf(&http.ErrWriteAfterFlush).Elem(),
		"Error":                               reflect.ValueOf(http.Error),
		"FileServer":                          reflect.ValueOf(http.FileServer),
		"Get":                                 reflect.ValueOf(http.Get),
		"Handle":                              reflect.ValueOf(http.Handle),
		"HandleFunc":                          reflect.ValueOf(http.HandleFunc),
		"Head":                                reflect.ValueOf(http.Head),
		"ListenAndServe":                      reflect.ValueOf(http.ListenAndServe),
		"ListenAndServeTLS":                   reflect.ValueOf(http.ListenAndServeTLS),
		"LocalAddrContextKey":                 reflect.ValueOf(&http.LocalAddrContextKey).Elem(),
		"MaxBytesReader":                      reflect.ValueOf(http.MaxBytesReader),
		"MethodConnect":                       reflect.ValueOf(constant.MakeFromLiteral("\"CONNECT\"", token.STRING, 0)),
		"MethodDelete":                        reflect.ValueOf(constant.MakeFromLiteral("\"DELETE\"", token.STRING, 0)),
		"MethodGet":                           reflect.ValueOf(constant.MakeFromLiteral("\"GET\"", token.STRING, 0)),
		"MethodHead":                          reflect.ValueOf(constant.MakeFromLiteral("\"HEAD\"", token.STRING, 0)),
		"MethodOptions":                       reflect.ValueOf(constant.MakeFromLiteral("\"OPTIONS\"", token.STRING, 0)),
		"MethodPatch":                         reflect.ValueOf(constant.MakeFromLiteral("\"PATCH\"", token.STRING, 0)),
		"MethodPost":                          reflect.ValueOf(constant.MakeFromLiteral("\"POST\"", token.STRING, 0)),
		"MethodPut":                           reflect.ValueOf(constant.MakeFromLiteral("\"PUT\"", token.STRING, 0)),
		"MethodTrace":                         reflect.ValueOf(constant.MakeFromLiteral("\"TRACE\"", token.STRING, 0)),
		"NewFileTransport":                    reflect.ValueOf(http.NewFileTransport),
		"NewRequest":                          reflect.ValueOf(http.NewRequest),
		"NewRequestWithContext":               reflect.ValueOf(http.NewRequestWithContext),
		"NewServeMux":                         reflect.ValueOf(http.NewServeMux),
		"NoBody":                              reflect.ValueOf(&http.NoBody).Elem(),
		"NotFound":                            reflect.ValueOf(http.NotFound),
		"NotFoundHandler":                     reflect.ValueOf(http.NotFoundHandler),
		"ParseHTTPVersion":                    reflect.ValueOf(http.ParseHTTPVersion),
		"ParseTime":                           reflect.ValueOf(http.ParseTime),
		"Post":                                reflect.ValueOf(http.Post),
		"PostForm":                            reflect.ValueOf(http.PostForm),
		"ProxyFromEnvironment":                reflect.ValueOf(http.ProxyFromEnvironment),
		"ProxyURL":                            reflect.ValueOf(http.ProxyURL),
		"ReadRequest":                         reflect.ValueOf(http.ReadRequest),
		"ReadResponse":                        reflect.ValueOf(http.ReadResponse),
		"Redirect":                            reflect.ValueOf(http.Redirect),
		"RedirectHandler":                     reflect.ValueOf(http.RedirectHandler),
		"SameSiteDefaultMode":                 reflect.ValueOf(http.SameSiteDefaultMode),
		"SameSiteLaxMode":                     reflect.ValueOf(http.SameSiteLaxMode),
		"SameSiteNoneMode":                    reflect.ValueOf(http.SameSiteNoneMode),
		"SameSiteStrictMode":                  reflect.ValueOf(http.SameSiteStrictMode),
		"Serve":                               reflect.ValueOf(http.Serve),
		"ServeContent":                        reflect.ValueOf(http.ServeContent),
		"ServeFile":                           reflect.ValueOf(http.ServeFile),
		"ServeTLS":                            reflect.ValueOf(http.ServeTLS),
		"ServerContextKey":                    reflect.ValueOf(&http.ServerContextKey).Elem(),
		"SetCookie":                           reflect.ValueOf(http.SetCookie),
		"StateActive":                         reflect.ValueOf(http.StateActive),
		"StateClosed":                         reflect.ValueOf(http.StateClosed),
		"StateHijacked":                       reflect.ValueOf(http.StateHijacked),
		"StateIdle":                           reflect.ValueOf(http.StateIdle),
		"StateNew":                            reflect.ValueOf(http.StateNew),
		"StatusAccepted":                      reflect.ValueOf(constant.MakeFromLiteral("202", token.INT, 0)),
		"StatusAlreadyReported":               reflect.ValueOf(constant.MakeFromLiteral("208", token.INT, 0)),
		"StatusBadGateway":                    reflect.ValueOf(constant.MakeFromLiteral("502", token.INT, 0)),
		"StatusBadRequest":                    reflect.ValueOf(constant.MakeFromLiteral("400", token.INT, 0)),
		"StatusConflict":                      reflect.ValueOf(constant.MakeFromLiteral("409", token.INT, 0)),
		"StatusContinue":                      reflect.ValueOf(constant.MakeFromLiteral("100", token.INT, 0)),
		"StatusCreated":                       reflect.ValueOf(constant.MakeFromLiteral("201", token.INT, 0)),
		"StatusEarlyHints":                    reflect.ValueOf(constant.MakeFromLiteral("103", token.INT, 0)),
		"StatusExpectationFailed":             reflect.ValueOf(constant.MakeFromLiteral("417", token.INT, 0)),
		"StatusFailedDependency":              reflect.ValueOf(constant.MakeFromLiteral("424", token.INT, 0)),
		"StatusForbidden":                     reflect.ValueOf(constant.MakeFromLiteral("403", token.INT, 0)),
		"StatusFound":                         reflect.ValueOf(constant.MakeFromLiteral("302", token.INT, 0)),
		"StatusGatewayTimeout":                reflect.ValueOf(constant.MakeFromLiteral("504", token.INT, 0)),
		"StatusGone":                          reflect.ValueOf(constant.MakeFromLiteral("410", token.INT, 0)),
		"StatusHTTPVersionNotSupported":       reflect.ValueOf(constant.MakeFromLiteral("505", token.INT, 0)),
		"StatusIMUsed":                        reflect.ValueOf(constant.MakeFromLiteral("226", token.INT, 0)),
		"StatusInsufficientStorage":           reflect.ValueOf(constant.MakeFromLiteral("507", token.INT, 0)),
		"StatusInternalServerError":           reflect.ValueOf(constant.MakeFromLiteral("500", token.INT, 0)),
		"StatusLengthRequired":                reflect.ValueOf(constant.MakeFromLiteral("411", token.INT, 0)),
		"StatusLocked":                        reflect.ValueOf(constant.MakeFromLiteral("423", token.INT, 0)),
		"StatusLoopDetected":                  reflect.ValueOf(constant.MakeFromLiteral("508", token.INT, 0)),
		"StatusMethodNotAllowed":              reflect.ValueOf(constant.MakeFromLiteral("405", token.INT, 0)),
		"StatusMisdirectedRequest":            reflect.ValueOf(constant.MakeFromLiteral("421", token.INT, 0)),
		"StatusMovedPermanently":              reflect.ValueOf(constant.MakeFromLiteral("301", token.INT, 0)),
		"StatusMultiStatus":                   reflect.ValueOf(constant.MakeFromLiteral("207", token.INT, 0)),
		"StatusMultipleChoices":               reflect.ValueOf(constant.MakeFromLiteral("300", token.INT, 0)),
		"StatusNetworkAuthenticationRequired": reflect.ValueOf(constant.MakeFromLiteral("511", token.INT, 0)),
		"StatusNoContent":                     reflect.ValueOf(constant.MakeFromLiteral("204", token.INT, 0)),
		"StatusNonAuthoritativeInfo":          reflect.ValueOf(constant.MakeFromLiteral("203", token.INT, 0)),
		"StatusNotAcceptable":                 reflect.ValueOf(constant.MakeFromLiteral("406", token.INT, 0)),
		"StatusNotExtended":                   reflect.ValueOf(constant.MakeFromLiteral("510", token.INT, 0)),
		"StatusNotFound":                      reflect.ValueOf(constant.MakeFromLiteral("404", token.INT, 0)),
		"StatusNotImplemented":                reflect.ValueOf(constant.MakeFromLiteral("501", token.INT, 0)),
		"StatusNotModified":                   reflect.ValueOf(constant.MakeFromLiteral("304", token.INT, 0)),
		"StatusOK":                            reflect.ValueOf(constant.MakeFromLiteral("200", token.INT, 0)),
		"StatusPartialContent":                reflect.ValueOf(constant.MakeFromLiteral("206", token.INT, 0)),
		"StatusPaymentRequired":               reflect.ValueOf(constant.MakeFromLiteral("402", token.INT, 0)),
		"StatusPermanentRedirect":             reflect.ValueOf(constant.MakeFromLiteral("308", token.INT, 0)),
		"StatusPreconditionFailed":            reflect.ValueOf(constant.MakeFromLiteral("412", token.INT, 0)),
		"StatusPreconditionRequired":          reflect.ValueOf(constant.MakeFromLiteral("428", token.INT, 0)),
		"StatusProcessing":                    reflect.ValueOf(constant.MakeFromLiteral("102", token.INT, 0)),
		"StatusProxyAuthRequired":             reflect.ValueOf(constant.MakeFromLiteral("407", token.INT, 0)),
		"StatusRequestEntityTooLarge":         reflect.ValueOf(constant.MakeFromLiteral("413", token.INT, 0)),
		"StatusRequestHeaderFieldsTooLarge":   reflect.ValueOf(constant.MakeFromLiteral("431", token.INT, 0)),
		"StatusRequestTimeout":                reflect.ValueOf(constant.MakeFromLiteral("408", token.INT, 0)),
		"StatusRequestURITooLong":             reflect.ValueOf(constant.MakeFromLiteral("414", token.INT, 0)),
		"StatusRequestedRangeNotSatisfiable":  reflect.ValueOf(constant.MakeFromLiteral("416", token.INT, 0)),
		"StatusResetContent":                  reflect.ValueOf(constant.MakeFromLiteral("205", token.INT, 0)),
		"StatusSeeOther":                      reflect.ValueOf(constant.MakeFromLiteral("303", token.INT, 0)),
		"StatusServiceUnavailable":            reflect.ValueOf(constant.MakeFromLiteral("503", token.INT, 0)),
		"StatusSwitchingProtocols":            reflect.ValueOf(constant.MakeFromLiteral("101", token.INT, 0)),
		"StatusTeapot":                        reflect.ValueOf(constant.MakeFromLiteral("418", token.INT, 0)),
		"StatusTemporaryRedirect":             reflect.ValueOf(constant.MakeFromLiteral("307", token.INT, 0)),
		"StatusText":                          reflect.ValueOf(http.StatusText),
		"StatusTooEarly":                      reflect.ValueOf(constant.MakeFromLiteral("425", token.INT, 0)),
		"StatusTooManyRequests":               reflect.ValueOf(constant.MakeFromLiteral("429", token.INT, 0)),
		"StatusUnauthorized":                  reflect.ValueOf(constant.MakeFromLiteral("401", token.INT, 0)),
		"StatusUnavailableForLegalReasons":    reflect.ValueOf(constant.MakeFromLiteral("451", token.INT, 0)),
		"StatusUnprocessableEntity":           reflect.ValueOf(constant.MakeFromLiteral("422", token.INT, 0)),
		"StatusUnsupportedMediaType":          reflect.ValueOf(constant.MakeFromLiteral("415", token.INT, 0)),
		"StatusUpgradeRequired":               reflect.ValueOf(constant.MakeFromLiteral("426", token.INT, 0)),
		"StatusUseProxy":                      reflect.ValueOf(constant.MakeFromLiteral("305", token.INT, 0)),
		"StatusVariantAlsoNegotiates":         reflect.ValueOf(constant.MakeFromLiteral("506", token.INT, 0)),
		"StripPrefix":                         reflect.ValueOf(http.StripPrefix),
		"TimeFormat":                          reflect.ValueOf(constant.MakeFromLiteral("\"Mon, 02 Jan 2006 15:04:05 GMT\"", token.STRING, 0)),
		"TimeoutHandler":                      reflect.ValueOf(http.TimeoutHandler),
		"TrailerPrefix":                       reflect.ValueOf(constant.MakeFromLiteral("\"Trailer:\"", token.STRING, 0)),

		// type definitions
		"Client":         reflect.ValueOf((*http.Client)(nil)),
		"CloseNotifier":  reflect.ValueOf((*http.CloseNotifier)(nil)),
		"ConnState":      reflect.ValueOf((*http.ConnState)(nil)),
		"Cookie":         reflect.ValueOf((*http.Cookie)(nil)),
		"CookieJar":      reflect.ValueOf((*http.CookieJar)(nil)),
		"Dir":            reflect.ValueOf((*http.Dir)(nil)),
		"File":           reflect.ValueOf((*http.File)(nil)),
		"FileSystem":     reflect.ValueOf((*http.FileSystem)(nil)),
		"Flusher":        reflect.ValueOf((*http.Flusher)(nil)),
		"Handler":        reflect.ValueOf((*http.Handler)(nil)),
		"HandlerFunc":    reflect.ValueOf((*http.HandlerFunc)(nil)),
		"Header":         reflect.ValueOf((*http.Header)(nil)),
		"Hijacker":       reflect.ValueOf((*http.Hijacker)(nil)),
		"ProtocolError":  reflect.ValueOf((*http.ProtocolError)(nil)),
		"PushOptions":    reflect.ValueOf((*http.PushOptions)(nil)),
		"Pusher":         reflect.ValueOf((*http.Pusher)(nil)),
		"Request":        reflect.ValueOf((*http.Request)(nil)),
		"Response":       reflect.ValueOf((*http.Response)(nil)),
		"ResponseWriter": reflect.ValueOf((*http.ResponseWriter)(nil)),
		"RoundTripper":   reflect.ValueOf((*http.RoundTripper)(nil)),
		"SameSite":       reflect.ValueOf((*http.SameSite)(nil)),
		"ServeMux":       reflect.ValueOf((*http.ServeMux)(nil)),
		"Server":         reflect.ValueOf((*http.Server)(nil)),
		"Transport":      reflect.ValueOf((*http.Transport)(nil)),

		// interface wrapper definitions
		"_CloseNotifier":  reflect.ValueOf((*_net_http_CloseNotifier)(nil)),
		"_CookieJar":      reflect.ValueOf((*_net_http_CookieJar)(nil)),
		"_File":           reflect.ValueOf((*_net_http_File)(nil)),
		"_FileSystem":     reflect.ValueOf((*_net_http_FileSystem)(nil)),
		"_Flusher":        reflect.ValueOf((*_net_http_Flusher)(nil)),
		"_Handler":        reflect.ValueOf((*_net_http_Handler)(nil)),
		"_Hijacker":       reflect.ValueOf((*_net_http_Hijacker)(nil)),
		"_Pusher":         reflect.ValueOf((*_net_http_Pusher)(nil)),
		"_ResponseWriter": reflect.ValueOf((*_net_http_ResponseWriter)(nil)),
		"_RoundTripper":   reflect.ValueOf((*_net_http_RoundTripper)(nil)),
	}
}

// _net_http_CloseNotifier is an interface wrapper for CloseNotifier type
type _net_http_CloseNotifier struct {
	WCloseNotify func() <-chan bool
}

func (W _net_http_CloseNotifier) CloseNotify() <-chan bool { return W.WCloseNotify() }

// _net_http_CookieJar is an interface wrapper for CookieJar type
type _net_http_CookieJar struct {
	WCookies    func(u *url.URL) []*http.Cookie
	WSetCookies func(u *url.URL, cookies []*http.Cookie)
}

func (W _net_http_CookieJar) Cookies(u *url.URL) []*http.Cookie { return W.WCookies(u) }
func (W _net_http_CookieJar) SetCookies(u *url.URL, cookies []*http.Cookie) {
	W.WSetCookies(u, cookies)
}

// _net_http_File is an interface wrapper for File type
type _net_http_File struct {
	WClose   func() error
	WRead    func(p []byte) (n int, err error)
	WReaddir func(count int) ([]os.FileInfo, error)
	WSeek    func(offset int64, whence int) (int64, error)
	WStat    func() (os.FileInfo, error)
}

func (W _net_http_File) Close() error                                 { return W.WClose() }
func (W _net_http_File) Read(p []byte) (n int, err error)             { return W.WRead(p) }
func (W _net_http_File) Readdir(count int) ([]os.FileInfo, error)     { return W.WReaddir(count) }
func (W _net_http_File) Seek(offset int64, whence int) (int64, error) { return W.WSeek(offset, whence) }
func (W _net_http_File) Stat() (os.FileInfo, error)                   { return W.WStat() }

// _net_http_FileSystem is an interface wrapper for FileSystem type
type _net_http_FileSystem struct {
	WOpen func(name string) (http.File, error)
}

func (W _net_http_FileSystem) Open(name string) (http.File, error) { return W.WOpen(name) }

// _net_http_Flusher is an interface wrapper for Flusher type
type _net_http_Flusher struct {
	WFlush func()
}

func (W _net_http_Flusher) Flush() { W.WFlush() }

// _net_http_Handler is an interface wrapper for Handler type
type _net_http_Handler struct {
	WServeHTTP func(a0 http.ResponseWriter, a1 *http.Request)
}

func (W _net_http_Handler) ServeHTTP(a0 http.ResponseWriter, a1 *http.Request) { W.WServeHTTP(a0, a1) }

// _net_http_Hijacker is an interface wrapper for Hijacker type
type _net_http_Hijacker struct {
	WHijack func() (net.Conn, *bufio.ReadWriter, error)
}

func (W _net_http_Hijacker) Hijack() (net.Conn, *bufio.ReadWriter, error) { return W.WHijack() }

// _net_http_Pusher is an interface wrapper for Pusher type
type _net_http_Pusher struct {
	WPush func(target string, opts *http.PushOptions) error
}

func (W _net_http_Pusher) Push(target string, opts *http.PushOptions) error {
	return W.WPush(target, opts)
}

// _net_http_ResponseWriter is an interface wrapper for ResponseWriter type
type _net_http_ResponseWriter struct {
	WHeader      func() http.Header
	WWrite       func(a0 []byte) (int, error)
	WWriteHeader func(statusCode int)
}

func (W _net_http_ResponseWriter) Header() http.Header          { return W.WHeader() }
func (W _net_http_ResponseWriter) Write(a0 []byte) (int, error) { return W.WWrite(a0) }
func (W _net_http_ResponseWriter) WriteHeader(statusCode int)   { W.WWriteHeader(statusCode) }

// _net_http_RoundTripper is an interface wrapper for RoundTripper type
type _net_http_RoundTripper struct {
	WRoundTrip func(a0 *http.Request) (*http.Response, error)
}

func (W _net_http_RoundTripper) RoundTrip(a0 *http.Request) (*http.Response, error) {
	return W.WRoundTrip(a0)
}
