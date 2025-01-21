package repositorios

import (
	"api/src/modelos"
	"database/sql"
	"fmt"
)

type Usuarios struct {
	db *sql.DB
}

// NovoRepositorioDeUsuarios cria um repositorio de usuarios
func NovoRepositorioDeUsuarios(db *sql.DB) *Usuarios {
	return &Usuarios{db}
}

// Criar insere um usuario no banco de dados
func (repositorio Usuarios) Criar(usuario modelos.Usuario) (uint64, error) {
	statement, err := repositorio.db.Prepare(
		"insert into usuarios (nome, nick, email, senha) values (?, ?, ?, ?)",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	resultado, err := statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, usuario.Senha)
	if err != nil {
		return 0, err
	}
	ultimoID, err := resultado.LastInsertId()
	if err != nil {
		return 0, err
	}
	return uint64(ultimoID), nil
}

// Buscar traz todos usuarios que atedem o filtro de nome ou nick
func (repositorio Usuarios) Buscar(nomeOuNick string) ([]modelos.Usuario, error) {
	nomeOuNick = fmt.Sprintf("%%%s%%", nomeOuNick) // %nomeOuNick%

	linhas, err := repositorio.db.Query(
		"select id, nome, nick, email, criadoEm from usuarios where nome LIKE ? or nick LIKE ?",
		nomeOuNick, nomeOuNick,
	)
	if err != nil {
		return nil, err
	}
	defer linhas.Close()

	var usuarios []modelos.Usuario

	for linhas.Next() {
		var usuario modelos.Usuario
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return nil, err
		}
		usuarios = append(usuarios, usuario)
	}
	return usuarios, nil
}

func (repositorio Usuarios) BuscarPorID(ID uint64) (modelos.Usuario, error) {
	linhas, err := repositorio.db.Query(
		"selecet id, nome, nick, email, criadoEm from usuarios where id = ?",
		ID,
	)
	if err != nil {
		return modelos.Usuario{}, err
	}
	defer linhas.Close()

	var usuario modelos.Usuario

	if linhas.Next() {
		if err = linhas.Scan(
			&usuario.ID,
			&usuario.Nome,
			&usuario.Nick,
			&usuario.Email,
			&usuario.CriadoEm,
		); err != nil {
			return modelos.Usuario{}, err
		}
	}
	return usuario, nil
}

// Atualizar altera as informacoes de um usuario
func (repositorio Usuarios) Atualizar(ID uint64, usuario modelos.Usuario) error {
	statement, err := repositorio.db.Prepare(
		"update usuarios set nome = ?, nick = ?, email = ? where id = ?",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(usuario.Nome, usuario.Nick, usuario.Email, ID); err != nil {
		return err
	}
	return nil
}
