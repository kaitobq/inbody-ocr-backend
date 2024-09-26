package service

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type TokenService struct {
}

func NewTokenService() TokenService {
	return TokenService{}
}


func (ts *TokenService) GenerateTokenFromID(userID, orgID string) (string, error) {
	tokenLifeSpanStr := os.Getenv("TOKEN_LIFE_SPAN")
	if(len(tokenLifeSpanStr) == 0) {
		return "", fmt.Errorf("TOKEN_LIFE_SPAN is not set in the environment")
	}
	tokenLifeSpan, err := strconv.Atoi(tokenLifeSpanStr)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"authorized": true,
		"user_id": userID,
		"org_id": orgID,
		"exp": time.Now().Add(time.Hour * time.Duration(tokenLifeSpan)).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}

func (ts *TokenService) TokenValid(c *gin.Context) (bool, error) {
	tokenStr, err := getTokenStringFromRequestHeader(c)
	if err != nil {
		return false, err
	}

	token, err := parseToken(tokenStr)

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, nil
	}

	return true, nil
}

func (ts *TokenService) ExtractIDsFromContext(c *gin.Context) (string, string, error) {
	tokenStr, err := getTokenStringFromRequestHeader(c)
	if err != nil {
		return "", "", err
	}

	token, err := parseToken(tokenStr)
	if err != nil {
		return "", "", err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", "", errors.New("error while parsing claims")
	}

	userID, ok := claims["user_id"].(string)
	if !ok {
		return "", "", errors.New("error while parsing id")
	}

	orgID, ok := claims["org_id"].(string)
	if !ok {
		return "", "", errors.New("error while parsing id")
	}

	return userID, orgID, nil
}

func (ts *TokenService) ExtractExpFromToken(tokenStr string) (*time.Time, error) {
	token, err := parseToken(tokenStr)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("error while parsing claims")
	}

	exp, ok := claims["exp"].(float64)
	if !ok {
		return nil, errors.New("error while parsing exp")
	}

	loc, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		return nil, err
	}

	expTime := time.Unix(int64(exp), 0).In(loc)

	return &expTime, nil
}

func getTokenStringFromRequestHeader(c *gin.Context) (string, error) {
	bearToken := c.Request.Header.Get("Authorization")
    strArr := strings.Split(bearToken, " ")
    if len(strArr) == 2 {
        return strArr[1], nil
    }

    return "", errors.New("no token found")
}

func parseToken(tokenString string) (*jwt.Token, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("there was an error while parsing the token")
        }
        return []byte(os.Getenv("JWT_SECRET")), nil
    })

    if err != nil {
        return nil, err
    }

    return token, nil
}
