package modelos

import (
	"api/src/seguranca"
	"errors"
	"github.com/badoux/checkmail"
	"strings"
	"time"
)

// usuario representa um usuario usando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempty"`
	Nome     string    `json:"nome,omitempty"`
	Nick     string    `json:"nick,omitempty"`
	Email    string    `json:"email,omitempty"`
	Senha    string    `json:"senha,omitempty"`
	CriadoEm time.Time `json:"CriadoEm,omitempty"`
}

// Preparar vai chamar os metodos para validar e formatar o usuario recebido
func (usuario *Usuario) Preparar(etapa string) error {
	if err := usuario.validar(etapa); err != nil {
		return err
	}
	if err := usuario.formatar(etapa); err != nil {
		return err
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome eh obrigatorio e nao pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick eh obrigatorio e nao pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O Email eh obrigatorio e nao pode estar em branco")
	}
	if err := checkmail.ValidateFormat(usuario.Email); err != nil {
		return errors.New("O Email inserido eh invalido")
	}
	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("O Senha eh obrigatorio e nao pode estar em branco")
	}
	return nil
}

func (usuario *Usuario) formatar(etapa string) error {
	usuario.Nome = strings.TrimSpace(usuario.Nome)
	usuario.Nick = strings.TrimSpace(usuario.Nick)
	usuario.Email = strings.TrimSpace(usuario.Email)

	if etapa == "cadastro" {
		senhaComHash, err := seguranca.Hash(usuario.Senha)
		if err != nil {
			return err
		}
		usuario.Senha = string(senhaComHash)
	}
	return nil
}
