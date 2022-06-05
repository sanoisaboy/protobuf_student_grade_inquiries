package main

import (
	"context"
	"errors"
	"flag"
	"log"
	"net"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jackc/pgx/v4"

	httpstatus "github.com/sanoisaboy/restful/pkg/http_status"
	pb "github.com/sanoisaboy/restful/pkg/proto/core/v1"
	re "github.com/sanoisaboy/restful/pkg/repository"
	"google.golang.org/grpc"
)

type Server struct {
	pb.UnimplementedStudentserviceServer
}

func (s *Server) ListStudent(ctx context.Context, req *pb.ListStudentRequest) (*pb.ListStudentResponse, error) {
	//connect to the database
	newRepository := re.NewCrdbRepository(connect_string)

	//get the data from the database
	student_list, err := newRepository.ListStudent(ctx)
	if err != nil {
		log.Println(err)
	}
	newUser := []pb.User{}
	newListuser := []*pb.User{}

	for i := 0; i < len(student_list); i++ {

		newUser = append(newUser, pb.User{StudentName: student_list[i].Student_name, ID: int32(student_list[i].ID), POINT: int32(student_list[i].POINT)})
		newListuser = append(newListuser, &newUser[i])
	}

	//select the reponse message type
	res := &pb.ListStudentResponse{
		User: newListuser,
	}

	return res, nil
}

func (s *Server) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {
	//connect to the database
	newRepository := re.NewCrdbRepository(connect_string)

	//get the data from the database
	student_get, err := newRepository.GetStudent(ctx, int(req.Id))
	if err != nil {
		log.Println(err)
	}

	//check the get data success or not
	http_status_code, err := httpstatus.Get_http_status(student_get.Student_name)

	if err != nil {
		log.Println(err)
	}

	//select the reponse message type
	res := &pb.GetStudentResponse{}
	if http_status_code != 200 {
		res = &pb.GetStudentResponse{
			Success: "Get error",
		}
	} else {
		res = &pb.GetStudentResponse{
			Success: "Get success",
			Name:    student_get.Student_name,
			Id:      int32(student_get.ID),
			Point:   int32(student_get.POINT),
		}
	}
	return res, nil
}

func (Server) CreateStudent(ctx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {
	//connect to the database
	newRepository := re.NewCrdbRepository(connect_string)
	err := req.Validate()
	if err != nil {
		return nil, errors.New("request format error")
	}

	//check data need query success or not
	http_status_code, err := httpstatus.Create_http_status(req.Name, string(req.Id), string(req.Point))
	if err != nil {
		log.Println(err)
	}

	//insert data to the database success or not
	http_status_code, err = newRepository.CreateStudent(ctx, req.Name, int(req.Id), int(req.Point))
	if err != nil {
		log.Println(err)
	}

	//select the reponse message type
	res := &pb.CreateStudentResponse{}
	if http_status_code != 201 {
		res = &pb.CreateStudentResponse{
			Success: "create error",
		}
	} else {
		res = &pb.CreateStudentResponse{
			Success: "create success",
			Name:    req.Name,
			Id:      req.Id,
			Point:   req.Point,
		}
	}

	return res, nil
}

func (Server) UpdateStudent(ctx context.Context, req *pb.UpdateStudentResquest) (*pb.UpdateStudentResponse, error) {
	//connect to the database
	newRepository := re.NewCrdbRepository(connect_string)

	//check data need query success or not
	http_status_code, err := httpstatus.Update_http_status(string(req.Id), string(req.Point))
	if err != nil {
		log.Println(err)
	}

	//Update data to the database success or not
	http_status_code, student_update, err := newRepository.UpdateStudent(ctx, int(req.Id), int(req.Point))

	if err != nil {
		log.Println(err)
	}

	//select the reponse message type
	res := &pb.UpdateStudentResponse{}
	if http_status_code != 200 {
		res = &pb.UpdateStudentResponse{
			Updatesuccess: "Update error",
		}
	} else {
		res = &pb.UpdateStudentResponse{
			Updatesuccess: "Update success",
			Name:          student_update[0].Student_name,
			Id:            int32(student_update[0].ID),
			Point:         int32(student_update[0].POINT),
		}
	}

	return res, nil
}

func (Server) DeleteStudent(ctx context.Context, req *pb.DeleteStudentResquest) (*pb.DeleteStudentResponse, error) {
	//connect to the database
	newRepository := re.NewCrdbRepository(connect_string)

	//check data need query success or not
	http_status_code, err := newRepository.DeleteStudent(ctx, int(req.Id))
	if err != nil {
		log.Println(err)
	}

	//Delete data to the database success or not
	http_status_code, err = httpstatus.Delete_http_status(int(req.Id))
	if err != nil {
		log.Println(err)
	}

	//select the reponse message type
	res := &pb.DeleteStudentResponse{}
	if http_status_code != 204 {
		res = &pb.DeleteStudentResponse{
			Deletesuccess: "Delete error",
		}
	} else {
		res = &pb.DeleteStudentResponse{
			Deletesuccess: "Delete success",
		}
	}

	return res, nil
}

var connect_string string

func init() {
	flag.StringVar(&connect_string, "connect_string", "Default string", "Set String")
}

func main() {
	flag.Parse()
	//staring database
	connn, err := pgx.Connect(context.Background(), connect_string)
	if err != nil {
		log.Println(err)
	}

	defer connn.Close(context.Background())
	log.Println("starting connect to cockroach db ...")

	//staring grpc
	log.Println("starting gRPC server...")

	grpcServer := grpc.NewServer()
	pb.RegisterStudentserviceServer(grpcServer, &Server{})

	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		log.Println("failed to listen: %v \n", err)
	}

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			log.Println("fail to serve:%v", err)
		}
	}()

	conn, err := grpc.Dial("localhost:3000", grpc.WithInsecure())
	if err != nil {
		log.Println("fail to dial server", err)
	}

	gmux := runtime.NewServeMux()

	err = pb.RegisterStudentserviceHandler(context.Background(), gmux, conn)
	if err != nil {
		log.Println("fail to register gateway", err)
	}

	gwServer := &http.Server{
		Addr:    ":3001",
		Handler: gmux,
	}
	log.Printf("Serving gRPC-Gateway on http://0.0.0.0:3001")
	log.Println(gwServer.ListenAndServe())
}
