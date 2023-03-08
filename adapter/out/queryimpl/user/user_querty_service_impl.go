package queryimpl

import (
	"github.com/lj19950508/ddd-demo-go/adapter/out/repositoryimpl/user/po"
	"github.com/lj19950508/ddd-demo-go/application/query/user"
	"github.com/lj19950508/ddd-demo-go/pkg/db"
	"github.com/lj19950508/ddd-demo-go/pkg/logger"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)


type UserQueryServiceImpl struct {
	*db.DB
	logger.Interface
}

func NewUserQueryServiceImpl(mysql *db.DB, logger logger.Interface) query.UserQueryService {
	return &UserQueryServiceImpl{
		mysql,
		logger,
	}
}

func (t *UserQueryServiceImpl) FindOne(cond *query.UserQuery) (*query.UserResult, error) {
	var userPo po.User
	//组装查询语句
	if cond.IdEq != nil {
		t.GormDb.Where("id=?", *cond.IdEq)
	}
	if cond.NameLike != nil {
		t.GormDb.Where("name LIKE ?", "%"+*cond.NameLike+"%")
	}

	if result := t.GormDb.First(&userPo); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, errors.WithStack(result.Error)
	}

	result := query.NewUserResult(userPo.ID, userPo.Name)
	return result, nil

}

func (t *UserQueryServiceImpl) FindList(cond *query.UserPageQuery) (*query.PageResult[query.UserResult], error) {
	var userPo []po.User
	var count int64
	//组装查询语句
	if result := t.GormDb.Scopes(func(d *gorm.DB) *gorm.DB {
		if cond.IdEq != nil {
			d.Where("id=?", *cond.IdEq)
		}
		if cond.NameLike != nil {
			d.Where("name LIKE ?", "%"+*cond.NameLike+"%")
		}
		d.Offset(cond.Page * cond.Size).Limit(cond.Size)
		return d
	}).Find(&userPo).Count(&count); result.Error != nil {
		return nil, errors.WithStack(result.Error)
	}

	var result []query.UserResult
	for i := range userPo {
		result = append(result, *query.NewUserResult(userPo[i].ID, userPo[i].Name))
	}
	//
	return query.NewPageResult(result, count), nil
}
