package mariadb

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" //Necessário para abertura da conexão com o DB
)

// MaxConnections - Número máximo de conexões abertas no DB
const MaxConnections int = 2

// IdleTimeMinutes - Tempo em minutos para manter conexões aberta no estado IDLE
const IdleTimeMinutes int = 5

// ConfiguracaoServidor - Estrutura contendo a configuração para acesso ao servidor de DB
type ConfiguracaoServidor struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	URL     string `json:"url"`
	Porta   string `json:"porta"`
	NomeDB  string `json:"dbnome"`
	DB      *sql.DB
}

// CarregaConfiguracoes - Le os arquivos de configuração de conexão com a DB
func carregaConfiguracoes(nomeArquivo string) (*ConfiguracaoServidor, error) {

	var configuracao ConfiguracaoServidor

	config, err := ioutil.ReadFile(nomeArquivo + ".json")

	if err != nil {
		log.Fatalf("Erro ao ler arquivo de configuração: %v", err)
		return &configuracao, err
	}
	//fmt.Println("Arquivo de configuração lido: " + string(config))

	err = json.Unmarshal(config, &configuracao)

	if err != nil {
		log.Fatalf("Erro ao ler informações do arquivo de configuração: %v", err)
		return &configuracao, err
	}

	log.Println("Arquivo de configuração carregado com sucesso: ", configuracao)

	return &configuracao, err
}

func montaURLConexao(configuracao *ConfiguracaoServidor) string {
	//formato da URL => user:password@tcp(127.0.0.1:3306)/database

	if configuracao.URL == "localhost" || configuracao.URL == "127.0.0.1" {
		return configuracao.Usuario + ":" + configuracao.Senha + "@tcp(:" + configuracao.Porta + ")/" + configuracao.NomeDB
	}

	return configuracao.Usuario + ":" + configuracao.Senha + "@tcp(" + configuracao.URL + ":" + configuracao.Porta + ")/" + configuracao.NomeDB

}

// InicializaDB - Realiza uma conexão com a base de dados e retorna a conexão.
func inicializaDB(config *ConfiguracaoServidor) error {
	var err error

	config.DB, err = sql.Open("mysql", montaURLConexao(config))

	if err != nil {
		log.Fatalf("Erro ao abrir conexao com DB: %v", err)
		return err
	}

	config.DB.SetMaxIdleConns(MaxConnections)

	config.DB.SetMaxOpenConns(MaxConnections)

	var tempoDuracao = time.Minute * time.Duration(IdleTimeMinutes)
	config.DB.SetConnMaxLifetime(tempoDuracao)

	if testaConexaoDB(config) == false {
		err = errors.New("falha no teste de conexão com o DB")
		//log.Fatalf("Falha testar conexao com DB.")
		log.Printf("Falha testar conexao com DB.")
		return err
	}

	//log.Println("Conexão realizada com sucesso.")

	return nil
}

// TestaConexaoDB - Envia um PING para o DB verificar se a conexão está ok.
func testaConexaoDB(config *ConfiguracaoServidor) bool {
	err := config.DB.Ping()

	if err != nil {
		log.Printf("Erro testar conexao com DB: %v", err)
		//log.Fatalf("Erro testar conexao com DB: %v", err)
		return false
	}

	return true
}

// ListaTabelasDB - Testa a conexão com o banco de dados.
func listaTabelasDB(db *sql.DB) error {
	rows, err := db.Query("SHOW TABLES")

	if err != nil {
		log.Fatalf("Erro: %v", err)
		return err
	}

	defer rows.Close()
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			log.Fatal(err)
			return err
		}
		fmt.Println(name)
	}
	if err := rows.Err(); err != nil {
		log.Fatal(err)
		return err
	}

	return nil
}
