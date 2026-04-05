package auth_service

import (
	"context"
	"errors"
	"task_manager/internal/dto"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

type UserRepo interface {
	CreateUser(ctx context.Context, user *dto.User) error
	GetUserByEmail(ctx context.Context, email string) (*dto.User, error)
}

type RoleRepo interface {
	GetRoleByName(ctx context.Context, roleName string) (*dto.Role, error)
}
type Service struct {
	userRepo  UserRepo
	jwtSecret string
	roleRepo  RoleRepo
	logger    *zap.Logger
}

func New(userRepo UserRepo, roleRepo RoleRepo, jwtSecret string, logger *zap.Logger) *Service {
	return &Service{
		userRepo:  userRepo,
		jwtSecret: jwtSecret,
		roleRepo:  roleRepo,
		logger:    logger,
	}
}

func (s *Service) Register(ctx context.Context, email, password, fullName string) error {
	if email == "" || password == "" {
		return errors.New("email and password required")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		s.logger.Warn("Failed to hash password", zap.Error(err))
		return err
	}

	role, err := s.roleRepo.GetRoleByName(ctx, "user")
	if err != nil {
		s.logger.Warn("Failed to get role", zap.Error(err))
		return err
	}

	user := &dto.User{
		Email:        email,
		FullName:     fullName,
		PasswordHash: string(hash),
		IsActive:     true,
		RoleID:       role.ID,
	}

	return s.userRepo.CreateUser(ctx, user)
}

func (s *Service) Login(ctx context.Context, email, password string) (*dto.TokenPair, error) {
	user, err := s.userRepo.GetUserByEmail(ctx, email)
	if err != nil {
		s.logger.Warn("Failed to get user by email", zap.Error(err))
		return nil, errors.New("invalid credentials")
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return nil, errors.New("invalid credentials")
	}

	accessToken, err := s.generateToken(user.ID, 15*time.Minute)
	if err != nil {
		s.logger.Warn("Failed to generate token", zap.Error(err))
		return nil, err
	}

	refreshToken, err := s.generateToken(user.ID, 7*24*time.Hour)
	if err != nil {
		s.logger.Warn("Failed to generate token", zap.Error(err))
		return nil, err
	}

	return &dto.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (s *Service) Refresh(ctx context.Context, refreshToken string) (*dto.TokenPair, error) {
	token, err := jwt.Parse(refreshToken, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid claims")
	}

	userIDFloat, ok := claims["user_id"].(float64)
	if !ok {
		return nil, errors.New("invalid user_id type")
	}

	userID := int64(userIDFloat)

	accessToken, err := s.generateToken(userID, 15*time.Minute)
	if err != nil {
		s.logger.Warn("Failed to generate accessToken", zap.Error(err))
		return nil, err
	}

	newRefreshToken, err := s.generateToken(userID, 7*24*time.Hour)
	if err != nil {
		s.logger.Warn("Failed to generate refreshToken", zap.Error(err))
		return nil, err
	}

	return &dto.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (s *Service) generateToken(userID int64, ttl time.Duration) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(ttl).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(s.jwtSecret))
}

func (s *Service) parseToken(tokenStr string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		return []byte(s.jwtSecret), nil
	})
	if err != nil || !token.Valid {
		s.logger.Warn("Failed to parse token", zap.Error(err))
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("invalid token claims")
	}

	return claims, nil
}
