package prisma

import (
	"log"
	"paynau-backend/prisma/db"

	"github.com/joho/godotenv"
)

var Prisma *db.PrismaClient

func InitPrisma() {
	// Cargar variables de entorno
	if err := godotenv.Load(); err != nil {
		log.Println("No se pudo cargar el archivo .env")
	}

	var err error
	Prisma = db.NewClient()
	if err = Prisma.Prisma.Connect(); err != nil {
		log.Fatalf("Error connecting to Prisma: %v", err)
	}

	log.Println("Prisma conectado exitosamente")
}

func ClosePrisma() {
	if Prisma != nil {
		_ = Prisma.Prisma.Disconnect()
	}
}

func GetPrisma() *db.PrismaClient {
	if Prisma == nil {
		log.Fatal("Prisma no ha sido inicializado. Asegúrate de llamar InitPrisma() primero")
	}
	return Prisma
}
