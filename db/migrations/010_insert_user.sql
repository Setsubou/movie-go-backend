-- Initialize default admin user with bcrypt password hash of 12 for testing purposes

INSERT INTO users (user_name, password_hash) VALUES 
    ('admin', '$2a$12$E9tK2mhW1xn.FavhPR2uQ.q02pB83UUDlnLTevWp5Bn0GrTTbxFjC')
