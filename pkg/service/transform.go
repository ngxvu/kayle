package service

import (
	"gitlab.com/merakilab9/meracrawler/kayle/pkg/repo/pg"
)

type Sizer interface {
	Size() int64
}

type TransformService struct {
	repo pg.PGInterface
}

func NewTransformService(repo pg.PGInterface) TransformInterface {
	return &TransformService{repo: repo}
}

type TransformInterface interface {
}
