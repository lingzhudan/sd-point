package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	rand2 "math/rand"
	"sd-point/app/user/service/internal/biz"
	"sd-point/app/user/service/internal/define"
	"strconv"
	"strings"
	"time"
)

type userRepo struct {
	data *Data
	log  *log.Helper
}

// NewUserRepo .
func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

type User struct {
	UID         uint32 `gorm:"column:uid;primaryKey;comment:用户编号;"`
	Name        string `gorm:"column:name;size:32;comment:用户名称;"`
	Account     string `gorm:"column:account;uniqueIndex;size:64;comment:系统账号;"`
	Password    string `gorm:"column:password;size:64;comment:加密系统密码;"`
	Openid      string `gorm:"column:openid;size:64;comment:微信openID;"`
	PhoneNumber string `gorm:"column:phone_number;size:16;comment:手机号;"`

	CreatedAt time.Time      `gorm:"column:created_at;autoCreateTime;comment:创建时间;"`
	UpdatedAt time.Time      `gorm:"column:updated_at;autoUpdateTime;comment:更新时间;"`
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at;comment:删除时间;"`
}

func (u *User) TableName() string {
	return "user"
}

type UserCond struct {
	Begin        int
	Count        int
	UIDs         []uint32
	Accounts     []string
	Passwords    []string
	Openids      []string
	PhoneNumbers []string
}

func (c *UserCond) ParseCond() (whereStage string, args []interface{}) {
	var whereStages []string
	if len(c.UIDs) != 0 {
		whereStages = append(whereStages, "`uid` IN ?")
		args = append(args, c.UIDs)
	}
	if len(c.Accounts) != 0 {
		whereStages = append(whereStages, "`account` IN ?")
		args = append(args, c.Accounts)
	}
	if len(c.Passwords) != 0 {
		whereStages = append(whereStages, "`password` IN ?")
		args = append(args, c.Passwords)
	}
	if len(c.Openids) != 0 {
		whereStages = append(whereStages, "`openid` IN ?")
		args = append(args, c.Openids)
	}
	if len(c.PhoneNumbers) != 0 {
		whereStages = append(whereStages, "`phone_number` IN ?")
		args = append(args, c.PhoneNumbers)
	}
	return strings.Join(whereStages, " AND "), args
}

func (r *userRepo) GetUser(ctx context.Context, uid uint32) (bU *biz.User, err error) {
	dU := &User{}
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{UID: uid}).
		First(dU).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("db error: %v", err)
		return
	}
	bU = newBizUser(dU)
	return
}

func (r *userRepo) ListUser(ctx context.Context, uids []uint32) (bUs []*biz.User, err error) {
	c := &UserCond{UIDs: uids}
	ws, args := c.ParseCond()
	var dUs []*User
	if err = r.data.db.
		WithContext(ctx).
		Where(ws, args...).
		Find(&dUs).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("db error: %v", err)
	}
	for _, u := range dUs {
		bUs = append(bUs, newBizUser(u))
	}
	return
}
func (r *userRepo) Register(ctx context.Context, account string, password string) (uid uint32, err error) {
	u := &User{
		Name:     account,
		Account:  account,
		Password: password,
	}
	encryptionPassword(u)
	if err = r.data.db.
		WithContext(ctx).
		Create(u).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			err = define.ErrDuplicateKey
		}
	}
	uid = u.UID
	return
}

func (r *userRepo) RegisterByWechat(ctx context.Context, openid string) (uid uint32, err error) {
	u := newUserOriginAccount()
	encryptionPassword(u)
	u.Openid = openid
	if err = r.data.db.
		WithContext(ctx).
		Create(u).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			err = define.ErrDuplicateKey
		}
	}
	uid = u.UID
	return
}

func (r *userRepo) RegisterByPhoneNumber(ctx context.Context, phoneNumber string) (uid uint32, err error) {
	u := newUserOriginAccount()
	encryptionPassword(u)
	u.PhoneNumber = phoneNumber
	if err = r.data.db.
		WithContext(ctx).
		Create(u).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			err = define.ErrDuplicateKey
		}
	}
	uid = u.UID
	return
}

