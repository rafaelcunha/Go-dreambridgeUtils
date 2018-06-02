package dreambridge_utils

import (
	"strconv"
)

func validaCPF(cpf string) bool {
    var digitos_iguais int = 1;
    var i int;

    if len(cpf) < 11 {
    	return false;
    }

    for i := 0; i < len(cpf) - 1; i++ {
		if (cpf[i] != cpf[i + 1])
        {
            digitos_iguais := 0;
        	break;
        }
	}

	var numeros, digitos string;

	var soma, resultado int;

	if digitos_iguais == 0 {
		numeros := cpf[0:9];
		digitos := cpf[9];
		soma := 0;
		for i := 10; i>1; i-- {
			soma += strconv.Atoi(numeros[10 - i]) * i;
		}

		if soma % 11 < 2 {
			resultado := 0;
		} 
		else {
			resultado := 11 - (soma % 11);
		}

		if resultado != digitos[0]) {
			return false;
		}

		numeros := cpf[0:10];
		soma := 0;

		for i := 11; i > 1; i--{
			soma := soma + strconv.Atoi(numeros[11 - i) * i;
		}  
		if soma % 11 < 2 {
			resultado := 0;
		} 
		else {
			resultado := 11 - (soma % 11);
		}
		if resultado != strconv.Atoi(digitos[1]) {
			return false;
		}
        
        return true;
	}

	return false;
}
