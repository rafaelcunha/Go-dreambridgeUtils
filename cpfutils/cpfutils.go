package cpfutils

import (
	"math/rand"
	"strconv"
)

// NumDigitosCPF - Quantidade de digitos de um CPF
const NumDigitosCPF = 11

// CPF - Estrutura para armazenar um CPF
type CPF struct {
	CpfDigito [NumDigitosCPF]uint
}

// ValidaCPF - retorna verdadeiro se o númeor do CPF é válido
func ValidaCPF(cpf *CPF) bool {

	/*Obtendo o primeiro digito verificador:

	  Os 9 primeiros algarismos são multiplicados pela sequência 10, 9, 8, 7, 6, 5, 4, 3, 2
	  (o primeiro por 10, o segundo por 9, e assim por diante);
	  Em seguida, calcula-se o resto “r1″ da divisão da soma dos resultados das multiplicações por 11,
	  e se o resto for zero ou 1, digito é zero, caso contrário digito = (11-r1) */
	var temp uint
	temp = 0
	for i := 0; i < 9; i++ {
		temp += (cpf.CpfDigito[i] * (10 - uint(i)))
	}

	temp %= 11

	var digito1 uint
	if temp >= 2 {
		digito1 = NumDigitosCPF - temp
	} else {
		digito1 = 0
	}

	/*Obtendo o segundo digito verificador:

	  O dígito2 é calculado pela mesma regra, porém inclui-se o primeiro digito verificador ao final
	  da sequencia. Os 10 primeiros algarismos são multiplicados pela sequencia 11, 10, 9, ... etc...
	  (o primeiro por 11, o segundo por 10, e assim por diante);
	  procedendo da mesma maneira do primeiro digito*/

	temp = 0
	for i := 0; i < 10; i++ {
		temp += (cpf.CpfDigito[i] * (NumDigitosCPF - uint(i)))
	}

	temp %= NumDigitosCPF

	var digito2 uint

	if temp >= 2 {
		digito2 = NumDigitosCPF - temp
	} else {
		digito2 = 0
	}

	/* Se os digitos verificadores obtidos forem iguais aos informados pelo usuário,
	   então o CPF é válido */

	if digito1 == cpf.CpfDigito[9] && digito2 == cpf.CpfDigito[10] {
		return true
	}
	return false
}

// GeradorCPF - Retorna um CPF válido gerado aleatoriamente
func GeradorCPF() *CPF {
	//int *cpf, pesos[11], vetSoma[11], soma, resto, digito,  i;

	// Aloca memória para o cpf
	cpf := new(CPF)

	// Gera 9 números aleatórios
	for i := 0; i < 9; i++ {
		cpf.CpfDigito[i] = uint(rand.Intn(9))
	}

	// Cálculo do primeiro dígito verificador

	// Gera os 9 pesos
	var pesos [11]uint
	for i := 0; i < 9; i++ {
		pesos[i] = 10 - uint(i)
	}

	// Multiplica os valores de cada coluna
	var vetSoma [11]uint
	for i := 0; i < 9; i++ {
		vetSoma[i] = cpf.CpfDigito[i] * pesos[i]
	}

	// Calcula o somatório dos resultados
	var soma uint
	soma = 0
	for i := 0; i < 9; i++ {
		soma += vetSoma[i]
	}

	// Realiza-se a divisão inteira do resultado por 11
	resto := soma % NumDigitosCPF

	// Verifica o resto da divisão
	var digito uint
	if resto < 2 {
		digito = 0
	} else {
		digito = 11 - resto
	}

	// Adiciona o 1º dígito verificador ao cpf
	cpf.CpfDigito[9] = digito

	// Cálculo do segundo dígito verificador

	// Gera os 10 pesos
	for i := 0; i < 10; i++ {
		pesos[i] = NumDigitosCPF - uint(i)
	}

	// Multiplica os valores de cada coluna
	for i := 0; i < 10; i++ {
		vetSoma[i] = cpf.CpfDigito[i] * pesos[i]
	}

	// Calcula o somatório dos resultados
	soma = 0
	for i := 0; i < 10; i++ {
		soma += vetSoma[i]
	}

	// Realiza-se a divisão inteira do resultado por 11
	resto = soma % NumDigitosCPF

	// Verifica o resto da divisão
	if resto < 2 {
		digito = 0
	} else {
		digito = NumDigitosCPF - resto
	}

	// Adiciona o 2º dígito verificador ao cpf
	cpf.CpfDigito[10] = digito

	return cpf
}

// CPFToString - Converte o CPF numérico em string formatada XXX.XXX.XXX.-XX
func CPFToString(cpf *CPF) string {
	var cpfString string

	cpfString = ""
	for i := 0; i < NumDigitosCPF; i++ {
		cpfString += strconv.Itoa(int(cpf.CpfDigito[i]))
	}

	return cpfString
}

// StringToCPF - Recebe um CPF, transforma em números e valida
func StringToCPF(cpfStr string) *CPF {
	length := len(cpfStr)
	if length > NumDigitosCPF {
		// Nesse caso o cpf já está errado
		return nil
	}

	cpf := new(CPF)
	offset := NumDigitosCPF - length
	for i := 0; i < length; i++ {
		if cpfStr[i] < '0' || cpfStr[i] > '9' {
			// Existe um caractere estranho na string
			return nil
		}

		cpf.CpfDigito[(offset + i)] = uint(cpfStr[i]) - uint('0')
	}

	if ValidaCPF(cpf) {
		return cpf
	}
	return nil
}
