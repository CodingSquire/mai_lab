package repository

import (
	"context"
	"errors"
	"fmt"
	"go_services_lab/models"
	"go_services_lab/pkg/order/repository"
	"go_services_lab/pkg/user/proto"
	"log"
	"strconv"

	"google.golang.org/grpc"
)

type OrderService struct {
	rep repository.Order
}

func NewOrderService(rep repository.Order) *OrderService {
	return &OrderService{rep: rep}
}

func (s *OrderService) Get(id int) (models.Order, error) {
	return s.rep.Get(id)
}

func (s *OrderService) Amount(id int) (float32, error) {
	return s.rep.Amount(id)
}

func (s *OrderService) GetAll() ([]models.Order, error) {
	return s.rep.GetAll()
}

func (s *OrderService) Delete(id int) (int, error) {
	return s.rep.Delete(id)
}

func (s *OrderService) Create(user_id int, products map[string]int) (int, error) {
	conn, err := grpc.Dial("go_users_service_container:8080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := proto.NewUsersClient(conn)
	res, err := client.IsExistById(context.Background(), &proto.IsExistByIdRequest{Id: int32(user_id)})
	if err != nil {
		log.Fatal(err)
	}

	isUserExist := res.GetIsExist()
	if !isUserExist {
		return 0, fmt.Errorf("There is no user with id == %d", user_id)
	}

	pr, err := atoiMap(products)

	if err != nil {
		return 0, err
	}

	return s.rep.Create(user_id, pr)
}

func atoiMap(products map[string]int) (map[int]int, error) {
	ret_map := make((map[int]int))

	for key, val := range products {
		k, err := strconv.Atoi(key)
		if err != nil {
			return ret_map, errors.New("Wronge key for product's ID.")
		} else {
			ret_map[k] = val
		}
	}

	return ret_map, nil
}
