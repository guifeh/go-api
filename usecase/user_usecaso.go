package usecase

import (
	"errors"
	"go-api/config"
	"go-api/model"
	"go-api/repository"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type UserUseCase struct {
	userRepository repository.UserRepository
}

func NewUserUseCase(userRepository repository.UserRepository) UserUseCase {
	return UserUseCase{
		userRepository: userRepository,
	}
}

func (u *UserUseCase) getJWTSecret() []byte {
	return config.GetJWTSecret()
}

func (u *UserUseCase) CreateUser(request model.RegisterRequest) (*model.User, error) {
	existingUser, err := u.userRepository.GetUserByEmail(request.Email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := model.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: string(hashedPassword),
	}
	userID, err := u.userRepository.CreateUser(user)
	if err != nil {
		return nil, err
	}
	user.ID = userID
	user.Password = ""
	return &user, nil
}

func (u *UserUseCase) LoginUser(credentials model.Credentials) (string, error) {
	user, err := u.userRepository.GetUserByEmail(credentials.Email)
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	token, err := u.generateJWT(user.ID, user.Email)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (u *UserUseCase) generateJWT(userID int, email string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"email":   email,
		"exp":     time.Now().Add(24 * time.Hour).Unix(),
		"iat":     time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(u.getJWTSecret())
}
