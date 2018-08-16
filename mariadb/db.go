package mariadb

import (
	"database/sql"
	"fmt"
	"log"
)

// Variaveis Globais
var nomeArquivoConfiguracao = "dbconfig"
var db *ConfiguracaoServidor

// AÇÕES GERAIS

// InicializaDB - realiza os processos para inicialização do DB e o mantém aberto.
func InicializaDB() error {
	var err error

	if db == nil {
		db, err = carregaConfiguracoes(nomeArquivoConfiguracao)

		if err != nil {
			return err
		}
	}

	err = inicializaDB(db)

	if err != nil {
		log.Fatalf("Erro: %v", err)
		return err
	}

	return nil
}

func finalizaDB() {
	if db != nil && db.DB != nil {
		db.DB.Close()
	}
}

// Finalizacao - Chama todas as finalizações necessárias
func Finalizacao() {
	finalizaDB()
}

// ExecutaQuery - Executa uma query e trata os resultados
func ExecutaQuery(query string, args ...interface{}) (*sql.Rows, error) {
	InicializaDB()

	// Prepara a query
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		log.Fatalf("ExecutaQuery - Erro ao preparar a query: " + query)
		log.Fatalf("ExecutaQuery - Erro: %v", err)
		return nil, err
	}
	defer stmt.Close()

	// Executa a query no banco de dados
	rows, err := stmt.Query(args...)

	if err != nil {
		log.Printf("ExecutaQuery - Erro ao executar a query: " + query)
		fmt.Println(err.Error())
		//log.Fatalf("Erro: %v", err)
		return nil, err
	}

	return rows, nil
}
