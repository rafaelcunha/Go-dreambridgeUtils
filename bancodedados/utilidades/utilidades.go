package utilidades

import (
	"log"
	"time"
)

// DateTimeToGoTime - Converte o retorno da consulta em um campo DATETIME para o Time do go,
// adicionando a informação da time zone
func DateTimeToGoTime(datetime string) (*time.Time, error) {
	t := time.Now()
	nz, _ := t.Zone()
	datetime += " " + nz

	layout := "2006-01-02 15:04:05 -07"
	goTime, err := time.Parse(layout, datetime)

	if err != nil {
		log.Println("utilidades.DateTimeToGoTime - Falha ao convertar string para data.")
		return nil, err
	}

	return &goTime, nil
}
