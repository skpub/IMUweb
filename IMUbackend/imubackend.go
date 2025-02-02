package imubackendapi

import (
	imubackend "IMUbackend/gen/imubackend"
	"context"

	"goa.design/clue/log"
)

// imubackend service example implementation.
// The example methods log the requests and return zero values.
type imubackendsrvc struct{}

// NewImubackend returns the imubackend service implementation.
func NewImubackend() imubackend.Service {
	return &imubackendsrvc{}
}

// create markdown file.
func (s *imubackendsrvc) Create(ctx context.Context, p *imubackend.Markdown) (err error) {
	log.Printf(ctx, "imubackend.create")
	return
}
