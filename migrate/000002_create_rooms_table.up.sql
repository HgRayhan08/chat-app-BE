CREATE TABLE rooms (
   id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
   title VARCHAR(255) NOT NULL,
   created_at TIMESTAMP DEFAULT NOW()
)