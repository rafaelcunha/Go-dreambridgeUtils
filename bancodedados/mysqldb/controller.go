package mysqldb

import (
	"database/sql"
	"errors"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //Necessário para abertura da conexão com o DB
)

// MaxConnections - Número máximo de conexões abertas no DB
const MaxConnections int = 1

// IdleTimeMinutes - Tempo em minutos para manter conexões aberta no estado IDLE
const IdleTimeMinutes int = 1

// TestaConexao - Envia um PING para o DB verificar se a conexão está ok.
func (conexaoDB *ConexaoMySQLDB) TestaConexao() error {

	if conexaoDB.DB == nil {
		mensagem := "mysqldb.ConexaoMySQLDB.TestaConexao - Ponteiro para conexão com o DB nulo"
		log.Println(mensagem)
		return errors.New(mensagem)
	}

	err := conexaoDB.DB.Ping()

	if err != nil {
		log.Println("mysqldb.ConexaoMySQLDB.TestaConexao - Erro testar conexao com DB.")
		return err
	}

	return nil
}

// InicializaConexao - Inicializa a conexão com o banco de dados.
func (conexaoDB *ConexaoMySQLDB) InicializaConexao() error {
	conexaoDB.mutex.Lock()
	defer conexaoDB.mutex.Unlock()

	var err error

	conexaoDB.DB, err = sql.Open("mysql", conexaoDB.URLConexao)

	if err != nil {
		log.Fatalf("mysqldb.ConexaoMySQLDB.InicializaConexao - Erro ao abrir conexao com DB.")
		return err
	}

	conexaoDB.DB.SetMaxIdleConns(MaxConnections) // Número máximo de conexões abertas sem uso
	conexaoDB.DB.SetMaxOpenConns(MaxConnections) // Número máximo de conexões abertas simultaneamente

	var tempoDuracao = time.Minute * time.Duration(IdleTimeMinutes)
	conexaoDB.DB.SetConnMaxLifetime(tempoDuracao) // Tempo máximo que uma conexão pode ser reutilizada. 0 = para sempre

	err = conexaoDB.TestaConexao()

	if err != nil {
		conexaoDB.DB = nil
		log.Printf("mysqldb.ConexaoMySQLDB.InicializaConexao  - Falha ao testar conexao com DB.")
		return err
	}

	return nil
}

// FinalizaConexao - Finaloiza a conexão com o banco de dados.
func (conexaoDB *ConexaoMySQLDB) FinalizaConexao() error {
	if conexaoDB.DB == nil {
		mensagem := "mysqldb.ConexaoMySQLDB.FinalizaConexao - Ponteiro para conexão com o DB nulo"
		log.Println(mensagem)
		return errors.New(mensagem)
	}

	conexaoDB.mutex.Lock()
	defer conexaoDB.mutex.Unlock()

	err := conexaoDB.DB.Close()
	conexaoDB.DB = nil

	if err != nil {
		log.Println("mysqldb.ConexaoMySQLDB.FinalizaConexao - Erro ao finalizar banco da dados.")
		return err
	}

	return nil
}

// ExecutaUpdateInsertDelete - Executa uma operação de delete, update ou insert no banco de dados
func (conexaoDB *ConexaoMySQLDB) ExecutaUpdateInsertDelete(query string, args ...interface{}) (*sql.Result, error) {
	if conexaoDB.DB == nil {
		mensagem := "mysqldb.ConexaoMySQLDB.ExecutaUpdateInsertDelete - Ponteiro para conexão com o DB nulo"
		log.Println(mensagem)
		return nil, errors.New(mensagem)
	}

	conexaoDB.mutex.Lock()
	defer conexaoDB.mutex.Unlock()

	// Prepara a query
	stmt, err := conexaoDB.DB.Prepare(query)

	if err != nil {
		log.Println("mysqldb.ConexaoMySQLDB.ExecutaUpdateInsertDelete - Erro ao preparar a query.")
		return nil, err
	}
	defer stmt.Close()

	// Executa a query no banco de dados
	result, err := stmt.Exec(args...)

	if err != nil {
		log.Println("mysqldb.ConexaoMySQLDB.ExecutaUpdateInsertDelete - Erro ao executar a query: " + query)
		return nil, err
	}

	return &result, nil
}

// ExecutaSelect - Executa uma operação de select no banco de dados
func (conexaoDB *ConexaoMySQLDB) ExecutaSelect(query string, args ...interface{}) (*sql.Rows, error) {
	if conexaoDB.DB == nil {
		mensagem := "mysqldb.ConexaoMySQLDB.ExecutaSelect - Ponteiro para conexão com o DB nulo"
		log.Println(mensagem)
		return nil, errors.New(mensagem)
	}

	conexaoDB.mutex.Lock()
	defer conexaoDB.mutex.Unlock()

	// Prepara a query
	stmt, err := conexaoDB.DB.Prepare(query)

	if err != nil {
		log.Println("mysqldb.ConexaoMySQLDB.ExecutaSelect - Erro ao preparar a query: " + query)
		return nil, err
	}

	defer stmt.Close()

	// Executa a query no banco de dados
	rows, err := stmt.Query(args...)
	if err != nil {
		log.Println("mysqldb.ConexaoMySQLDB.ExecutaSelect - Erro ao executar a query: " + query)
		return nil, err
	}

	if rows == nil {
		mensagem := "mysqldb.ConexaoMySQLDB.ExecutaSelect - Resultado nulo."
		log.Println(mensagem)
		return nil, errors.New(mensagem)
	}

	return rows, nil
}
