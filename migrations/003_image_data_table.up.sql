CREATE TABLE IF NOT EXISTS image_data (
    id char(26) PRIMARY KEY,
    organization_id char(26) NOT NULL,
    user_id char(26) NOT NULL,
    weight float NOT NULL,
    height float NOT NULL,
    muscle_weight float NOT NULL,
    fat_weight float NOT NULL,
    fat_percent float NOT NULL,
    body_water float NOT NULL,
    protein float NOT NULL,
    mineral float NOT NULL,
    point int NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
