package equipment

import (
	"gym-system/src/inventory/Equipments/application/useCases"
	equipamentControllers "gym-system/src/inventory/Equipments/infraestructure/controllers"
	"gym-system/src/inventory/Equipments/infraestructure/database"

	"github.com/gin-gonic/gin"
)

func SetupRoutesEquipament(r *gin.Engine){

	dbInstance := equipment.NewMySQLEquipament()

	listEquipamentController := equipamentControllers.NewListEquipmentController(*equipamentusecases.NewListEquipment(dbInstance))
	createEquipamentController := equipamentControllers.NewCreateEquipamentController(*equipamentusecases.NewCreateEquipament(dbInstance))
	getEquipmentById := equipamentControllers.NewEquipmentByIdController(*equipamentusecases.NewEquipmentById(dbInstance))
	getEquipmentCondition := equipamentControllers.NewEquipmentCondition(equipamentusecases.NewEquipmentByCondition(dbInstance))
	updateEquipment := equipamentControllers.NewUpdateEquipmentController(*equipamentusecases.NewUpdateEquipment(dbInstance))
	deleteEquipment := equipamentControllers.NewDeleteEquipment(*equipamentusecases.NewDeleteEquipment(dbInstance))

	r.GET("/equipments",listEquipamentController.Execute)
	r.POST("/equipments",createEquipamentController.Execute)
	r.GET("/equipments/:id",getEquipmentById.Execute)
	r.GET("equipments/condition/:condition",getEquipmentCondition.Execute)
	r.PUT("/equipments/:id",updateEquipment.Execute)
	r.DELETE("/equipments/:id",deleteEquipment.Execute)

}