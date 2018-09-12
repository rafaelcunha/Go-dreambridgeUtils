package fileconfig

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

// GetURLConcexao - Retorna a URL de conexão
func (dadosConexao *DadosArquivoConexao) GetURLConcexao() string {
	//formato da URL => user:password@tcp(127.0.0.1:3306)/database

	if dadosConexao.URL == "localhost" || dadosConexao.URL == "127.0.0.1" {
		return dadosConexao.Usuario + ":" + dadosConexao.Senha + "@tcp(:" + dadosConexao.Porta + ")/" + dadosConexao.NomeDB
	}

	return dadosConexao.Usuario + ":" + dadosConexao.Senha + "@tcp(" + dadosConexao.URL + ":" + dadosConexao.Porta + ")/" + dadosConexao.NomeDB
}

// CarregaArquivoConfig - Carrega as informações de conexão do arquivo.
func (dadosConexao *DadosArquivoConexao) CarregaArquivoConfig(nomeArquivo string) error {

	config, err := ioutil.ReadFile(nomeArquivo)

	if err != nil {
		log.Println("fileconfigdb.DadosArquivoConexao.CarregaArquivo - Erro ao ler arquivo de configuração: " + nomeArquivo)
		return err
	}

	err = json.Unmarshal(config, dadosConexao)

	if err != nil {
		log.Println("fileconfigdb.DadosArquivoConexao.CarregaArquivo - Erro ao ler informações do arquivo de configuração.")
		return err
	}

	return nil
}
