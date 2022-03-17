package services

import (
	"webapi/models/domains"
	"webapi/models/dtos"
	"webapi/models/mapper"
	"webapi/repositories"
)

type IDemoService interface {
	CreateDemo() (*dtos.DemoDto, *dtos.ErrorDto)
	GetAllDemos() ([]*dtos.DemoDto, *dtos.ErrorDto)
}

type DemoService struct {
	demoRepository repositories.IDemoRepository
}

func NewDemoService(demoRepository repositories.IDemoRepository) IDemoService {
	return &DemoService{
		demoRepository: demoRepository,
	}
}

func (that *DemoService) CreateDemo() (*dtos.DemoDto, *dtos.ErrorDto) {
	var demoToCreate domains.Demo
	createdDemo, err := that.demoRepository.Create(&demoToCreate)
	if err != nil {
		return nil, dtos.NewDatabaseError(err)
	}
	demoDto := &dtos.DemoDto{}
	mapper.Map(createdDemo, demoDto)
	return demoDto, nil
}

func (that *DemoService) GetAllDemos() ([]*dtos.DemoDto, *dtos.ErrorDto) {
	demos, err := that.demoRepository.GetAll()
	if err != nil {
		return nil, dtos.NewDatabaseError(err)
	}
	demoDtos := make([]*dtos.DemoDto, 0)
	mapper.Map(demos, &demoDtos, func(i int, s *domains.Demo, d *dtos.DemoDto) {
	})
	return demoDtos, nil
}
