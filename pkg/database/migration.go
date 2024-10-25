package database

import (
	"context"
	"errors"
	"fmt"
	"inbody-ocr-backend/internal/domain/entity"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/exp/rand"
)

func (db *DB) Migrate() error {
	driver, err := mysql.WithInstance(db.DB.DB, &mysql.Config{})
	if err != nil {
		return err
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file:///migrations",
		"mysql", driver,
	)
	if err != nil {
		return err
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

// SeedDataメソッドをDB構造体に追加
func (db *DB) SeedData() error {
	rand.Seed(uint64(time.Now().UnixNano()))
	ctx := context.Background()

	// 組織のシードデータを生成
	organizationIDs := []string{"org1", "org2"}
	for _, orgID := range organizationIDs {
		organization := entity.Organization{
			ID:        orgID,
			Name:      orgID,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}
		// 組織をデータベースに挿入
		_, err := db.DB.ExecContext(ctx, `
            INSERT INTO organizations (id, name, created_at, updated_at)
            VALUES (?, ?, ?, ?)
        `, organization.ID, organization.Name, organization.CreatedAt, organization.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to insert organization: %w", err)
		}
	}

	// ユーザーのシードデータを生成
	userIDs := []string{"user1", "user2", "user3", "user4", "user5"}
	bytes, _ := bcrypt.GenerateFromPassword([]byte("password"), 14)
	hashedPassword := string(bytes)
	for _, userID := range userIDs {
		user := entity.User{
			ID:             userID,
			Name:           userID,
			Email:          fmt.Sprintf("%s@example.com", userID),
			Password:       hashedPassword,
			OrganizationID: "org1",
			Role:           "member",
			CreatedAt:      time.Now(),
			UpdatedAt:      time.Now(),
		}
		// ユーザーをデータベースに挿入
		_, err := db.DB.ExecContext(ctx, `
            INSERT INTO users (id, name, email, password, organization_id, role, created_at, updated_at)
            VALUES (?, ?, ?, ?, ?, ?, ?, ?)
        `, user.ID, user.Name, user.Email, user.Password, user.OrganizationID, user.Role, user.CreatedAt, user.UpdatedAt)
		if err != nil {
			return fmt.Errorf("failed to insert user: %w", err)
		}

		// 各ユーザーに対して画像データを生成
		numRecords := 10 // 1～3個のレコード
		for i := 0; i < numRecords; i++ {
			weight := rand.Float64()*50 + 50    // 50kg～100kg
			height := rand.Float64()*30 + 150   // 150cm～180cm
			fatPercent := rand.Float64()*25 + 5 // 5%～30%
			muscleWeight := weight * (1 - fatPercent/100) * (rand.Float64()*0.1 + 0.4)

			imageData := entity.ImageData{
				ID:             fmt.Sprintf("%s-%d", userID, i),
				OrganizationID: user.OrganizationID,
				UserID:         user.ID,
				Weight:         weight,
				Height:         height,
				MuscleWeight:   muscleWeight,
				FatWeight:      weight * fatPercent / 100,
				FatPercent:     fatPercent,
				BodyWater:      rand.Float64()*40 + 40, // 40%～80%
				Protein:        rand.Float64()*20 + 10, // 10%～30%
				Mineral:        rand.Float64()*5 + 3,   // 3kg～8kg
				Point:          uint(rand.Intn(100)),
				CreatedAt:      time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
				UpdatedAt:      time.Now().Add(-time.Duration(rand.Intn(1000)) * time.Hour),
			}
			// 画像データをデータベースに挿入
			_, err := db.DB.ExecContext(ctx, `
                INSERT INTO image_data (
                    id, organization_id, user_id, weight, height, muscle_weight, fat_weight, fat_percent,
                    body_water, protein, mineral, point, created_at, updated_at
                ) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
            `, imageData.ID, imageData.OrganizationID, imageData.UserID, imageData.Weight, imageData.Height,
				imageData.MuscleWeight, imageData.FatWeight, imageData.FatPercent, imageData.BodyWater,
				imageData.Protein, imageData.Mineral, imageData.Point, imageData.CreatedAt, imageData.UpdatedAt)
			if err != nil {
				return fmt.Errorf("failed to insert image data: %w", err)
			}
		}
	}
	return nil
}
