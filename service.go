package main

import (
	"context"
	"fmt"
	"github.com/pete911/zap-examples/logger"
	"go.uber.org/zap"
)

// --- service A ---

type ServiceA struct {
	logger *zap.Logger
	store  Store
}

func NewServiceA(logger *zap.Logger, store Store) ServiceA {

	return ServiceA{
		logger: logger.With(zap.String("component", "service-a")),
		store:  store,
	}
}

func (s ServiceA) GetUser(ctx context.Context) error {

	ctx = logger.GetRequestContext(ctx, "GetUser")
	s.logger.Debug("get user", logger.GetLoggerFields(ctx)...)

	user, err := s.store.GetUser(ctx)
	if err != nil {
		s.logger.Error("get user", logger.GetLoggerFields(ctx, zap.Error(err))...)
		// hide internal log details and return only request id to the user, so we can track error by request
		return fmt.Errorf("internal server error, request id: %s", logger.GetRequestContextId(ctx))
	}

	s.logger.Debug(fmt.Sprintf("user %s", user), logger.GetLoggerFields(ctx)...)
	return nil
}

// --- service B ---

type ServiceB struct {

	logger *zap.Logger
	store  Store
}

func NewServiceB(logger *zap.Logger, store Store) ServiceB {

	return ServiceB{
		logger: logger.With(zap.String("component", "service-b")),
		store:  store,
	}
}

func (s ServiceB) GetUser(ctx context.Context) error {

	ctx = logger.GetRequestContext(ctx, "GetUser")
	s.logger.Debug("get user", logger.GetLoggerFields(ctx)...)

	user, err := s.store.GetUser(ctx)
	if err != nil {
		s.logger.Error("get user", logger.GetLoggerFields(ctx, zap.Error(err))...)
		// hide internal log details and return only request id to the user, so we can track error by request
		return fmt.Errorf("internal server error, request id: %s", logger.GetRequestContextId(ctx))
	}

	s.logger.Debug(fmt.Sprintf("user %s", user), logger.GetLoggerFields(ctx)...)
	return nil
}
