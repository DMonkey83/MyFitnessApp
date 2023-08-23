-- Add created_at column to User table
ALTER TABLE "User"
ADD COLUMN created_at TIMESTAMP DEFAULT NOW();

-- Add created_at column to UserProfile table
ALTER TABLE UserProfile
ADD COLUMN created_at TIMESTAMP DEFAULT NOW();
