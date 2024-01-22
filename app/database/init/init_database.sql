-- usersテーブル
CREATE TABLE IF NOT EXISTS users(
    user_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    user_icon VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP
);

-- eventsテーブル
CREATE TABLE IF NOT EXISTS events(
    event_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(1000) NOT NULL,
    event_date TIMESTAMP NOT NULL,
    join_deadline_date TIMESTAMP NOT NULL,
    capacity INTEGER NOT NULL
);

-- report_imageテーブル
CREATE TABLE IF NOT EXISTS report_images(
    report_image_id SERIAL PRIMARY KEY,
    report_img VARCHAR(255) NOT NULL
);

-- equipment_categoriesテーブル
CREATE TABLE IF NOT EXISTS equipment_categories(
    equipment_category_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

-- inquiriesテーブル
CREATE TABLE IF NOT EXISTS inquiries(
    inquiry_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    body VARCHAR(1000) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id)
);

-- event_join_usersテーブル
CREATE TABLE IF NOT EXISTS event_join_users(
    event_join_user_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    event_id INTEGER NOT NULL,
    deleted_at TIMESTAMP,
    is_owner BOOLEAN NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (event_id) REFERENCES events(event_id)
);

-- announcementsテーブル
CREATE TABLE IF NOT EXISTS announcements(
    announcement_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(1000) NOT NULL,
    publishd_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    user_id INTEGER NOT NULL
);

-- equipmentsテーブル
CREATE TABLE IF NOT EXISTS equipments(
    equipment_id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    explanation VARCHAR(1000) NOT NULL,
    equipment_category_id INTEGER NOT NULL,
    equipment_img VARCHAR(255) NOT NULL,
    is_available BOOLEAN NOT NULL,
    FOREIGN KEY (equipment_category_id) REFERENCES equipment_categories(equipment_category_id)
);

-- reportsテーブル
CREATE TABLE IF NOT EXISTS reports(
    report_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    body VARCHAR(1000) NOT NULL,
    report_image_id INTEGER NOT NULL,
    report_date TIMESTAMP NOT NULL,
    FOREIGN KEY (report_image_id) REFERENCES report_images(report_image_id)
);

-- equipment_reservationsテーブル
CREATE TABLE IF NOT EXISTS equipment_reservations(
    equipment_reservation_id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    equipment_id INTEGER NOT NULL,
    reservation_start_time TIMESTAMP NOT NULL,
    reservation_end_time TIMESTAMP NOT NULL,
    activity_start_time TIMESTAMP NOT NULL,
    activity_end_time TIMESTAMP NOT NULL,
    deleted_at TIMESTAMP,
    report_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (equipment_id) REFERENCES equipments(equipment_id),
    FOREIGN KEY (report_id) REFERENCES reports(report_id)
);

-- usersテーブルにデータを挿入
INSERT INTO users (name, email, password, user_icon) VALUES
('John Doe', 'john@example.com', 'hashed_password', 'user_icon.jpg'),
('Jane Doe', 'jane@example.com', 'hashed_password', 'user_icon.jpg');
-- usersテーブルのデータを確認
SELECT * FROM users;

-- eventsテーブルにデータを挿入
INSERT INTO events (title, body, event_date, join_deadline_date, capacity) VALUES
('Event 1', 'Description for Event 1', '2024-01-19 12:00:00', '2024-01-18 12:00:00', 50),
('Event 2', 'Description for Event 2', '2024-02-01 14:00:00', '2024-01-30 14:00:00', 30);
-- eventsテーブルのデータを確認
SELECT * FROM events;

-- report_imagesテーブルにデータを挿入
INSERT INTO report_images (report_img) VALUES
('image1.jpg'),
('image2.jpg');
-- report_imagesテーブルのデータを確認
SELECT * FROM report_images;

-- equipment_categoriesテーブルにデータを挿入
INSERT INTO equipment_categories (name) VALUES
('Category 1'),
('Category 2');
-- equipment_categoriesテーブルのデータを確認
SELECT * FROM equipment_categories;

-- inquiriesテーブルにデータを挿入
INSERT INTO inquiries (user_id, body) VALUES
(1, 'Inquiry body from user 1'),
(2, 'Inquiry body from user 2');
-- inquiriesテーブルのデータを確認
SELECT * FROM inquiries;

-- event_join_usersテーブルにデータを挿入
INSERT INTO event_join_users (user_id, event_id, deleted_at, is_owner) VALUES
(1, 1, NULL, true),
(2, 1, NULL, false);
-- event_join_usersテーブルのデータを確認
SELECT * FROM event_join_users;

-- announcementsテーブルにデータを挿入
INSERT INTO announcements (title, body, user_id) VALUES
('Announcement 1', 'Announcement body 1', 1),
('Announcement 2', 'Announcement body 2', 2);
-- announcementsテーブルのデータを確認
SELECT * FROM announcements;

-- equipmentsテーブルにデータを挿入
INSERT INTO equipments (name, explanation, equipment_category_id, equipment_img, is_available) VALUES
('Equipment 1', 'Description for Equipment 1', 1, 'equipment_img1.jpg', true),
('Equipment 2', 'Description for Equipment 2', 2, 'equipment_img2.jpg', true),
('Equipment 3', 'Description for Equipment 3', 2, 'equipment_img2.jpg', false);
-- equipmentsテーブルのデータを確認
SELECT * FROM equipments;

-- reportsテーブルにデータを挿入
INSERT INTO reports (title, body, report_image_id, report_date) VALUES
('Report 1', 'Report body 1', 1, '2024-01-20 10:00:00'),
('Report 2', 'Report body 2', 2, '2024-02-05 15:30:00');
-- reportsテーブルのデータを確認
SELECT * FROM reports;

-- equipment_reservationsテーブルにデータを挿入
INSERT INTO equipment_reservations (user_id, equipment_id, reservation_start_time, reservation_end_time, 
    activity_start_time, activity_end_time, deleted_at, report_id) VALUES
(1, 1, '2024-01-21 09:00:00', '2024-01-21 12:00:00', '2024-01-21 10:00:00', '2024-01-21 11:30:00', NULL, 1),
(2, 2, '2024-02-08 14:00:00', '2024-02-08 17:00:00', '2024-02-08 15:00:00', '2024-02-08 16:30:00', NULL, 2);
-- equipment_reservationsテーブルのデータを確認
SELECT * FROM equipment_reservations;
