package profile

import (
	"context"
	"google.golang.org/grpc"
	profileProtos "profile/protos/gen/dota_traker.profile.v1"
)

type Api struct {
	profileProtos.UnimplementedProfileServiceServer
}

func ServerApi(grpc *grpc.Server) {
	profileProtos.RegisterProfileServiceServer(grpc, &Api{})
}

func (p *Api) ProfileInfo(ctx context.Context, r *profileProtos.ProfileRequest) (*profileProtos.ProfileResponse, error) {
	return &profileProtos.ProfileResponse{}, nil
}
