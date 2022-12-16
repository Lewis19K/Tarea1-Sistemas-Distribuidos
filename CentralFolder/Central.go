package main

import (
	"context"
	"fmt"
	"os"
	"time"

	//"sync"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea1-Grupo08/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

var escuadra = 2 // Placeholder para numero de escuadra
var portAux string
var hostS string

const lab1 = "Laboratorio Renca (la lleva) - Chile"
const lab2 = "Laboratorio Pohang - Korea"
const lab3 = "Laboratorio Kampala - Uganda"
const lab4 = "Laboratorio Pripiat - Rusia"

const port1 = ":50051"
const port2 = ":50052"
const port3 = ":50053"
const port4 = ":50054"

func registroSolicitudes(f *os.File, labName string, contadorConsultas int) {
	_, err := f.WriteString(string(labName))
	if err != nil {
		fmt.Println(err)
	}
	_, err2 := f.WriteString(";")
	if err2 != nil {
		fmt.Println(err)
	}
	_, err3 := f.WriteString(fmt.Sprintf("%d\n", contadorConsultas))
	if err3 != nil {
		fmt.Println(err)
	}
}

func loopEscuadron(msgs <-chan amqp.Delivery, hostQ string, f *os.File, escuadra int) {
	for d := range msgs {
		// laboratorio que pide ayuda
		labName := d.Body
		fmt.Println("-------------------------------------------------------------")
		fmt.Printf("Mensaje asincrono recibido de %s leido\n", labName)
		fmt.Printf("Se envia Escuadra %d a %s\n", escuadra, labName)

		switch string(labName) {
		case lab1:
			portAux = port1
			hostS = "dist029"
		case lab2:
			portAux = port2
			hostS = "dist030"
		case lab3:
			portAux = port3
			hostS = "dist031"
		case lab4:
			portAux = port4
			hostS = "dist032"
		}

		connS, errS := grpc.Dial(hostS+portAux, grpc.WithInsecure()) //crea la conexion sincrona Central-Laboratorio

		if errS != nil {
			panic("No se pudo conectar con el servidor" + errS.Error())
		}

		serviceClient := pb.NewMessageServiceClient(connS)

		// Enviar escuadron
		switch string(labName) {
		case lab1:
			_, errS2 := serviceClient.EnviarEscuadron(context.Background(),
				&pb.ReqEnviarEscuadron{Escuadron: int64(escuadra)})

			if errS2 != nil {
				panic("No se puede crear el mensaje " + errS2.Error())
			}
		case lab2:
			_, errS2 := serviceClient.EnviarEscuadron2(context.Background(),
				&pb.ReqEnviarEscuadron{Escuadron: int64(escuadra)})

			if errS2 != nil {
				panic("No se puede crear el mensaje " + errS2.Error())
			}
		case lab3:
			_, errS2 := serviceClient.EnviarEscuadron3(context.Background(),
				&pb.ReqEnviarEscuadron{Escuadron: int64(escuadra)})

			if errS2 != nil {
				panic("No se puede crear el mensaje " + errS2.Error())
			}
		case lab4:
			_, errS2 := serviceClient.EnviarEscuadron4(context.Background(),
				&pb.ReqEnviarEscuadron{Escuadron: int64(escuadra)})

			if errS2 != nil {
				panic("No se puede crear el mensaje " + errS2.Error())
			}
		}

		contadorConsultas := 0
		for {
			time.Sleep(5 * time.Second)

			// (3) Solicitando status escuadron ********
			contadorConsultas += 1
			if string(labName) == lab1 {
				resStatus, errS3 := serviceClient.EstadoLab(context.Background(),
					&pb.ReqEstadoLab{})

				if errS3 != nil {
					panic("No se puede crear el mensaje " + errS3.Error())
				}
				if resStatus.Ready == true {
					fmt.Printf("Status Escuadra %d: [LISTO]\n", escuadra)
					fmt.Printf("Retorno a Central Escuadra %d, Conexion Laboratorio %s Cerrada\n", escuadra, labName)
					//fmt.Println("-------------------------------------------------------------")
					break
				} else {
					fmt.Printf("Status Escuadra %d: [NO LISTO]\n", escuadra)
				}

			} else if string(labName) == lab2 {
				resStatus, errS3 := serviceClient.EstadoLab2(context.Background(),
					&pb.ReqEstadoLab{})

				if errS3 != nil {
					panic("No se puede crear el mensaje " + errS3.Error())
				}
				if resStatus.Ready == true {
					fmt.Printf("Status Escuadra %d: [LISTO]\n", escuadra)
					fmt.Printf("Retorno a Central Escuadra %d, Conexion Laboratorio %s Cerrada\n", escuadra, labName)
					//fmt.Println("-------------------------------------------------------------")
					break
				} else {
					fmt.Printf("Status Escuadra %d: [NO LISTO]\n", escuadra)
				}
			}
		}
		registroSolicitudes(f, string(labName), contadorConsultas)
	}
	//wg.Done()
}

func main() {
	hostQ := "localhost"
	//hostS := "localhost"

	// Conexion a RabbitMQ
	conn, err := amqp.Dial("amqp://guest:guest@" + hostQ + ":5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	fmt.Println("Central conectada a RabbitMQ")

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	// (1) Se crea cola asincronta en RabbitMQ para la Central
	q, err := ch.QueueDeclare(
		"CentralQueue",
		false,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		fmt.Println(err)
		panic(err)
	}

	fmt.Println(q)

	msgs, err := ch.Consume(
		"CentralQueue",
		"",
		true,
		false,
		false,
		false,
		nil,
	)

	// Se crea archivo .txt para las solicitudes
	f, err := os.Create("SOLICITUDES.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	// ejecutar rutinas de escuadrones (falta ver como coordinar disponibilidad)

	forever := make(chan bool)

	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//go func(){
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	go loopEscuadron(msgs, hostQ, f, 1)
	//wg.Wait()

	//}

	go loopEscuadron(msgs, hostQ, f, 2)

	fmt.Println(" [*] - Esperando estallidos")
	<-forever
}
