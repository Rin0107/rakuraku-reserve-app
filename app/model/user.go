package model

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// usersテーブル情報
// userIdをprimaryKeyに指定している
type User struct {
	UserId int `gorm:"primaryKey;autoIncrement"`
	Name string
	Email string
	Password string
	UserIcon string
	Role string
	PasswordResetToken string
	CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt
}

// DBからユーザー一覧を取得するためのメソッド
// 削除されていないユーザーを取得（deleted_at IS NULL）
func GetUsers()(users []User){
	result := Db.Where("deleted_at IS NULL").Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// ユーザー詳細を取得するためのメソッド
// userIdが0の時該当ユーザーが存在しないと判断する
func GetUserDetailByUserId(userId int)(User,error){
	user:=User{}
	result:=Db.Select("user_id","name","email","user_icon","role","created_at","updated_at").Where("user_id=?",userId).Find(&user)
	if result.Error!=nil||user.UserId==0 {
		return user,fmt.Errorf("該当のユーザーが存在しません")
	}
	return user,nil
}

// ユーザーを登録するためのメソッド
//　初期パスワードはハッシュ化したpasswordとする
func CreateUsers(name,email,role string,password string)(){
	encryptPw,err:=PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return
	}
	user:=User{}
	user.Name=name
	user.Email=email
	user.Role=role
	user.Password=encryptPw
	
	result:=Db.Select("Name","Email","Role","Password").Create(&user)
	if result.Error != nil {
	panic(result.Error)
	}
	return
}

// メール情報からユーザー情報を取得するためのメソッド
func IsEmail(email string) (users []User){
	result := Db.Where("email=?",email).Find(&users)
	if result.Error != nil {
		panic(result.Error)
	}
	return
}

// emailからユーザー情報を取得するためのメソッド
func GetUserPasswordByEmail(email string) (password string,userId int,role string){
	user:=User{}
	result:=Db.Select("Password","UserId","Role").Where("email=?",email).Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	return user.Password,user.UserId,user.Role
}

// トークンをDBに保存するためのメソッド
func SaveTokenToUser(email string,token string){
	result:=Db.Table("users").Where("email = ?", email).Updates(User{PasswordResetToken: token,UpdatedAt: time.Now()})
	if result.Error != nil {
		panic(result.Error)
	}
}

// パスワードトークンを使ってユーザーIDを取得するためのメソッド
func GetUserIdForPasswordToken(passwordToken string) (int,error){
	user:=User{}
	result := Db.Select("UserId").Where("password_reset_token=?",passwordToken).Find(&user)
	if result.Error != nil {
		panic(result.Error)
	}
	if user.UserId == 0 {
		return 0,fmt.Errorf("該当のユーザーが存在しません")
	}
	return user.UserId,result.Error
}

// ユーザー情報を理論削除するためのメソッド
func DeleteUser(userId int)error{
	// ユーザーが存在するか確認
	user,err:=GetUserDetailByUserId(userId)
	if err != nil||user.UserId==0 {
		return fmt.Errorf("該当のユーザーが存在しません")
	}
	// ユーザーを理論削除
	result:=Db.Table("users").Where("user_id = ?", userId).Update("deleted_at",time.Now())
	fmt.Println(result)
	if result.Error!=nil {
		return result.Error
	}
	return nil
}

/* パスワードをリセットするメソッド
	パスワードをリセットする
	update_atカラムを更新する
	パスワードリセットトークンを削除する */
func ResetPassword(userId int, password string) error{
	// パスワードをハッシュ化する
	encryptPw,err:=PasswordEncrypt(password)
	if err != nil {
		fmt.Println("パスワード暗号化中にエラーが発生しました。：", err)
		return err
	}
	result:=Db.Table("users").Where("user_id = ?", userId).Updates(map[string]any{
		"password": encryptPw,
		"updated_at": time.Now(),
		"password_reset_token": nil,
	})
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// ユーザー情報を変更するためのメソッド
func UpdateUserInformation(userId int,name,email,userIcon string) error{
	// ユーザーが存在するか確認
	user,err:=GetUserDetailByUserId(userId)
	if err != nil||user.UserId==0 {
		return fmt.Errorf("該当のユーザーが存在しません")
	}

	// ユーザー情報を変更する処理
	result:=Db.Table("users").Where("user_id = ?", userId).Updates(User{Name: name,Email: email,UserIcon: userIcon,UpdatedAt: time.Now()})
	if result.Error != nil {
		panic(result.Error)
	}
	return nil
}

// 暗号(Hash)化するためのメソッド
func PasswordEncrypt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hash), err
}