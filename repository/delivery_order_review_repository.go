package repository

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/infrastructure"
	"myapp/model"

	"github.com/Masterminds/squirrel"
)

type DeliveryOrderReviewRepository interface {
	// create
	Insert(ctx context.Context, deliveryOrderReview *model.DeliveryOrderReview) error
	InsertMany(ctx context.Context, deliveryOrderReviews []model.DeliveryOrderReview, options ...data_type.RepositoryOption) error

	// read
	Count(ctx context.Context, options ...model.DeliveryOrderReviewQueryOption) (int, error)
	Fetch(ctx context.Context, options ...model.DeliveryOrderReviewQueryOption) ([]model.DeliveryOrderReview, error)
	FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderReview, error)
	Get(ctx context.Context, id string) (*model.DeliveryOrderReview, error)
	IsExistByDeliveryOrderId(ctx context.Context, deliveryOrderId string) (bool, error)

	// delete
	Truncate(ctx context.Context) error
}

type deliveryOrderReviewRepository struct {
	db infrastructure.DBTX
}

func NewDeliveryOrderReviewRepository(
	db infrastructure.DBTX,
	loggerStack infrastructure.LoggerStack,
) DeliveryOrderReviewRepository {
	return &deliveryOrderReviewRepository{
		db: db,
	}
}

func (r *deliveryOrderReviewRepository) fetch(ctx context.Context, stmt squirrel.Sqlizer) ([]model.DeliveryOrderReview, error) {
	deliveryOrderReviews := []model.DeliveryOrderReview{}
	if err := fetch(r.db, ctx, &deliveryOrderReviews, stmt); err != nil {
		return nil, err
	}

	return deliveryOrderReviews, nil
}

func (r *deliveryOrderReviewRepository) get(ctx context.Context, stmt squirrel.Sqlizer) (*model.DeliveryOrderReview, error) {
	deliveryOrderReview := model.DeliveryOrderReview{}
	if err := get(r.db, ctx, &deliveryOrderReview, stmt); err != nil {
		return nil, err
	}

	return &deliveryOrderReview, nil
}

func (r *deliveryOrderReviewRepository) prepareQuery(option model.DeliveryOrderReviewQueryOption) squirrel.SelectBuilder {
	stmt := stmtBuilder.Select().
		From(fmt.Sprintf("%s dor", model.DeliveryOrderReviewTableName))

	stmt = model.Prepare(stmt, &option)

	return stmt
}

func (r *deliveryOrderReviewRepository) Insert(ctx context.Context, deliveryOrderReview *model.DeliveryOrderReview) error {
	return defaultInsert(r.db, ctx, deliveryOrderReview, "*")
}

func (r *deliveryOrderReviewRepository) InsertMany(ctx context.Context, deliveryOrderReviews []model.DeliveryOrderReview, options ...data_type.RepositoryOption) error {
	arr := []model.BaseModel{}
	for i := range deliveryOrderReviews {
		arr = append(arr, &deliveryOrderReviews[i])
	}

	return defaultInsertMany(r.db, ctx, arr, "*")
}

func (r *deliveryOrderReviewRepository) Count(ctx context.Context, options ...model.DeliveryOrderReviewQueryOption) (int, error) {
	option := model.DeliveryOrderReviewQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}
	option.IsCount = true

	stmt := r.prepareQuery(option)

	return count(r.db, ctx, stmt)
}

func (r *deliveryOrderReviewRepository) Fetch(ctx context.Context, options ...model.DeliveryOrderReviewQueryOption) ([]model.DeliveryOrderReview, error) {
	option := model.DeliveryOrderReviewQueryOption{}
	if len(options) > 0 {
		option = options[0]
	}

	stmt := r.prepareQuery(option)

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderReviewRepository) FetchByDeliveryOrderIds(ctx context.Context, deliveryOrderIds []string) ([]model.DeliveryOrderReview, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderReviewTableName).
		Where(squirrel.Eq{"delivery_order_id": deliveryOrderIds})

	return r.fetch(ctx, stmt)
}

func (r *deliveryOrderReviewRepository) Get(ctx context.Context, id string) (*model.DeliveryOrderReview, error) {
	stmt := stmtBuilder.Select("*").
		From(model.DeliveryOrderReviewTableName).
		Where(squirrel.Eq{"id": id})

	return r.get(ctx, stmt)
}

func (r *deliveryOrderReviewRepository) IsExistByDeliveryOrderId(ctx context.Context, deliveryOrderId string) (bool, error) {
	stmt := stmtBuilder.Select().Column(
		stmtBuilder.Select("*").
			From(model.DeliveryOrderReviewTableName).
			Where(squirrel.Eq{"delivery_order_id ": deliveryOrderId}).
			Prefix("EXISTS (").Suffix(")"),
	)

	return isExist(r.db, ctx, stmt)
}

func (r *deliveryOrderReviewRepository) Truncate(ctx context.Context) error {
	return truncate(r.db, ctx, model.DeliveryOrderReviewTableName)
}
