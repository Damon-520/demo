package live_room

import (
	"context"
	"demoapi/internal/pkg/gorm_builder"
	"demoapi/internal/repository"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
)

type LiveRoomDao struct {
	aDb *repository.ActivityDB
	log *log.Helper
}

func NewLiveRoomDao(asDb *repository.ActivityDB, logger log.Logger) ILiveRoomDao {
	dao := LiveRoomDao{
		aDb: asDb,
		log: log.NewHelper(log.With(logger, "x_module", "repository/NewLiveRoomDao")),
	}
	return &dao
}

func (l *LiveRoomDao) Info(ctx context.Context, options gorm_builder.Options) (info LiveRoom, err error) {

	db := l.aDb.ActivityR.WithContext(ctx).Model(&LiveRoom{})

	cond, values := gorm_builder.BuildWhere(options)
	if cond != "" {
		db = db.Where(cond, values...)
	}

	err = db.First(&info).Error

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return
	}

	return info, nil
}

func (l *LiveRoomDao) List(ctx context.Context, options gorm_builder.Options) (list []LiveRoom, total int64, err error) {

	db := l.aDb.ActivityR.
		WithContext(ctx).
		Model(&LiveRoom{})

	cond, values := gorm_builder.BuildWhere(options)
	if cond != "" {
		db = db.Where(cond, values...)
	}

	if options.IsCount { // 获取总数
		err = db.Count(&total).Error
		if err != nil {
			return
		}
		if total == 0 {
			return
		}
	}

	if options.Limit > 0 {
		db = db.Limit(options.Limit)
	}
	if options.Offset > 0 {
		db = db.Offset(options.Offset)
	}
	if options.Order != "" {
		db = db.Order(options.Order)
	}

	err = db.Find(&list).Error

	return
}

func (l *LiveRoomDao) Count(ctx context.Context, options gorm_builder.Options) (total int, err error) {
	// TODO implement me
	panic("implement me")
}

func (l *LiveRoomDao) Save(ctx context.Context, data LiveRoom) (lastId int, err error) {

	r := l.aDb.ActivityW.WithContext(ctx).Model(LiveRoom{}).Create(&data)

	if r.Error != nil {
		return 0, r.Error
	}

	return data.ID, nil
}

func (l *LiveRoomDao) Update(ctx context.Context, options gorm_builder.Options, data map[string]interface{}) (int64, error) {

	db := l.aDb.ActivityW.WithContext(ctx).Model(LiveRoom{})

	cond, values := gorm_builder.BuildWhere(options)

	if cond == "" {
		return 0, errors.New("更新条件不能为空")
	}

	db = db.Where(cond, values...)

	db.Updates(data)

	return db.RowsAffected, db.Error
}
