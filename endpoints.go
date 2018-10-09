package butterflytracker

import(
	"context"
	"github.com/go-kit/kit/endpoint"
)

// Butterfly structure
type Butterfly struct {
	Name string `json:"name"`
}

type trackButterflyRequest struct{
	Butterfly Butterfly
}

type trackButterflyResponse struct {
	Err error `json:"error:,omitempty"`
}
func (r trackButterflyResponse) error() error { return r.Err }

//Endpoints collects all of the endpoints that compose a profile service
type Endpoints struct{
	TrackButterflyEndpoint	endpoint.Endpoint
}
//Service: make endpoint
func MakeTrackButterflyEndpoint(svc Service) endpoint.Endpoint{
	return func(ctx context.Context, request interface{}) (interface{}, error){
		req := request.(trackButterflyRequest)
		e := svc.Track(ctx, req.Butterfly)
		return trackButterflyResponse{Err: e}, nil
	}
}