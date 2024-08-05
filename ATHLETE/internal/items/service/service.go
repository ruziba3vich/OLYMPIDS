package service

import (
	"github.com/ruziba3vich/OLYMPIDS/ATHLETE/internal/items/repository"
)

type (
	Service struct {
		storage repository.IAthleteRepo
	}
)

func New(storage repository.IAthleteRepo) repository.IAthleteRepo {
	return &Service{
		storage: storage,
	}
}