func (r *userRepo) GetUserByAccount(ctx context.Context, account string) (bU *biz.User, err error) {
	dU := &User{}
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{Account: account}).
		First(dU).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("db error: %v", err)
		return
	}
	bU = newBizUser(dU)
	return
}

func (r *userRepo) GetUserByWechat(ctx context.Context, openid string) (bU *biz.User, err error) {
	dU := &User{}
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{Openid: openid}).
		First(dU).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("db error: %v", err)
		return
	}
	bU = newBizUser(dU)
	return
}

func (r *userRepo) GetUserByPhoneNumber(ctx context.Context, phoneNumber string) (bU *biz.User, err error) {
	dU := &User{}
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{PhoneNumber: phoneNumber}).
		First(dU).
		Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			err = define.ErrRecordNotFound
			return
		}
		r.log.Errorf("db error: %v", err)
		return
	}
	bU = newBizUser(dU)
	return
}

func (r *userRepo) BindWechat(ctx context.Context, uid uint32, openid string) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{UID: uid}).
		Updates(&User{Openid: openid}).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
		return
	}
	return
}

func (r *userRepo) BindPhoneNumber(ctx context.Context, uid uint32, phoneNumber string) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{UID: uid}).
		Updates(&User{PhoneNumber: phoneNumber}).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) UnbindWechat(ctx context.Context, openid string) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{Openid: openid}).
		Update("`openid`", "").
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) UnbindPhoneNumber(ctx context.Context, phoneNumber string) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Where(&User{PhoneNumber: phoneNumber}).
		Update("`phone_number`", "").
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func (r *userRepo) DeleteUser(ctx context.Context, uid uint32) (err error) {
	if err = r.data.db.
		WithContext(ctx).
		Delete(&User{}, uid).
		Error; err != nil {
		r.log.Errorf("db error: %v", err)
	}
	return
}

func newBizUser(dU *User) (bU *biz.User) {
	bU = &biz.User{
		UID:         dU.UID,
		Name:        dU.Name,
		Account:     dU.Account,
		Password:    dU.Password,
		Openid:      dU.Openid,
		PhoneNumber: dU.PhoneNumber,
		CreatedAt:   dU.CreatedAt,
		UpdatedAt:   dU.UpdatedAt,
		DeletedAt:   dU.DeletedAt.Time,
	}
	return
}

// 通过用户的某些属性和密码混合加密生成加密数据写回用户结构体中
func encryptionPassword(u *User) {
	// TODO 自定义生成相关机密
	if len(u.Password) == 0 {
		panic("empty password to encrypt")
	}
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}
	salt := u.CreatedAt
	u.Password = strconv.FormatInt(salt.Unix(), 36) + u.Password
	return
}

// 创建随机的原生账号密码
//   - 用户名默认为账号
func newUserOriginAccount() (u *User) {
	u = new(User)
	// TODO 自定义生成相关账号密码
	// 随机账号为`sd_`+`10位时间生成的36进制数`+`3位随机36进制数`
	// 随机密码为`10位随机36进制数`
	var ab strings.Builder
	var pb strings.Builder
	ab.WriteString("sd_")
	ab.WriteString(strconv.FormatInt(time.Now().Unix(), 36))
	pb.WriteString(strconv.FormatInt(rand2.Int63n(time.Now().UnixMicro()), 36))
	for ab.Len() < 13 {
		ab.WriteString("0")
	}
	for ab.Len() < 16 {
		ab.WriteString(strconv.FormatInt(rand2.Int63n(36), 36))
	}
	for pb.Len() < 10 {
		pb.WriteString(strconv.FormatInt(rand2.Int63n(36), 36))
	}
	u.Account = ab.String()
	u.Password = pb.String()
	u.Name = u.Account
	return
}
