package entity_profile

import "server-client/domain/entities"

type Profile struct {
	entities.Model
	Name string
}

type profile struct {
	repository Repository
}

func New(r Repository) *profile {
	return &profile{repository: r}
}

func (p *profile) Create(name string) (*Profile, error) {
	return p.repository.Create(name)
}
