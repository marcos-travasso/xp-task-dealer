package openai_dealer

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"xp-task-dealer/core/models"
)

func TestOpenAIDealer_GetDeveloperForTask(t *testing.T) {
	task := models.NewTask("Desenvolver tela de login", "Como um usuário, gostaria de poder acessar a aplicação através de uma tela de login, onde possa inserir meu email e senha, para que eu possa ter acesso a plataforma.")

	developer1 := models.NewDeveloper("Alice", "Alice é uma desenvolvedora frontend experiente em criar interfaces elegantes e funcionais, com habilidades em HTML, CSS e JavaScript.")
	developer2 := models.NewDeveloper("Bob", "Bob é um desenvolvedor backend experiente em sistemas escaláveis, com habilidades em diversas linguagens e bancos de dados.")

	developers := []models.Developer{developer1, developer2}

	d := Init()

	chosenDev, err := d.GetDeveloperForTask(task, developers)
	assert.NoError(t, err)
	assert.Equal(t, developer1.ID, chosenDev.ID)
}

func TestOpenAIDealer_GetTaskForDeveloper(t *testing.T) {
	developer := models.NewDeveloper("Bob", "Bob é um desenvolvedor backend experiente em sistemas escaláveis, com habilidades em diversas linguagens e bancos de dados.")

	task1 := models.NewTask("Enviar email de recuperação de senha", "Como um usuário, gostaria de receber um email contendo as instruções para recuperar a senha.")
	task2 := models.NewTask("Desenvolver tela de login", "Como um usuário, gostaria de poder acessar a aplicação através de uma tela de login, onde possa inserir meu email e senha, para que eu possa ter acesso a plataforma.")

	tasks := []models.Task{task1, task2}

	d := Init()

	chosenTask, err := d.GetTaskForDeveloper(developer, tasks)
	assert.NoError(t, err)
	assert.Equal(t, task1.ID, chosenTask.ID)
}

func TestOpenAIDealer_GetPairForDeveloper(t *testing.T) {
	task := models.NewTask("Desenvolver tela de login", "Como um usuário, gostaria de poder acessar a aplicação através de uma tela de login, onde possa inserir meu email e senha, para que eu possa ter acesso a plataforma.")

	mainDeveloper := models.NewDeveloper("Alice", "Alice é uma desenvolvedora frontend experiente em criar interfaces elegantes e funcionais, com habilidades em HTML, CSS e JavaScript.")

	developer1 := models.NewDeveloper("Bob", "Bob é um desenvolvedor backend experiente em sistemas escaláveis, com habilidades em diversas linguagens e bancos de dados.")
	developer2 := models.NewDeveloper("Charlie", "Charlie é um estagiário em frontend que fez algumas páginas web e uma aplicação simples de CRUD.")

	developers := []models.Developer{developer1, developer2}

	d := Init()

	chosenDev, err := d.GetPairForDeveloper(mainDeveloper, task, developers)
	assert.NoError(t, err)
	assert.Equal(t, developer2.ID, chosenDev.ID)
}
