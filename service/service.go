package service

import "context"

// Service is meant to be a layer that holds all the business logic functions
// All service types must contain Serve function

type Service interface {
	Serve(ctx context.Context, inputOutput ...interface{}) (errorCode int, errorMessage error)
}
