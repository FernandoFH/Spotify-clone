package patients

import (
	"database/sql"
	"github.com/tomiok/patients-API/patients/domain"
)

type PatientInfrastructure interface {
	CreatePatient(p *patients.CreatePatientCMD) (*patients.Patient, error)
	GetPatients() []*patients.Patient
	GetPatientByID(id int64) (*patients.Patient, error)
}
type CreatePatientInDB struct {
	PatientStorage
}

func (c *CreatePatientInDB) CreatePatient(p *patients.CreatePatientCMD) (*patients.Patient, error) {
	return c.createPatientDB(p)
}

func (c *CreatePatientInDB) GetPatients() []*patients.Patient {
	return c.getPatientsDB()
}

func (c *CreatePatientInDB) GetPatientByID(id int64) (*patients.Patient, error) {
	return c.getPatientByIDBD(id)
}

func NewPatientInfrastructure(db *sql.DB) PatientInfrastructure {
	return &CreatePatientInDB{NewPatientStorageInfrastructure(db)}
}