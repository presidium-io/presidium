package local

import (
	"errors"

	"github.com/presid-io/presidio/pkg/platform"
)

func (s *store) CreateJob(name string, image string, commands []string) error {
	return errors.New("Not implemented")
}

func (s *store) CreateCronJob(name string, schedule string, containerDetailsArray []platform.ContainerDetails) error {
	return errors.New("Not implemented")
}

func (s *store) ListJobs() ([]string, error) {
	return nil, errors.New("Not implemented")
}

func (s *store) ListCronJobs() ([]string, error) {
	return nil, errors.New("Not implemented")
}

func (s *store) DeleteJob(name string) error {
	return errors.New("Not implemented")
}

func (s *store) DeleteCronJob(name string) error {
	return errors.New("Not implemented")
}
