package define

import (
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

// IsErrDuplicateKey 是否是重复唯一索引错误
func IsErrDuplicateKey(err error) bool {
	var mysqlErr *mysql.MySQLError
	return errors.As(err, &mysqlErr) && mysqlErr.Number == 1062
}

// IsErrRecordNotFound 是否是未找到记录错误
func IsErrRecordNotFound(err error) bool {
	return errors.Is(err, gorm.ErrRecordNotFound)
}
