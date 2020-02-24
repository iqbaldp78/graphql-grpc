package repository

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"github.com/mbizmarket/dmp/radiance/models"
	"github.com/mbizmarket/dmp/radiance/proto/users/pb"
)

//Service --
type Service struct {
	DB *gorm.DB
}

//GetAllUsers --
func (s *Service) GetAllUsers(ctx context.Context, in *empty.Empty, out *pb.RespUserData) error {
	log.Println("GetAllUsers received .....")
	users := []models.Users{}
	pbUsers := []*pb.Users{}
	s.DB.Debug().Limit(10000).Find(&users)
	copier.Copy(&pbUsers, &users)

	for i := range users {
		if users[i].LastLogin == nil {
			pbUsers[i].LastLogin = nil
		} else {
			pbUsers[i].LastLogin, _ = ptypes.TimestampProto(*users[i].LastLogin)
		}

		if users[i].DeletedAt == nil {
			pbUsers[i].DeletedAt = nil
		} else {
			pbUsers[i].DeletedAt, _ = ptypes.TimestampProto(*users[i].DeletedAt)
		}

		pbUsers[i].CreatedAt, _ = ptypes.TimestampProto(users[i].CreatedAt)
		pbUsers[i].UpdatedAt, _ = ptypes.TimestampProto(users[i].UpdatedAt)

	}
	out.Data = pbUsers
	return nil
}
