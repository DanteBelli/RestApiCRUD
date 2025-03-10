package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestDataBaseConnection(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Fatalf("Error cargando el archivo de conexion a la Base de Datos")
	}
	dsm := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)
	db, err := gorm.Open(mysql.Open(dsm), &gorm.Config{})
	if err != nil {
		t.Fatalf("No se puede conectar a la base: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		t.Fatalf("No se pudo obtener la data: %v", err)
	}
	err = sqlDB.Ping()
	if err != nil {
		t.Fatalf("Fallo la conexion a la BD: %v", err)
	} else {
		t.Log("Conexion a la base Exitosa")
	}
}
