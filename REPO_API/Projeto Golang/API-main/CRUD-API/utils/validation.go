package utils

import (
	"regexp"
)

var (
	cpfCnpjRegex = regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$|^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
)

func IsValidCPF(cpf string) bool {
	return cpfCnpjRegex.MatchString(cpf)
}

func IsValidCNPJ(cnpj string) bool {
	return cpfCnpjRegex.MatchString(cnpj)
}

func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return emailRegex.MatchString(email)
}

func IsValidBankAccount(agencia, conta, banco string) bool {
	agenciaRegex := regexp.MustCompile(`^\d{4}$`)
	contaRegex := regexp.MustCompile(`^\d{1,12}$`)
	bancoRegex := regexp.MustCompile(`^\d{3}$`)
	return agenciaRegex.MatchString(agencia) && contaRegex.MatchString(conta) && bancoRegex.MatchString(banco)
}

func IsValidPix(pix string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	phoneRegex := regexp.MustCompile(`^\+?[1-9]\d{1,14}$`)
	cpfCnpjRegex := regexp.MustCompile(`^\d{3}\.\d{3}\.\d{3}-\d{2}$|^\d{2}\.\d{3}\.\d{3}/\d{4}-\d{2}$`)
	return emailRegex.MatchString(pix) || phoneRegex.MatchString(pix) || cpfCnpjRegex.MatchString(pix)
}

func AllValidationsPass(cpfCnpj, email, agencia, conta, banco, pix string) bool {
	return IsValidCPF(cpfCnpj) || IsValidCNPJ(cpfCnpj) && IsValidEmail(email) && IsValidBankAccount(agencia, conta, banco) && IsValidPix(pix)
}
