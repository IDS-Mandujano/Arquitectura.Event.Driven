package main

import (
	"log"
    "github.com/joho/godotenv"
    "github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
		equipamentInfra "gym-system/src/inventory/Equipments/infraestructure/routes"
		machineInfra "gym-system/src/inventory/Machines/infraestructure/routes"
)


func main () {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error al cargar el archivo.env: %v", err)
	}

	r := gin.Default()

	r.Use(cors.New(cors.Config{
        AllowOrigins:     []string{"*"},
        AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
        AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
        AllowCredentials: true,
    }))

	equipamentInfra.SetupRoutesEquipament(r)
	machineInfra.SetupRoutesMachine(r)

	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v",err)
	}
}