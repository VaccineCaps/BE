package services

import (
	config "BE/configs"
	"BE/domain"
	"BE/model"
)

type svcVaccine struct {
	c    config.Config
	repo domain.AdapterRepositoryVaccineHospital
}

func (s *svcVaccine) CreateStokService(stok model.VaccineHospitals) error {
	return s.repo.CreateStokVaccineHospital(stok)
}

func (s *svcVaccine) UpdateVaccineStokService(hospital_id, vaccine_id int, stok model.VaccineHospitals) error {
	return s.repo.UpdateStokByID(hospital_id, vaccine_id, stok)
}

func (s *svcVaccine) GetAllStokByHospitalService(hospital_id int) (stok []model.VaccineHospitals, err error) {

	return s.repo.GetAllStokByHospital(hospital_id)
}

func (s *svcVaccine) GetStokByHospitalVaccineService(hospital_id, vaccine_id int) (stok model.VaccineHospitals, err error) {
	return s.repo.GetStokByHospitalVaccine(hospital_id, vaccine_id)
}

func (s *svcVaccine) DeleteVaccineStokByIDService(hospital_id, vaccine_id int) error {
	return s.repo.DeleteVaccineStokByID(hospital_id, vaccine_id)
}

func NewServiceVaccineHospital(repo domain.AdapterRepositoryVaccineHospital, c config.Config) domain.AdapterServiceVaccineHospital {
	return &svcVaccine{
		repo: repo,
		c:    c,
	}
}
