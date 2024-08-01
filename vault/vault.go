package vault

import (
	"log"

	"github.com/go-resty/resty/v2"
)

type VaultInterface interface {
	GetSecret() (Secret, error)
}

type vault struct {
	vaultClient *resty.Client
	loginURL    string
	secretURL   string
	password    string
}

func NewVault(
	baseURL string,
	loginURL string,
	secretURL string,
	password string,
) VaultInterface {
	return &vault{
		vaultClient: resty.New().
			SetBaseURL(baseURL).
			SetDebug(false),
		loginURL:  loginURL,
		secretURL: secretURL,
		password:  password,
	}
}

func (v *vault) GetSecret() (Secret, error) {
	login := LoginResponse{}
	secret := Secret{}

	_, err := v.vaultClient.R().
		SetBody(
			Password{
				Password: v.password,
			},
		).SetResult(&login).
		Post(v.loginURL)
	if err != nil {
		return secret, err
	}

	_, err = v.vaultClient.R().
		SetHeader(
			xVaultToken,
			login.Auth.ClientToken,
		).
		SetResult(&secret).
		Get(v.secretURL)
	if err != nil {
		log.Println(err)
		return secret, err
	}

	return secret, nil
}
