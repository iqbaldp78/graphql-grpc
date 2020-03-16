package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"log"

	"github.com/mbizmarket/dmp/graphql-aghanim/graph/generated"
	"github.com/mbizmarket/dmp/graphql-aghanim/graph/model"
	"github.com/mbizmarket/dmp/graphql-aghanim/helper"
	pb "github.com/mbizmarket/dmp/graphql-aghanim/proto/pb/rfqs"
)

func (r *mutationResolver) CreateTodo2(ctx context.Context, input string) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetRfqByCompany(ctx context.Context, companyID *int) ([]*model.Rfqs, error) {
	var output []*model.Rfqs
	res, err := r.ResolverClientAghanim.GetRfqsByUser(context.Background(), &pb.Req{ID: int64(*companyID)})
	if err != nil {
		log.Println("err nya ", err)
		return output, err
	}
	for _, v := range res.Data {
		out := &model.Rfqs{
			ID:              int(v.ID),
			TransactionID:   int(v.TransactionId),
			RequestedBy:     int(v.RequestedBy),
			CreatedBy:       int(v.CreatedBy),
			PaymentType:     int(v.PaymentType),
			PaymentDuration: int(v.PaymentDuration),
			CompanyBuyerID:  int(v.CompanyBuyerId),
			CompanySellerID: int(v.CompanySellerId),
			RfqNo:           v.RfqNo,
			ReferenceNo:     v.ReferenceNo,
			Notes:           v.Notes,
			NoteForSeller:   v.NoteForSeller,
			SubTotal:        float64(v.SubTotal),
			TaxBasis:        float64(v.TaxBasis),
			Ppn:             float64(v.Ppn),
			Pph:             float64(v.Pph),
			Rounding:        float64(v.Rounding),
			GrandTotal:      float64(v.GrandTotal),
			Status:          int(v.Status),
			StatusReason:    v.StatusReason,
			ExpiredAt:       helper.TimestampProtoToPointerTime(v.ExpiredAt),
			CreatedAt:       helper.TimestampProtoToPointerTime(v.CreatedAt),
			UpdatedAt:       helper.TimestampProtoToPointerTime(v.UpdatedAt),
			DeletedAt:       helper.TimestampProtoToPointerTime(v.DeletedAt),
		}
		output = append(output, out)
	}
	return output, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
