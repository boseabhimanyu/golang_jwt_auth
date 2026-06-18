package app

import (
	"context"
	"fmt"
	"golang-jwt-auth/internal/config"
	"golang-jwt-auth/internal/db"
	"time"

	"go.mongodb.org/mongo-driver/v2/mongo"
)

type App struct {
	Config      config.Config
	MongoClient *mongo.Client
	DB          *mongo.Database
}

func New(ctx context.Context) (*App, error) {
	//load env first
	cfg, err := config.Load()
	if err != nil {
		return nil, err
	}
	// do the db connection second
	mongoCli, err := db.Connect(ctx, cfg)
	if err != nil {
		return nil, err
	}

	return &App{
		Config:      cfg,
		MongoClient: mongoCli.Client,
		DB:          mongoCli.DB,
	}, nil
}

func (a *App) Close(ctx context.Context) error {
	if a.MongoClient == nil {
		return nil
	}
	closeCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := a.MongoClient.Disconnect(closeCtx); err != nil {
		return fmt.Errorf("Mongo disconnect failed: %w", err)
	}

	return nil
}
