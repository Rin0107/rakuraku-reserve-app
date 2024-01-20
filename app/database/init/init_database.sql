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
    deleted_at TIMESTAMP NOT NULL,
    is_orner BOOLEAN NOT NULL,
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
    deleted_at TIMESTAMP NOT NULL,
    report_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users(user_id),
    FOREIGN KEY (equipment_id) REFERENCES equipments(equipment_id),
    FOREIGN KEY (report_id) REFERENCES reports(report_id)
);