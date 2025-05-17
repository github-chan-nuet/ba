import { useEffect, useState } from "react";
import useAuth from "../../../utils/auth/useAuth";
import { useNavigate } from "react-router";
import { createLessonCompletion, getLessonCompletionsOfCourseAndUser } from "../../../api";
import { getCourse } from "../../../data/courses";
import { Button, Title1 } from "@fluentui/react-components";
import CourseProgress from "../../../components/CourseProgress";
import type { Route } from "./+types/lesson";

import CourseLessonStyles from '../../../styles/CourseLesson.module.scss';

export const handle = async ({ params }: Route.LoaderArgs) => {
  const course = await getCourse(params.courseHandle);
  if (course) {
    return course.label;
  }
}

export async function clientLoader({ params }: Route.LoaderArgs) {
  const course = await getCourse(params.courseHandle);
  if (course) {
    const lesson = course.lessons.find(lesson => lesson.handle === params.lessonHandle);
    return { course, lesson };
  }
  return { course, lesson: null };
}

export default function CourseLesson({ loaderData }: Route.ComponentProps) {
  const { course, lesson } = loaderData;
  if (!course || !lesson) {
    throw new Response("Not Found", { status: 404 });
  }

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
          setCompletedLessons(lessonIds => [...lessonIds, currentLesson.id]);
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