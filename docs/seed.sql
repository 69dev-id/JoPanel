-- Default Admin User Seed
-- Username: admin
-- Password: password123

INSERT INTO users (username, email, password_hash, role, status, created_at, updated_at)
VALUES (
    'admin', 
    'admin@example.com', 
    '$2a$14$VpJsND4YGZL6vxsHhijzR.UgLkepU3O46n93udh90pBc6xhcdWL1S', 
    'admin', 
    'active', 
    NOW(), 
    NOW()
);
