// Code generated by protoc-gen-go-http. DO NOT EDIT.
// Source: queue.proto

package pb

import (
	context "context"
	http1 "net/http"
	strconv "strconv"
	strings "strings"

	transport "github.com/erda-project/erda-infra/pkg/transport"
	http "github.com/erda-project/erda-infra/pkg/transport/http"
	httprule "github.com/erda-project/erda-infra/pkg/transport/http/httprule"
	runtime "github.com/erda-project/erda-infra/pkg/transport/http/runtime"
	urlenc "github.com/erda-project/erda-infra/pkg/urlenc"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the "github.com/erda-project/erda-infra/pkg/transport/http" package it is being compiled against.
const _ = http.SupportPackageIsVersion1

// QueueServiceHandler is the server API for QueueService service.
type QueueServiceHandler interface {
	// POST /api/pipeline-queues
	CreateQueue(context.Context, *QueueCreateRequest) (*QueueCreateResponse, error)
	// GET /api/pipeline-queues/{queueID}
	GetQueue(context.Context, *QueueGetRequest) (*QueueGetResponse, error)
	// GET /api/pipeline-queues
	PagingQueue(context.Context, *QueuePagingRequest) (*QueuePagingResponse, error)
	// PUT /api/pipeline-queues/{queueID}
	UpdateQueue(context.Context, *QueueUpdateRequest) (*QueueUpdateResponse, error)
	// DELETE /api/pipeline-queues/{queueID}
	DeleteQueue(context.Context, *QueueDeleteRequest) (*QueueDeleteResponse, error)
}

// RegisterQueueServiceHandler register QueueServiceHandler to http.Router.
func RegisterQueueServiceHandler(r http.Router, srv QueueServiceHandler, opts ...http.HandleOption) {
	h := http.DefaultHandleOptions()
	for _, op := range opts {
		op(h)
	}
	encodeFunc := func(fn func(http1.ResponseWriter, *http1.Request) (interface{}, error)) http.HandlerFunc {
		handler := func(w http1.ResponseWriter, r *http1.Request) {
			out, err := fn(w, r)
			if err != nil {
				h.Error(w, r, err)
				return
			}
			if err := h.Encode(w, r, out); err != nil {
				h.Error(w, r, err)
			}
		}
		if h.HTTPInterceptor != nil {
			handler = h.HTTPInterceptor(handler)
		}
		return handler
	}

	add_CreateQueue := func(method, path string, fn func(context.Context, *QueueCreateRequest) (*QueueCreateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueueCreateRequest))
		}
		var CreateQueue_info transport.ServiceInfo
		if h.Interceptor != nil {
			CreateQueue_info = transport.NewServiceInfo("erda.core.pipeline.queue.QueueService", "CreateQueue", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, CreateQueue_info)
				}
				r = r.WithContext(ctx)
				var in QueueCreateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_GetQueue := func(method, path string, fn func(context.Context, *QueueGetRequest) (*QueueGetResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueueGetRequest))
		}
		var GetQueue_info transport.ServiceInfo
		if h.Interceptor != nil {
			GetQueue_info = transport.NewServiceInfo("erda.core.pipeline.queue.QueueService", "GetQueue", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, GetQueue_info)
				}
				r = r.WithContext(ctx)
				var in QueueGetRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "queueID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.QueueID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_PagingQueue := func(method, path string, fn func(context.Context, *QueuePagingRequest) (*QueuePagingResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueuePagingRequest))
		}
		var PagingQueue_info transport.ServiceInfo
		if h.Interceptor != nil {
			PagingQueue_info = transport.NewServiceInfo("erda.core.pipeline.queue.QueueService", "PagingQueue", srv)
			handler = h.Interceptor(handler)
		}
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, PagingQueue_info)
				}
				r = r.WithContext(ctx)
				var in QueuePagingRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_UpdateQueue := func(method, path string, fn func(context.Context, *QueueUpdateRequest) (*QueueUpdateResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueueUpdateRequest))
		}
		var UpdateQueue_info transport.ServiceInfo
		if h.Interceptor != nil {
			UpdateQueue_info = transport.NewServiceInfo("erda.core.pipeline.queue.QueueService", "UpdateQueue", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, UpdateQueue_info)
				}
				r = r.WithContext(ctx)
				var in QueueUpdateRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "queueID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.QueueID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_DeleteQueue := func(method, path string, fn func(context.Context, *QueueDeleteRequest) (*QueueDeleteResponse, error)) {
		handler := func(ctx context.Context, req interface{}) (interface{}, error) {
			return fn(ctx, req.(*QueueDeleteRequest))
		}
		var DeleteQueue_info transport.ServiceInfo
		if h.Interceptor != nil {
			DeleteQueue_info = transport.NewServiceInfo("erda.core.pipeline.queue.QueueService", "DeleteQueue", srv)
			handler = h.Interceptor(handler)
		}
		compiler, _ := httprule.Parse(path)
		temp := compiler.Compile()
		pattern, _ := runtime.NewPattern(httprule.SupportPackageIsVersion1, temp.OpCodes, temp.Pool, temp.Verb)
		r.Add(method, path, encodeFunc(
			func(w http1.ResponseWriter, r *http1.Request) (interface{}, error) {
				ctx := http.WithRequest(r.Context(), r)
				ctx = transport.WithHTTPHeaderForServer(ctx, r.Header)
				if h.Interceptor != nil {
					ctx = context.WithValue(ctx, transport.ServiceInfoContextKey, DeleteQueue_info)
				}
				r = r.WithContext(ctx)
				var in QueueDeleteRequest
				if err := h.Decode(r, &in); err != nil {
					return nil, err
				}
				var input interface{} = &in
				if u, ok := (input).(urlenc.URLValuesUnmarshaler); ok {
					if err := u.UnmarshalURLValues("", r.URL.Query()); err != nil {
						return nil, err
					}
				}
				path := r.URL.Path
				if len(path) > 0 {
					components := strings.Split(path[1:], "/")
					last := len(components) - 1
					var verb string
					if idx := strings.LastIndex(components[last], ":"); idx >= 0 {
						c := components[last]
						components[last], verb = c[:idx], c[idx+1:]
					}
					vars, err := pattern.Match(components, verb)
					if err != nil {
						return nil, err
					}
					for k, val := range vars {
						switch k {
						case "queueID":
							val, err := strconv.ParseUint(val, 10, 64)
							if err != nil {
								return nil, err
							}
							in.QueueID = val
						}
					}
				}
				out, err := handler(ctx, &in)
				if err != nil {
					return out, err
				}
				return out, nil
			}),
		)
	}

	add_CreateQueue("POST", "/api/pipeline-queues", srv.CreateQueue)
	add_GetQueue("GET", "/api/pipeline-queues/{queueID}", srv.GetQueue)
	add_PagingQueue("GET", "/api/pipeline-queues", srv.PagingQueue)
	add_UpdateQueue("PUT", "/api/pipeline-queues/{queueID}", srv.UpdateQueue)
	add_DeleteQueue("DELETE", "/api/pipeline-queues/{queueID}", srv.DeleteQueue)
}
