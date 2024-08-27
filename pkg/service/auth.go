package service

import (
	"Notes/models"
	"Notes/pkg/repository"
	"crypto/hmac"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	salt       = "jsdjfn392uujuhduh804ufjduu48dn984"
	signingKey = "urjehr84ugvujdnjn94vunevjfi4"
	tokenTTl   = 12 * time.Hour
)

type TokenClaims struct {
	UserId int   `json:"user_id"`
	Exp    int64 `json:"exp"`
	Iat    int64 `json:"iat"`
}

type AuthService struct {
	repo repository.Authorization
}

func NewAuthService(repo repository.Authorization) *AuthService {
	return &AuthService{repo: repo}
}

func (s *AuthService) CreateUser(user models.Users) (int, error) {
	user.Password = generatePasswordHash(user.Password)

	return s.repo.CreateUser(user)
}

func generatePasswordHash(password string) string {
	hash := sha1.New()
	hash.Write([]byte(password))

	return fmt.Sprintf("%x", hash.Sum([]byte(salt)))
}

func (s *AuthService) GenerateToken(email, password string) (string, error) {
	user, err := s.repo.GetUser(email, generatePasswordHash(password))
	if err != nil {
		return "", err
	}

	claims := TokenClaims{
		UserId: user.Id,
		Exp:    time.Now().Add(tokenTTl).Unix(),
		Iat:    time.Now().Unix(),
	}

	payload, err := json.Marshal(claims)
	if err != nil {
		return "", err
	}

	payloadEncoded := base64.RawURLEncoding.EncodeToString(payload)
	header := base64.RawURLEncoding.EncodeToString([]byte(`{"alg":"HS256","typ":"JWT"}`))
	unsignedToken := fmt.Sprintf("%s.%s", header, payloadEncoded)

	signature, err := createSignature(unsignedToken, signingKey)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s.%s", unsignedToken, signature), nil
}

func createSignature(unsignedToken, key string) (string, error) {
	h := hmac.New(sha256.New, []byte(key))
	_, err := h.Write([]byte(unsignedToken))
	if err != nil {
		return "", err
	}

	signature := h.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(signature), nil
}

func (s *AuthService) ParseToken(token string) (int, error) {
	// Разбиваем токен на три части: header, payload, signature
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return 0, fmt.Errorf("invalid token format")
	}

	unsignedToken := fmt.Sprintf("%s.%s", parts[0], parts[1])

	// Валидация подписи
	expectedSignature, err := createSignature(unsignedToken, signingKey)
	if err != nil {
		return 0, err
	}

	if parts[2] != expectedSignature {
		return 0, fmt.Errorf("invalid token signature")
	}

	// Декодируем payload (claims)
	payload, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return 0, err
	}

	var claims TokenClaims
	if err := json.Unmarshal(payload, &claims); err != nil {
		return 0, err
	}

	// Проверка времени истечения токена
	if claims.Exp < time.Now().Unix() {
		return 0, fmt.Errorf("token expired")
	}

	return claims.UserId, nil
}
