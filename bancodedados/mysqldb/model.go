package mysqldb

import (
	"database/sql"
	"sync"
)

// InterfaceDadosConexao - Gerencia a conexão com um banco de dados MySQL/MariaDB
type InterfaceDadosConexao interface {
	GetURLConexao() string
}

// InterfaceConexaoMySQLDB - Gerencia a conexão com um banco de dados MySQL/MariaDB
type InterfaceConexaoMySQLDB interface {
	InicializaConexao() error
	FinalizaConexao() error
	TestaConexao() error
	ExecutaSelect(query string, args ...interface{}) (*sql.Rows, error)
	ExecutaUpdateInsertDelete(query string, args ...interface{}) (*sql.Result, error)
}

// ConexaoMySQLDB - Estrutura de dados contendo informação sobre a conexão con o DB
type ConexaoMySQLDB struct {
	DadosConexao InterfaceDadosConexao
	DB           *sql.DB
	mutex        sync.Mutex
}
