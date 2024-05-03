package profile_service

import (
	"context"
	"log/slog"
)

type Profile struct {
	log             *slog.Logger
	profileProvider ProfileProvider
}

type ProfileProvider interface {
	ProfileInfo(ctx context.Context, id int64) error
}

func (p *Profile) ProfileAllInfo(ctx context.Context, id int64) {
	if err := p.profileProvider.ProfileInfo(ctx, id); err != nil {
		return
	}
}
