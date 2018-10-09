package butterflytracker
import(
	"context"
	"net/http"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"github.com/gorilla/mux"
	httptransport "github.com/go-kit/kit/transport/http"
)
type errorer interface {
	error() error
}

//Transform json input into Buttefly object
func decodeTrackButterflyRequest(_ context.Context, r *http.Request) (interface{}, error){
	var req trackButterflyRequest
	if e := json.NewDecoder(r.Body).Decode(&req.Butterfly); e != nil {
		return nil, e
	}
	return req, nil
}

//Transform response object to json
func encodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	if e, ok := response.(errorer); ok && e.error() != nil {
		// Not a Go kit transport error, but a business-logic error.
		// Provide those as HTTP errors.
		
		encodeError(ctx, e.error(), w)
		return nil
	}
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusCreated)
	return json.NewEncoder(w).Encode(response)
}

func encodeError(_ context.Context, err error, w http.ResponseWriter) {
	if err == nil {
		panic("encodeError with nil error")
	}
	w.Header().Set("Content-Type", "application/json; charset=uff-8")
	w.WriteHeader(http.StatusConflict)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"error": err.Error(),
	})
}

//MakeHTTPHandler 
func MakeHTTPHandler(ctx context.Context, endpoints Endpoints, logger log.Logger) http.Handler{
	r := mux.NewRouter()
	options := []httptransport.ServerOption{
		httptransport.ServerErrorLogger(logger),
		httptransport.ServerErrorEncoder(encodeError),
	}
	r.Methods("POST").Path("/butterflies/").Handler(httptransport.NewServer(
		endpoints.TrackButterflyEndpoint,
		decodeTrackButterflyRequest,
		encodeResponse,
		options...,
	))
	return r
}