package usecases

// CryptoHandler eh a interface do cryptohandler e contem todas as funcoes e retornos.
type CryptoHandler interface {
	EncryptString(string) (string, error)
	DecryptString(string) (string, error)
}
