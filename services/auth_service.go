package services

import (
	"datingapp-api/models"
	"datingapp-api/repositories"
	"github.com/astaxie/beego/logs"
	"golang.org/x/crypto/bcrypt"
	"github.com/juusechec/jwt-beego"
	"fmt"
	"time"
)

// AuthService struct
type AuthService struct {
	userRepo      repositories.IUserRepository
}

// NewAuthService function to create AuthService
func NewAuthService(ur repositories.IUserRepository) *AuthService {
	return &AuthService{
		userRepo: ur,
	}
}

// Register func to register user
func (svc *AuthService) Register(registerReq *models.RegisterRequest) models.Response {
	logs.Info("Start Service Register")

	if (registerReq.Username == "") {
		return models.GetResponse(nil, 400,  "Username is required")
	}

	if (registerReq.Password == "") {
		return models.GetResponse(nil, 400,  "Password is required")
	}

	if (registerReq.Email == "") {
		return models.GetResponse(nil, 400,  "Email is required")
	}

	if (registerReq.Fullname == "") {
    return models.GetResponse(nil, 400,  "Fullname is required")
  }

	if (registerReq.Gender == "") {
    return models.GetResponse(nil, 400,  "Gender is required")
  }

	var user models.User

	user.Email = registerReq.Email
	hashPass, err := GeneratehashPassword(registerReq.Password)
	if err!= nil {
		message := fmt.Sprintf("failed generate hash password %s", registerReq.Password)
    logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 400, message)
	}
	user.Password = hashPass
	user.Username = registerReq.Username

	svc.userRepo.BeginTrans()

	userId, err := svc.userRepo.CreateUser(user)
	if err != nil {
		message := "failed create user"
		logs.Error(message)
		logs.Error(err)
		svc.userRepo.RollbackTrans()
		return models.GetResponse(nil, 503, message)
	}

	user.Id = userId

	var profile models.Profile
	profile.User = &user
	profile.Fullname = registerReq.Fullname
	profile.Gender = registerReq.Gender
	profile.BirthDate = registerReq.BirthDate
	profile.ImgPath = registerReq.ImgPath

	_, err = svc.userRepo.CreateProfile(profile)
	if err != nil {
		message := "failed create profile"
		logs.Error(message)
		logs.Error(err)
		svc.userRepo.RollbackTrans()
		return models.GetResponse(nil, 503, message)
	}

	svc.userRepo.CommitTrans()

	token := CreateToken(user)

	dataResp := models.NewAuthResponse(&user, token)
	return models.GetResponse(dataResp, 200, "success")
}

func GeneratehashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CreateToken(user models.User) string {
	et := jwtbeego.EasyToken{
		Username: user.Username,
		Expires:  time.Now().Unix() + 3600, //Segundos
	}
	tokenString, _ := et.GetToken()

	return tokenString
}

func CheckPasswordHash(password, hash string) bool {
  err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
  return err == nil
}

// Login func to user login
func (svc *AuthService) Login(loginReq *models.LoginRequest) models.Response {
	logs.Info("Start Service Login")

	if (loginReq.Username == "" || loginReq.Password == "") {
		return models.GetResponse(nil, 400, "username or password is empty")
	}

	user, err := svc.userRepo.GetUserByUsername(loginReq.Username)
	if err!= nil {
		message := fmt.Sprintf("failed get user with username %s", loginReq.Username)
    logs.Error(message)
		logs.Error(err)
		return models.GetResponse(nil, 400, message)
	}

	check := CheckPasswordHash(loginReq.Password, user.Password)

	if !check {
		message := "Username or Password is incorrect"
		logs.Error(message)
		return models.GetResponse(nil, 401, message)
	}

	logs.Info("%v", user)

	token := CreateToken(user)
	dataResp := models.NewAuthResponse(&user, token)

	return models.GetResponse(dataResp, 200, "success")
}