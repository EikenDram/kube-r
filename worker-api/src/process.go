package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Task struct {
	Id      int64  `json:"id"`
	Path    string `json:"path"`
	Payload string `json:"payload"`
}

func process(message amqp.Delivery) {
	log.Printf(" > Received message: %s\n", message.Body)

	// get task data
	var task Task
	err := json.Unmarshal(message.Body, &task)
	if err != nil {
		log.Printf("Error decoding message body: %s\n", err)
		ackMessage(message)
		return
	}

	// update status to processing
	_ = updateStatus(task, "processing", "")

	// call server API
	err = serverAPI(task)
	if err != nil {
		log.Printf("Error running plumber API: %s\n", err)
		// update status to error
		_ = updateStatus(task, "error", err.Error())
		ackMessage(message)
		return
	}

	// save result somewhere?

	// update status to finished
	_ = updateStatus(task, "success", "")

	// remove message from queue
	log.Printf(" > Message processed\n")
	ackMessage(message)
}

func serverAPI(task Task) error {
	// call server api
	url := fmt.Sprintf("http://%s/%s", config.Server.Host, task.Path)
	var jsonStr = []byte(task.Payload)

	req, err := http.NewRequest("GET", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.SetBasicAuth(serverUser, serverPass)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response Body:", string(body))

	return nil
}

func updateStatus(task Task, status string, msg string) error {
	// call manager api
	url := fmt.Sprintf("http://%s/report/update", config.Manager.Host)
	json := fmt.Sprintf(`{"id": "%d", "status": "%s" }`, task.Id, status)
	var jsonStr = []byte(json)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	if err != nil {
		return err
	}

	req.SetBasicAuth(managerUser, managerPass)
	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	fmt.Println("Response Body:", string(body))

	return nil
}

func ackMessage(message amqp.Delivery) {
	if err := message.Ack(false); err != nil {
		log.Printf(" > Error acknowledging message: %s\n", err)
	} else {
		log.Printf(" > Message processed\n")
	}
}
