package pg

import (
	"context"
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/model"
	"gorm.io/gorm"
)

func (r *RepoPG) CreateMedia(ctx context.Context, ob *model.Media, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	return tx.Debug().Create(ob).Error
}

func (r *RepoPG) UpdateMedia(ctx context.Context, update *model.Media, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}

	return tx.WithContext(ctx).Where("id = ?", update.ID).Save(&update).Error
}

func (r *RepoPG) DeleteMedia(ctx context.Context, d *model.Media, tx *gorm.DB) error {
	var cancel context.CancelFunc
	if tx == nil {
		tx, cancel = r.DBWithTimeout(ctx)
		defer cancel()
	}
	return tx.WithContext(ctx).Delete(&d).Error
}
