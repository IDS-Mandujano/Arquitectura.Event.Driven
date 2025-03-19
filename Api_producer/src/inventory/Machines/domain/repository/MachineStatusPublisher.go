package repository

type MachineStatusPublisher interface {
    PublishMachineStatus(machineID int, status string) error
}