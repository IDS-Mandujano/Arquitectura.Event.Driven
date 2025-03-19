package machine

import (
	"gym-system/src/core"
	machineusecases "gym-system/src/inventory/Machines/application/useCases"
	"gym-system/src/inventory/Machines/domain/repository"
	"gym-system/src/inventory/Machines/infraestructure/adapters"
	machineControllers "gym-system/src/inventory/Machines/infraestructure/controllers"
	"log"

	"github.com/gin-gonic/gin"
)

func SetupRoutesMachine(r *gin.Engine) {
	dbInstance := adapters.NewMySQLMachine()

	rabbitMQInstance, err := core.NewRabbitMQ()
	if err != nil {
		log.Fatalf("No se pudo conectar a RabbitMQ: %v", err)
	}

	var publisher repository.MachineStatusPublisher = adapters.NewRabbitMQProducer(rabbitMQInstance)

	publishMachineStatusService := machineusecases.NewPublishMachineStatusService(publisher)

	listMachineController := machineControllers.NewListMachineController(*machineusecases.NewListMachine(dbInstance))
	createMachineController := machineControllers.NewCreateMachineController(*machineusecases.NewCreateMachine(dbInstance))
	getMachineById := machineControllers.NewMachineByIdController(*machineusecases.NewMachineById(dbInstance))
	getStatusMachine := machineControllers.NewStatusMachine(machineusecases.NewMachineStatus(dbInstance))
	updateMachine := machineControllers.NewUpdateMachineController(*machineusecases.NewUpdateMachine(dbInstance, publishMachineStatusService))
	deleteMachine := machineControllers.NewDeleteMachine(*machineusecases.NewDeleteMachine(dbInstance))

	r.GET("/machines", listMachineController.Execute)
	r.POST("/machines", createMachineController.Execute)
	r.GET("/machines/:id", getMachineById.Execute)
	r.GET("/machines/status/:id", getStatusMachine.Execute)
	r.PUT("/machines/:id", updateMachine.Execute)
	r.DELETE("/machines/:id", deleteMachine.Execute)
}