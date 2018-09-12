package snipetconfig

import (
	"fmt"

	"github.com/mercadolibre/go-meli-toolkit/gomelipass"
)

// GetURLConexao - Retorna a URL de conexão
func (dadosConexao *DadosSnipetConexao) GetURLConexao() string {
	//formato da URL => user:password@tcp(127.0.0.1:3306)/database

	return fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", dadosConexao.Usuario, dadosConexao.Senha, dadosConexao.URL, dadosConexao.NomeDB)
}

// CarregaConfig - Carrega as informações de conexão do snipet.
func (dadosConexao *DadosSnipetConexao) CarregaConfig() error {

	dadosConexao.NomeDB = "prefdata"
	dadosConexao.Usuario = "prefdata_RPROD"
	dadosConexao.Senha = gomelipass.GetEnv("DB_MYSQL_INTEGRACOES01_PREFDATA_PREFDATA_RPROD")
	dadosConexao.URL = gomelipass.GetEnv("DB_MYSQL_INTEGRACOES01_PREFDATA_PREFDATA_ENDPOINT")

	return nil
}
