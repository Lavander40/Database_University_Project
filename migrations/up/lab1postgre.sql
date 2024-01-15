
-- in: id of students witch should visit needed lecture (1, 2, 3, 4, 5, 6)
-- in: id of needed lessons (1, 2, 3)

select students.id, 
(Select count(id) From attendances Where stud_id = students.id and sched_id in (SELECT id FROM schedules Where lesson_id in (1,2,3)))
/
(Select count(id) from lessons where id In (1,2,3))::float -- lecture with code phrase id
as score 
from students 
where id in (1,2,3,4,5,6)
order by score asc 
limit 10;
