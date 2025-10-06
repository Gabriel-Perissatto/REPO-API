package models

import (
	"fmt"
	"time"

	"github.com/gustavolubacheski/API/CRUD-API/utils"
)

type User struct {
	Id            int    `json:"id"`
	Nome          string `json:"name"`
	Email         string `json:"email"`
	CPF_CNPJ      string `json:"cpf_cnpj"`
	ContaBancaria `json:"conta_bancaria"`
	CreatedAt     string `json:"created_at"`
}

type ContaBancaria struct {
	Agencia string `json:"agencia"`
	Conta   string `json:"conta"`
	Banco   string `json:"banco"`
	Pix     string `json:"pix"`
}

func NewUser(u User) (bool, error) {
	con := Connect()
	defer con.Close()

	if !utils.AllValidationsPass(u.CPF_CNPJ, u.Email, u.ContaBancaria.Agencia, u.ContaBancaria.Conta, u.ContaBancaria.Banco, u.ContaBancaria.Pix) {
		return false, fmt.Errorf("validation failed")
	}

	u.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	sql := "INSERT INTO users (nome, email, cpf_cnpj, agencia, conta, banco, pix, created_at) VALUES (?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := con.Prepare(sql)
	if err != nil {
		return false, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(u.Nome, u.Email, u.CPF_CNPJ, u.ContaBancaria.Agencia, u.ContaBancaria.Conta, u.ContaBancaria.Banco, u.ContaBancaria.Pix, u.CreatedAt)
	if err != nil {
		return false, err
	}

	return true, nil
}

func GetUsers() ([]User, error) {
	con := Connect()
	sql := "SELECT * FROM users"
	rows, err := con.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	defer con.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.Id, &user.Nome, &user.Email, &user.CPF_CNPJ, &user.ContaBancaria.Agencia, &user.ContaBancaria.Conta, &user.ContaBancaria.Banco, &user.ContaBancaria.Pix, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func GetUserByCPF(cpf string) (*User, error) {
	con := Connect()
	defer con.Close()
	sql := "SELECT * FROM users WHERE cpf_cnpj = ?"
	row := con.QueryRow(sql, cpf)
	var user User
	err := row.Scan(&user.Id, &user.Nome, &user.Email, &user.CPF_CNPJ, &user.ContaBancaria.Agencia, &user.ContaBancaria.Conta, &user.ContaBancaria.Banco, &user.ContaBancaria.Pix, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func DeleteUserByCPF(cpf string) (int64, error) {
	con := Connect()
	defer con.Close()
	sql := "DELETE FROM users WHERE cpf_cnpj = ?"
	rs, err := con.Exec(sql, cpf)
	if err != nil {
		return 0, err
	}
	return rs.RowsAffected()
}

func UpdateUser(user User) (int64, error) {
	con := Connect()
	defer con.Close()
	sqlStatement := "UPDATE users SET nome = ?, email = ?, agencia = ?, conta = ?, banco = ?, pix = ? WHERE cpf_cnpj = ?"

	if !utils.AllValidationsPass(user.CPF_CNPJ, user.Email, user.ContaBancaria.Agencia, user.ContaBancaria.Conta, user.ContaBancaria.Banco, user.ContaBancaria.Pix) {
		return 0, fmt.Errorf("validation failed")
	}

	stmt, err := con.Prepare(sqlStatement)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	rs, err := stmt.Exec(user.Nome, user.Email, user.ContaBancaria.Agencia, user.ContaBancaria.Conta, user.ContaBancaria.Banco, user.ContaBancaria.Pix, user.CPF_CNPJ)
	if err != nil {
		return 0, err
	}

	rows, err := rs.RowsAffected()
	if err != nil {
		return 0, err
	}

	return rows, nil
}
