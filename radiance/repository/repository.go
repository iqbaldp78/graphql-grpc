package repository

import (
	"context"
	"log"

	"github.com/golang/protobuf/ptypes"
	"github.com/golang/protobuf/ptypes/empty"
	"github.com/jinzhu/copier"
	"github.com/jinzhu/gorm"

	"github.com/mbizmarket/dmp/radiance/models"
	pbRfqs "github.com/mbizmarket/dmp/radiance/proto/pb/rfqs"
	pbUser "github.com/mbizmarket/dmp/radiance/proto/pb/user"
)

//Service --
type Service struct {
	DB            *gorm.DB
	ClientAghanim pbRfqs.AghanimServicesService
}

//GetAllUsers --
func (s *Service) GetAllUsers(ctx context.Context, in *empty.Empty, out *pbUser.RespUserData) error {
	log.Println("GetAllUsers received .....")
	users := []models.Users{}
	pbUsers := []*pbUser.Users{}
	// s.DB.Debug().Limit(30).Find(&users)
	s.DB.Debug().Find(&users)
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

//GetUserNRfqs --
func (s *Service) GetUserNRfqs(ctx context.Context, in *pbRfqs.Req, out *pbRfqs.RespData) error {
	log.Println("GetUserNRfqs received .....")
	pbRfqs, err := s.ClientAghanim.GetRfqsByUser(ctx, in)
	if err != nil {
		return err
	}

	out.Data = pbRfqs.Data
	return nil
}
