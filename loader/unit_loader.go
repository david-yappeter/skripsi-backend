package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type UnitLoader struct {
	loaderById dataloader.Loader
}

func (l *UnitLoader) loadById(id string) (*model.Unit, error) {
	thunk := l.loaderById.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Unit), nil
}

func (l *UnitLoader) ProductUnitFn(productUnit *model.ProductUnit) func() error {
	return func() error {
		unit, err := l.loadById(productUnit.UnitId)
		if err != nil {
			return err
		}

		productUnit.Unit = unit

		return nil
	}
}

func (l *UnitLoader) ProductUnitToUnitIdFn(productUnit *model.ProductUnit) func() error {
	return func() error {
		unit, err := l.loadById(*productUnit.ToUnitId)
		if err != nil {
			return err
		}

		productUnit.ToUnit = unit

		return nil
	}
}

func NewUnitLoader(unitRepository repository.UnitRepository) *UnitLoader {
	batchByUnitIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		units, err := unitRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		unitById := map[string]model.Unit{}
		for _, unit := range units {
			unitById[unit.Id] = unit
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var unit *model.Unit
			if v, ok := unitById[k.String()]; ok {
				unit = &v
			}

			result := &dataloader.Result{Data: unit, Error: nil}
			if unit == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &UnitLoader{
		loaderById: NewDataloader(batchByUnitIdFn),
	}
}
