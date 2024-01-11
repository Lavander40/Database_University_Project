INSERT INTO equipments (equip) VALUES ('Лабораторное оборудование');
INSERT INTO equipments (equip) VALUES ('Практическое оборудование');
INSERT INTO equipments (equip) VALUES ('Лекционное оборудование');

INSERT INTO rooms (number) VALUES ('301а');
INSERT INTO rooms (number) VALUES ('104');
INSERT INTO rooms (number) VALUES ('263б');

INSERT INTO courses(naming, description) VALUES ('биоинженерия', 'описание предмета?');

INSERT INTO specialities (naming, code, depart_id) VALUES ('Информационные системы и технологии', '9.1.3', 1);
INSERT INTO specialities (naming, code, depart_id) VALUES ('Разработка программных решений', '9.1.4', 1);
INSERT INTO specialities (naming, code, depart_id) VALUES ('Программирование на графических процессорах', '9.1.5', 1);

INSERT INTO groups (depart_id, spec_id) VALUES (1, 1);
INSERT INTO groups (depart_id, spec_id) VALUES (2, 1);
INSERT INTO groups (depart_id, spec_id) VALUES (1, 2);
INSERT INTO groups (depart_id, spec_id) VALUES (2, 2);

INSERT INTO students (full_name, group_id) VALUES ('Vladimir', 1);
INSERT INTO students (full_name, group_id) VALUES ('Vladislav', 1);
INSERT INTO students (full_name, group_id) VALUES ('Victor', 1);
INSERT INTO students (full_name, group_id) VALUES ('Vyacheslav', 1);
INSERT INTO students (full_name, group_id) VALUES ('Valeriy', 1);
INSERT INTO students (full_name, group_id) VALUES ('Viniamin', 1);

INSERT INTO students (full_name, group_id) VALUES ('Dmitriy', 2);
INSERT INTO students (full_name, group_id) VALUES ('Denis', 2);
INSERT INTO students (full_name, group_id) VALUES ('David', 2);
INSERT INTO students (full_name, group_id) VALUES ('Dementiy', 2);
INSERT INTO students (full_name, group_id) VALUES ('Sergey', 2);
INSERT INTO students (full_name, group_id) VALUES ('Nikolay', 2);

INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', NOW(), 1, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', NOW(), 2, 3, 1);
INSERT INTO lessons (typing, date, lection_id, equip_id, course_id) VALUES ('лекция', NOW(), 3, 3, 1);

INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 1, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 2, 1);
INSERT INTO schedules (group_id, lesson_id, room_id) VALUES (1, 3, 1);

INSERT INTO attendances (stud_id, sched_id) VALUES (1, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 1);
INSERT INTO attendances (stud_id, sched_id) VALUES (3, 1);

INSERT INTO attendances (stud_id, sched_id) VALUES (1, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (2, 2);
INSERT INTO attendances (stud_id, sched_id) VALUES (4, 2);
					   
INSERT INTO attendances (stud_id, sched_id) VALUES (1, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (4, 3);
INSERT INTO attendances (stud_id, sched_id) VALUES (5, 3);
					   