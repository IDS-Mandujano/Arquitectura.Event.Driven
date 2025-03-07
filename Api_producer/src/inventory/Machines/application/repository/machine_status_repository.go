package repository

type MachineStatusRepository interface {
	PublishMachineStatus(machineID int, status string) error
}