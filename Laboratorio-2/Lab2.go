package main

import (
	"context"
	"fmt"
	"math/rand"
	"net"
	"time"

	pb "github.com/Sistemas-Distribuidos-2022-2/Tarea1-Grupo08/Proto"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

// ############### Server Interface ###############
type server struct {
	pb.UnimplementedMessageServiceServer
}

func (s *server) EnviarEscuadron2(ctx context.Context, msg *pb.ReqEnviarEscuadron) (*pb.ResEnviarEscuadron, error) {
	fmt.Printf("Llega Escuadron %d, conteniendo estallido...\n", msg.Escuadron)
	escuadronConteniendo = int(msg.Escuadron)
	return &pb.ResEnviarEscuadron{}, nil
}

// nueva variable para saber que debe volver a generar estallidos
var generarEstallidos = true

func (s *server) EstadoLab2(ctx context.Context, msg *pb.ReqEstadoLab) (*pb.ResEstadoLab, error) {
	// (3) Estado de contencion

	var estadoLab = false

	// Se crea un numero aleatorio del 0 al 9
	fmt.Println("Revisando estado Escuadron:")
	s2 := rand.NewSource(time.Now().UnixNano())
	r2 := rand.New(s2)
	arreglar := r2.Intn(10)

	if arreglar >= 4 {
		fmt.Println("[LISTO]")
		fmt.Println("Estallido contenido, Escuadron ", escuadronConteniendo, " Retornando...")
		fmt.Println("--------------------------------------------------------")
		estadoLab = true

		generarEstallidos = true
	} else {
		fmt.Println("[NO LISTO]")
	}

	return &pb.ResEstadoLab{Ready: estadoLab}, nil
}

var escuadronConteniendo = 0

func loopEstallidos(ch *amqp.Channel, labName string) {
	for {
		time.Sleep(5 * time.Second)

		if generarEstallidos == false {
			loopEstallidos(ch, labName)
		}

		// (1) Analizar probabilidad estallido
		// Se crea un numero aleatorio del 0 al 9
		fmt.Println("Analizando estado Laboratorio:")
		s1 := rand.NewSource(time.Now().UnixNano())
		r1 := rand.New(s1)
		estallido := r1.Intn(10)

		if estallido >= 8 {
			fmt.Println("[OK]")
		} else if estallido < 8 {

			// (2) Se produce un estallido
			fmt.Println("[ESTALLIDO]")
			generarEstallidos = false

			err := ch.Publish(
				"",
				"CentralQueue",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(labName), // Publica un mensaje a la cola CentralQueue
				},
			)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			fmt.Println("SOS Enviado a Central. Esperando respuesta...")
		}
	}
}

func main() {
	labName := "Laboratorio Pohang - Korea"
	hostQ := "dist030"

	fmt.Println(labName)

	// Conexion a RabbitMQ
	conn, err := amqp.Dial("amqp://test:test@" + hostQ + ":5672/")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	defer ch.Close()

	// ############### Ciclo Laboratorio ###############
	go loopEstallidos(ch, labName)

	// ############### gRPC ###############
	listener, err := net.Listen("tcp", ":50052")

	if err != nil {
		panic("No se pudo crear la conexion TCP" + err.Error())
	}

	serv := grpc.NewServer()
	for {
		pb.RegisterMessageServiceServer(serv, &server{})
		if err := serv.Serve(listener); err != nil {
			panic("El server no se pudo iniciar" + err.Error())
		}
	}
}
