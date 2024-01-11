CREATE TABLE equipments(
    id SERIAL PRIMARY KEY,
    equip VARCHAR(60) NOT NULL
);

CREATE TABLE rooms(
    id SERIAL PRIMARY KEY,
    number VARCHAR(6) NOT NULL
);

CREATE TABLE courses(
    id SERIAL PRIMARY KEY,
    naming VARCHAR(60) NOT NULL,
    description VARCHAR(60)
);

CREATE TABLE specialities(
    id SERIAL PRIMARY KEY,
    naming VARCHAR(60) NOT NULL,
    code VARCHAR(10),
    depart_id INT NOT NULL
);

CREATE TABLE groups(
    id SERIAL PRIMARY KEY,
    depart_id INT NOT NULL,
    spec_id INT NOT NULL,
    FOREIGN KEY (spec_id) REFERENCES specialities(id)
);

CREATE TABLE students(
    id SERIAL PRIMARY KEY,
    full_name VARCHAR(60) NOT NULL,
    group_id INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id)
);

CREATE TABLE lessons(
    id SERIAL PRIMARY KEY,
    typing varchar(30) NOT NULL,
    date TIMESTAMP NOT NULL,
    lection_id INT NOT NULL,
    equip_id INT NOT NULL,
    course_id INT NOT NULL,
    FOREIGN KEY (equip_id) REFERENCES equipments(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);

CREATE TABLE schedules(
    id SERIAL PRIMARY KEY,
    group_id INT NOT NULL,
    lesson_id INT NOT NULL,
    room_id INT NOT NULL,
    FOREIGN KEY (group_id) REFERENCES groups(id),
    FOREIGN KEY (lesson_id) REFERENCES lessons(id),
    FOREIGN KEY (room_id) REFERENCES rooms(id)
);

CREATE TABLE attendances(
    id SERIAL PRIMARY KEY,
    stud_id INT NOT NULL,
    sched_id INT NOT NULL,
    FOREIGN KEY (stud_id) REFERENCES students(id),
    FOREIGN KEY (sched_id) REFERENCES schedules(id)
);

CREATE TABLE specialities_courses(
    id SERIAL PRIMARY KEY,
    spec_id INT NOT NULL,
    course_id INT NOT NULL,
    FOREIGN KEY (spec_id) REFERENCES specialities(id),
    FOREIGN KEY (course_id) REFERENCES courses(id)
);
