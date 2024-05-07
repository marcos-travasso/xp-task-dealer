package openai_dealer

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	_ "github.com/joho/godotenv/autoload"
	"github.com/sashabaranov/go-openai"
	"os"
	"strings"
	"xp-task-dealer/core/models"
)

var ErrDeveloperNotChosen = errors.New("dealer could not choose developer")
var ErrTaskNotChosen = errors.New("dealer could not choose a task")

type OpenAIDealer struct {
	client *openai.Client
}

func Init() *OpenAIDealer {
	return &OpenAIDealer{
		client: openai.NewClient(os.Getenv("OPENAI_KEY")),
	}
}

func (o *OpenAIDealer) GetDeveloperForTask(task models.Task, developers []models.Developer) (models.Developer, error) {
	var developersPrompt strings.Builder
	developersMap := make(map[string]models.Developer)

	for _, dev := range developers {
		developersPrompt.WriteString(fmt.Sprintf("Nome: %s - Descrição: %s\n", dev.Name, dev.Description))
		developersMap[dev.Name] = dev
	}

	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "Responda qual o desenvolvedor mais indicado para desempenhar a tarefa enviada.\n" +
						"A sua resposta deve ser um objeto JSON com um único campo \"nome\" contendo o nome do desenvolvedor escolhido.\n" +
						"Os desenvolvedores disponíveis são:\n" + developersPrompt.String() +
						"A tarefa é:\n" + task.Title + ": " + task.Description,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		return models.Developer{}, err
	}

	var response map[string]string
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &response)
	if err != nil {
		return models.Developer{}, err
	}

	devName, ok := response["nome"]
	if !ok {
		return models.Developer{}, ErrDeveloperNotChosen
	}

	dev, ok := developersMap[devName]
	if !ok {
		return models.Developer{}, ErrDeveloperNotChosen
	}

	return dev, nil
}

func (o *OpenAIDealer) GetTaskForDeveloper(developer models.Developer, tasks []models.Task) (models.Task, error) {
	var tasksPrompt strings.Builder
	tasksMap := make(map[string]models.Task)

	for _, task := range tasks {
		tasksPrompt.WriteString(fmt.Sprintf("Título: %s - Descrição: %s\n", task.Title, task.Description))
		tasksMap[task.Title] = task
	}

	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "Responda qual a tarefa mais indicada para o desenvolvedor em questão.\n" +
						"A sua resposta deve ser um objeto JSON com um único campo \"title\" contendo o título da tarefa escolhida.\n" +
						"As tarefas disponíveis são:\n" + tasksPrompt.String() +
						"O desenvolvedor é:\n" + developer.Name + ": " + developer.Description,
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		return models.Task{}, err
	}

	var response map[string]string
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &response)
	if err != nil {
		return models.Task{}, err
	}

	taskName, ok := response["title"]
	if !ok {
		return models.Task{}, ErrTaskNotChosen
	}

	dev, ok := tasksMap[taskName]
	if !ok {
		return models.Task{}, ErrTaskNotChosen
	}

	return dev, nil
}

func (o *OpenAIDealer) GetPairForDeveloper(mainDeveloper models.Developer, task models.Task, developers []models.Developer) (models.Developer, error) {
	var developersPrompt strings.Builder
	developersMap := make(map[string]models.Developer)

	for _, dev := range developers {
		if dev.ID == mainDeveloper.ID {
			continue
		}

		developersPrompt.WriteString(fmt.Sprintf("Nome: %s - Descrição: %s\n", dev.Name, dev.Description))
		developersMap[dev.Name] = dev
	}

	resp, err := o.client.CreateChatCompletion(
		context.Background(),
		openai.ChatCompletionRequest{
			Model: openai.GPT3Dot5Turbo,
			Messages: []openai.ChatCompletionMessage{
				{
					Role: openai.ChatMessageRoleSystem,
					Content: "Responda qual o desenvolvedor mais indicado para realizar programação pareada com o desenvolvedor descrito para a tarefa especificada.\n" +
						"A sua resposta deve ser um objeto JSON com um único campo \"nome\" contendo o nome do desenvolvedor escolhido.\n" +
						"O desenvolvedor principal:\n" + mainDeveloper.Description +
						"A tarefa é:\n" + task.Title + ": " + task.Description +
						"Os desenvolvedores disponíveis são:\n" + developersPrompt.String(),
				},
			},
			ResponseFormat: &openai.ChatCompletionResponseFormat{
				Type: openai.ChatCompletionResponseFormatTypeJSONObject,
			},
		},
	)
	if err != nil {
		return models.Developer{}, err
	}

	var response map[string]string
	err = json.Unmarshal([]byte(resp.Choices[0].Message.Content), &response)
	if err != nil {
		return models.Developer{}, err
	}

	devName, ok := response["nome"]
	if !ok {
		return models.Developer{}, ErrDeveloperNotChosen
	}

	dev, ok := developersMap[devName]
	if !ok {
		return models.Developer{}, ErrDeveloperNotChosen
	}

	return dev, nil
}
