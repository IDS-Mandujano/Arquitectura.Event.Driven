package adapters

import (
	"fmt"
	"gym-system/src/core"
	"log"
)

type MySQLMachine struct {
	conn *core.Conn_MySQL
}

func NewMySQLMachine() *MySQLMachine {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQLMachine{conn: conn}
}

func (mysql *MySQLMachine) Save(cname string, ctype string, cstatus string) {
	query := "INSERT INTO machines (cname, ctype, cstatus) VALUES (?, ?, ?)"

	result, err := mysql.conn.ExecutePreparedQuery(query, cname, ctype, cstatus)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Equipo guardado: %d", rowsAffected)
	}
}

func (mysql *MySQLMachine) GetAll() ([]map[string]interface{}, error) {
	query := "SELECT * FROM machines"
	rows := mysql.conn.FetchRows(query)
	defer rows.Close()

	var machines []map[string]interface{}
	for rows.Next() {
		var id int
		var cname, ctype, cstatus string
		if err := rows.Scan(&id, &cname, &ctype, &cstatus); err != nil {
			return nil, err
		}
		machine := map[string]interface{}{
			"id":        id,
			"cname":      cname,
			"ctype":  ctype,
			"cstatus": cstatus,
		}
		machines = append(machines, machine)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

func (mysql *MySQLMachine) GetById(id int) ([]map[string]interface{}, error) {
	query := "SELECT * FROM machines WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	if rows == nil {
		return nil, fmt.Errorf("no se pudo ejecutar la consulta o no hay resultados")
	}
	defer rows.Close()

	var machines []map[string]interface{}
	for rows.Next() {
		var id int
		var cname, ctype, cstatus string
		if err := rows.Scan(&id, &cname, &ctype, &cstatus); err != nil {
			return nil, err
		}
		machine := map[string]interface{}{
			"id":	id,
			"name":	cname,
			"ctype":  ctype,
			"cstatus": cstatus,
		}
		machines = append(machines, machine)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return machines, nil
}

func (mysql *MySQLMachine) GetStatus(id int) (string, error) {
	query := "SELECT cstatus FROM machines WHERE id = ?"
	rows := mysql.conn.FetchRows(query, id)
	if rows == nil {
		return "", fmt.Errorf("no se pudo ejecutar la consulta o no hay resultados")
	}
	defer rows.Close()

	var status string
	if rows.Next() {
		if err := rows.Scan(&status); err != nil {
			return "", err
		}
	}

	if err := rows.Err(); err != nil {
		return "", err
	}

	return status, nil
}

func (mysql *MySQLMachine) Update(id int, cname string, ctype string, cstatus string) {
	query := "UPDATE machines SET cname = ?, ctype = ?, cstatus = ? WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, cname, ctype, cstatus, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Equipo actualizado: %d", rowsAffected)
	}
}

func (mysql *MySQLMachine) Delete(id int) {
	query := "DELETE FROM machines WHERE id = ?"
	result, err := mysql.conn.ExecutePreparedQuery(query, id)
	if err != nil {
		log.Fatalf("Error al ejecutar la consulta: %v", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 1 {
		log.Printf("[MySQL] - Equipo eliminado: %d", rowsAffected)
	}
}