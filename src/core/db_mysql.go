package core

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type ConnMySQL struct {
	DB  *sql.DB
	Err string
}

// GetDBPool inicializa y devuelve una conexión a la base de datos
func GetDBPool() (*ConnMySQL, error) {
	err := godotenv.Load()
	if err != nil {
		log.Println("Advertencia: No se pudo cargar el archivo .env, usando variables de entorno del sistema.")
	}

	dbHost := os.Getenv("DB_HOST")     // localhost
	dbPort := os.Getenv("DB_PORT")     // 3306
	dbUser := os.Getenv("DB_USER")     // root
	dbPass := os.Getenv("DB_PASS")     // (vacío, si es el caso)
	dbSchema := os.Getenv("DB_SCHEMA") // nombre de la base de datos

	// Verificar que se hayan definido las variables necesarias
	if dbHost == "" || dbPort == "" || dbUser == "" || dbSchema == "" {
		return nil, fmt.Errorf("faltan variables de entorno: DB_HOST, DB_PORT, DB_USER y/o DB_SCHEMA")
	}

	// Construir el DSN usando el puerto desde la variable de entorno
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", dbUser, dbPass, dbHost, dbPort, dbSchema)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, fmt.Errorf("error al abrir la base de datos: %v", err)
	}

	// Configuración de conexión
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(60 * time.Minute)

	// Verificar la conexión
	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("error al verificar la conexión a la base de datos: %v", err)
	}

	log.Println("Conexión a la base de datos establecida correctamente.")
	return &ConnMySQL{DB: db}, nil
}

// ExecutePreparedQuery ejecuta una consulta preparada con parámetros
func (conn *ConnMySQL) ExecutePreparedQuery(query string, values ...interface{}) (sql.Result, error) {
	if conn.DB == nil {
		return nil, fmt.Errorf("la conexión a la base de datos no está inicializada")
	}

	stmt, err := conn.DB.Prepare(query)
	if err != nil {
		return nil, fmt.Errorf("error al preparar la consulta: %w", err)
	}
	defer stmt.Close()

	result, err := stmt.Exec(values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta preparada: %w", err)
	}

	return result, nil
}

// FetchRows ejecuta una consulta SELECT y devuelve múltiples filas
func (conn *ConnMySQL) FetchRows(query string, values ...interface{}) (*sql.Rows, error) {
	if conn.DB == nil {
		return nil, fmt.Errorf("la conexión a la base de datos no está inicializada")
	}

	rows, err := conn.DB.Query(query, values...)
	if err != nil {
		return nil, fmt.Errorf("error al ejecutar la consulta SELECT: %w", err)
	}
	return rows, nil
}

// FetchRow ejecuta una consulta SELECT y devuelve una sola fila
func (conn *ConnMySQL) FetchRow(query string, values ...interface{}) *sql.Row {
	if conn.DB == nil {
		log.Println("Advertencia: la conexión a la base de datos no está inicializada")
		return nil
	}
	return conn.DB.QueryRow(query, values...)
}

// Close cierra la conexión a la base de datos
func (conn *ConnMySQL) Close() {
	if conn.DB != nil {
		conn.DB.Close()
		log.Println("Conexión a la base de datos cerrada correctamente.")
	}
}
