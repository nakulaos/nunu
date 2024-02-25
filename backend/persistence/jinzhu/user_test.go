/**
 ******************************************************************************
 * @file           : user_test.go
 * @author         : nakulaos
 * @brief          : None
 * @attention      : None
 * @date           : 2024/2/13
 ******************************************************************************
 */

package jinzhu

import (
	"fmt"
	"gorm.io/gorm"
	gorm2 "nunu/backend/init/gorm"
	"nunu/backend/model"
	"nunu/backend/repo"
	"reflect"
	"testing"
)

func TestNewUserManageRepo(t *testing.T) {
	type args struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	tests := []struct {
		name string
		args args
		want repo.IUserManageRepo
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserManageRepo(tt.args.db, tt.args.ums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserManageRepo() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_CreateUser(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test user",
			fields: fields{
				db:  gorm2.GetGormDB(),
				ums: NewUserMetricsRepo(gorm2.GetGormDB()),
			},
			args: args{user: &model.User{
				Model:    nil,
				Nickname: "admin4",
				Username: "zhangsannnn",
				Phone:    "15112312233",
				Password: "asdasd",
				Salt:     "",
				Status:   0,
				Avatar:   "",
				Balance:  9999999,
				IsAdmin:  true,
				Mail:     "282119@qq.com",
			}},
			want: &model.User{
				Nickname: "admin",
				Username: "zhangsan",
				Phone:    "15112312233",
				Password: "asdasd",
				Salt:     "",
				Status:   0,
				Avatar:   "",
				Balance:  9999999,
				IsAdmin:  true,
				Mail:     "282119@qq.com",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.CreateUser(tt.args.user)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateUser() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetRegisterUserCount(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	tests := []struct {
		name    string
		fields  fields
		wantRes int64
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test user",
			fields: fields{
				db:  gorm2.GetGormDB(),
				ums: NewUserMetricsRepo(gorm2.GetGormDB()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			gotRes, err := u.GetRegisterUserCount()
			fmt.Sprintln(gotRes)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetRegisterUserCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotRes != tt.wantRes {
				t.Errorf("GetRegisterUserCount() gotRes = %v, want %v", gotRes, tt.wantRes)
			}
		})
	}
}

func TestUserManageRepo_GetUserByID(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		id uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test user",
			fields: fields{
				db:  gorm2.GetGormDB(),
				ums: NewUserMetricsRepo(gorm2.GetGormDB()),
			},
			args: args{id: 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUserByID(tt.args.id)
			fmt.Println(got)
			fmt.Println(got.ID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetUserByKeyword(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		keyword string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "test user",
			fields: fields{
				db:  gorm2.GetGormDB(),
				ums: NewUserMetricsRepo(gorm2.GetGormDB()),
			},
			args: args{keyword: "zhang"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUserByKeyword(tt.args.keyword)
			fmt.Println(got[0])
			fmt.Println(got[1])
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByKeyword() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByKeyword() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetUserByMail(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		mail string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUserByMail(tt.args.mail)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByMail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByMail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetUserByPhone(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		phone string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUserByPhone(tt.args.phone)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByPhone() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByPhone() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetUserByUsername(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUserByUsername(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserByUsername() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByUsername() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetUserProfileByName(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		username string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserProfile
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUserProfileByName(tt.args.username)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUserProfileByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserProfileByName() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_GetUsersByIDS(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		ids []uint64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*model.User
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			got, err := u.GetUsersByIDS(tt.args.ids)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetUsersByIDS() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUsersByIDS() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserManageRepo_UpdateUser(t *testing.T) {
	type fields struct {
		db  *gorm.DB
		ums repo.IUserMetricsRepo
	}
	type args struct {
		user *model.User
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &UserManageRepo{
				db:  tt.fields.db,
				ums: tt.fields.ums,
			}
			if err := u.UpdateUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UpdateUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
