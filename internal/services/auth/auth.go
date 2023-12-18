package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/kitanoyoru/media-system-service/internal/domain/dtos"
	"github.com/kitanoyoru/media-system-service/internal/domain/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

var jwtKey = []byte("my_secret_key")

var (
	ErrTokenInvalid = errors.New("token is invalid")
)

type Claims struct {
	jwt.RegisteredClaims
	Username string `json:"username"`
}

type AuthService struct {
	db *gorm.DB
}

func (as *AuthService) GetJWTToken(dto *dtos.LoginRequestDTO) (string, error) {
	medicalWorker := models.MedicalWorker{}
	err := as.db.First(&medicalWorker, "username = ?", dto.Username).Error
	if err != nil {
		return "", err
	}

	err = bcrypt.CompareHashAndPassword([]byte(medicalWorker.Password), []byte(dto.Password))
	if err != nil {
		return "", err
	}

	claims := &Claims{
		Username: dto.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(5 * time.Minute)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (as *AuthService) VerifyJWTToken(token string) error {
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		return err
	}

	_, ok := parsedToken.Claims.(*Claims)
	if !ok || !parsedToken.Valid {
		return errors.New("token is invalid")
	}

	return nil
}

func (as *AuthService) RefreshJWTToken(token string) (string, error) {
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (any, error) {
		return jwtKey, nil
	})
	if err != nil {
		return "", err
	}

	if !tkn.Valid {
		return "", ErrTokenInvalid
	}

	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(5 * time.Minute))

	newToken := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
