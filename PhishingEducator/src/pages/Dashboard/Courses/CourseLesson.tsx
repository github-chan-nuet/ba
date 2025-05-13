import { useNavigate } from "react-router";
import { CourseDef, LessonDef } from "../../../data/courses";
import { Button, Title1 } from "@fluentui/react-components";
import CourseLessonStyles from "./CourseLesson.module.scss";
import CourseProgress from "./CourseProgress";
import { createLessonCompletion, getLessonCompletionsOfCourseAndUser } from "../../../api";
import useAuth from "../../../auth/useAuth";
import { useEffect, useState } from "react";

type CourseProps = {
  course: CourseDef,
  lesson: LessonDef
};

const CourseLesson = ({ course, lesson }: CourseProps) => {
  const { token } = useAuth();
  const navigate = useNavigate();
  const [completedLessons, setCompletedLessons] = useState<string[]>([]);
  
    useEffect(() => {
      const fetchCompletions = async () => {
        try {
          const result = await getLessonCompletionsOfCourseAndUser({
            headers: {
              Authorization: `Bearer ${token}`
            },
            path: {
              courseId: course.id
            }
          });
          if (result.response.status === 200 && result.data) {
            setCompletedLessons(() => result.data);
          }
        } catch (e) {
          console.error('Failed to fetch completions', e);
        }
      };
      fetchCompletions();
    }, [token, course.id]);

  const currentPath = location.pathname.split('/').pop();
  const currentIndex = course.lessons.findIndex(lesson => lesson.handle === currentPath);

  const handleNextClick = async () => {
    if (currentIndex !== -1) {
      const currentLesson = course.lessons[currentIndex];
      try {
        const resp = await createLessonCompletion({
          path: {
            courseId: course.id,
          },
          body: {
            lessonId: currentLesson.id,
          },
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        if (resp.data) {
          setCompletedLessons(lessonIds => [...lessonIds, currentLesson.id])
        }
      } catch (e) {
        console.error('Failed to create completion', e);
      }
    }

    if (currentIndex !== -1 && currentIndex < course.lessons.length - 1) {
      const nextLesson = course.lessons[currentIndex + 1];
      navigate(`/dashboard/courses/${course.handle}/${nextLesson.handle}`);
    } else {
      navigate(`/dashboard/courses`);
    }
  }

  return (
    <div>
      <Title1>{ course.label }</Title1>
      <div className={CourseLessonStyles.Container}>
        <div style={{
          maxWidth: 900
        }}>
          {lesson.contentElement}
          <div style={{
            marginTop: 24
          }}>
            <Button appearance="primary" onClick={handleNextClick}>Weiter</Button>
          </div>
        </div>
        <CourseProgress lessons={course.lessons} currentLesson={lesson} completedLessons={completedLessons} />
      </div> 
    </div>
  );
}

export default CourseLesson;