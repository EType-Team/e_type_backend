package model

import "time"

type Provider struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	UserId            uint      `json:"user_id"`
	Type              string    `json:"type"`
	Provider          string    `json:"provider"`
	ProviderAccountID string    `json:"provider_account_id"`
	RefreshToken      string    `json:"refresh_token"`
	AccessToken       string    `json:"access_token"`
	ExpiresAt         time.Time `json:"expires_at"`
	TokenType         string    `json:"token_type"`
	Scope             string    `json:"scope"`
	IDToken           string    `json:"id_token"`
	SessionState      string    `json:"session_state"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}

type ProviderResponse struct {
	ID                uint      `json:"id" gorm:"primaryKey"`
	UserID            uint      `json:"user_id"`
	Provider          string    `json:"provider"`
	ProviderAccountID string    `json:"provider_account_id"`
	ExpiresAt         time.Time `json:"expires_at"`
}
