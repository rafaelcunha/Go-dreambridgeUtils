package fileconfig

// DadosArquivoConexao - Estrutura do arquivo com as configurações do servidor
type DadosArquivoConexao struct {
	Usuario string `json:"usuario"`
	Senha   string `json:"senha"`
	URL     string `json:"url"`
	Porta   string `json:"porta"`
	NomeDB  string `json:"dbnome"`
}
