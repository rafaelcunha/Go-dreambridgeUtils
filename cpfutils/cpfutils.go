package cpfutils

import (
	"errors"
	"strconv"
	"strings"
)

func ValidaCPF(cpf string) error {
	cpf = strings.Replace(cpf, ".", "", -1)
	cpf = strings.Replace(cpf, "-", "", -1)

	if len(cpf) != 11 {
		return errors.New("CPF inválido")
	}

	var eq bool
	var dig string

	for _, val := range cpf {
		if len(dig) == 0 {
			dig = string(val)
		}
		if string(val) == dig {
			eq = true
			continue
		}
		eq = false
		break
	}

	if eq {
		return errors.New("CPF inválido")
	}

	i := 10
	sum := 0
	for index := 0; index < len(cpf)-2; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}

	prod := sum * 10
	mod := prod % 11
	if mod == 10 {
		mod = 0
	}
	digit1, _ := strconv.Atoi(string(cpf[9]))
	if mod != digit1 {
		return errors.New("CPF inválido")
	}

	i = 11
	sum = 0
	for index := 0; index < len(cpf)-1; index++ {
		pos, _ := strconv.Atoi(string(cpf[index]))
		sum += pos * i
		i--
	}
	prod = sum * 10
	mod = prod % 11
	if mod == 10 {
		mod = 0
	}
	digit2, _ := strconv.Atoi(string(cpf[10]))
	if mod != digit2 {
		return errors.New("CPF inválido")
	}
	return nil
}
