package service

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JWTService interface {
	GenerateToken(userId uuid.UUID, role string) string
	ValidateToken(token string) (*jwt.Token, error)
	GetDataByToken(token string) (uuid.UUID, string, error)
}

type jwtUserClaim struct {
	UserId uuid.UUID `json:"user_id"`
	Role   string    `json:"role"`
	jwt.RegisteredClaims
}

type jwtService struct {
	secretKey string
	issuer    string
}

func NewJWTService() JWTService {
	return &jwtService{
		secretKey: getSecretKey(),
		issuer:    "tugas-rpl-cuy",
	}
}

func getSecretKey() string {
	secretKey := os.Getenv("JWT_SECRET")
	if secretKey == "" {
		secretKey = "tugas-rpl-cuy"
	}
	return secretKey
}

func (j *jwtService) GenerateToken(userId uuid.UUID, role string) string {
	claims := jwtUserClaim{
		UserId: userId,
		Role:   role,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(12 * time.Hour)),
			Issuer:    j.issuer,
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tx, err := token.SignedString([]byte(j.secretKey))
	if err != nil {
		log.Println(err)
	}
	return tx
}

func (j *jwtService) parseToken(t_ *jwt.Token) (any, error) {
	if _, ok := t_.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method %v", t_.Header["alg"])
	}
	return []byte(j.secretKey), nil
}

func (j *jwtService) ValidateToken(token string) (*jwt.Token, error) {
	return jwt.ParseWithClaims(token, &jwtUserClaim{}, j.parseToken)
}

func (j *jwtService) GetDataByToken(token string) (uuid.UUID, string, error) {
	t_Token, err := j.ValidateToken(token)
	if err != nil {
		return uuid.Nil, "", err
	}

	claims, ok := t_Token.Claims.(*jwtUserClaim)
	if !ok || !t_Token.Valid {
		return uuid.Nil, "", fmt.Errorf("invalid token claims")
	}

	return claims.UserId, claims.Role, nil
}
