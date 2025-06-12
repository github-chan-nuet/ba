import type { Route } from "./+types/lesson";
import { useNavigate } from "react-router";
import { Helmet } from "react-helmet-async";
import { Button, Title1 } from "@fluentui/react-components";
import useAuth from "@utils/auth/useAuth";
import { createLessonCompletion, getLessonCompletionsOfCourseAndUser } from "@api/index";
import { getCourse } from "@data/courses";
import CourseProgress from "@components/(Dashboard)/CourseProgress";

import CourseLessonStyles from '@styles/CourseLesson.module.scss';

export const handle = async ({ params }: Route.LoaderArgs) => {
  const course = await getCourse(params.courseHandle);
  if (course) {
    return course.label;
  }
}

export async function clientLoader({ params }: Route.ClientLoaderArgs) {
  const course = await getCourse(params.courseHandle);
  if (course) {
    const lesson = course.lessons.find(lesson => lesson.handle === params.lessonHandle);
    const { data: completedLessons } = await getLessonCompletionsOfCourseAndUser({
      path: {
        courseId: course.id
      }
    });
    return { course, lesson, completedLessons: completedLessons ?? [] };
  }
  return { course, lesson: null, completedLessons: [] };
}

export default function CourseLesson({ loaderData }: Route.ComponentProps) {
  const { onExperienceGain } = useAuth();
  const navigate = useNavigate();

  const { course, lesson, completedLessons } = loaderData;
  if (!course || !lesson) {
    throw new Response("Not Found", { status: 404 });
  }

  const currentPath = location.pathname.split('/').pop();
  const currentIndex = course.lessons.findIndex(lesson => lesson.handle === currentPath);

  const handleNextClick = async () => {
    if (currentIndex !== -1) {
      const currentLesson = course.lessons[currentIndex];
      const resp = await createLessonCompletion({
        path: {
          courseId: course.id,
        },
        body: {
          lessonId: currentLesson.id,
        }
      });
      if (resp.data && resp.response.status === 201) {
        onExperienceGain(resp.data.newExperienceGained, resp.data.newLevel);
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
    <>
      <Helmet>
        <title>Securaware - { course.label } Kurs</title>
      </Helmet>
      <Title1>{ course.label }</Title1>
      <div className={CourseLessonStyles.Container}>
        <div className={CourseLessonStyles.Content}>
          {lesson.contentElement}
          <div className={CourseLessonStyles.Actions}>
            <Button appearance="primary" onClick={handleNextClick}>Weiter</Button>
          </div>
        </div>
        <CourseProgress lessons={course.lessons} currentLesson={lesson} completedLessons={completedLessons} />
      </div>
    </>
  );
}