package pg

import (
	"context"
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/model"
	"gorm.io/gorm"
	"time"
)

//go:generate mockgen -source base_repo.go -destination mocks/base_repo.go

const (
	generalQueryTimeout = 60 * time.Second
	defaultPageSize     = 30
	maxPageSize         = 1000
)

type RepoPG struct {
	DB    *gorm.DB
	debug bool
}

func (r *RepoPG) GetRepo() *gorm.DB {
	return r.DB
}

func NewPGRepo(db *gorm.DB) PGInterface {
	return &RepoPG{DB: db}
}

type PGInterface interface {
	// DB
	DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc)

	CreateMedia(ctx context.Context, ob *model.Media, tx *gorm.DB) error
	UpdateMedia(ctx context.Context, update *model.Media, tx *gorm.DB) error
	DeleteMedia(ctx context.Context, d *model.Media, tx *gorm.DB) error
}

func (r *RepoPG) DBWithTimeout(ctx context.Context) (*gorm.DB, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(ctx, generalQueryTimeout)
	return r.DB.WithContext(ctx), cancel
}
