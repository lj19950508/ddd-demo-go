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
	//tx:=context.GET(Tx)
	//if tx==nil -> tx.create // tx.trascation(func(db *grom)) -> db.Create
	// PROPAGATION_REQUIRED：如果不存在外层事务，就主动创建事务；否则使用外层事务
	// PROPAGATION_SUPPORTS：如果不存在外层事务，就不开启事务；否则使用外层事务
	// PROPAGATION_MANDATORY：如果不存在外层事务，就抛出异常；否则使用外层事务
	// PROPAGATION_REQUIRES_NEW：总是主动开启事务；如果存在外层事务，就将外层事务挂起
	// PROPAGATION_NOT_SUPPORTED：总是不开启事务；如果存在外层事务，就将外层事务挂起
	// PROPAGATION_NEVER：总是不开启事务；如果存在外层事务，则抛出异常
	// PROPAGATION_NESTED：如果不存在外层事务，就主动创建事务；否则创建嵌套的子事务
	var userPo []po.User
	var count int64
	//组装查询语句
	scope := func(d *gorm.DB) *gorm.DB {
		if cond.IdEq != nil {
			d.Where("id=?", *cond.IdEq)
		}
		if cond.NameLike != nil {
			d.Where("name LIKE ?", "%"+*cond.NameLike+"%")
		}
		return d
	}

	//读一致性
	err := t.GormDb.Transaction(func(tx *gorm.DB) error {
		if db := tx.Scopes(scope).Model(&userPo).Count(&count).Offset(cond.Page * cond.Size).Find(&userPo); db.Error != nil {
			return errors.WithStack(db.Error)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	var result []query.UserResult
	for i := range userPo {
		result = append(result, *query.NewUserResult(userPo[i].ID, userPo[i].Name))
	}
	//
	return query.NewPageResult(result, count), nil
}
