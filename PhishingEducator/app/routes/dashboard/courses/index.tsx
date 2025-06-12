import type { Route } from "./+types";
import { Link } from "react-router";
import { Body1, Button, Card, CardFooter, CardHeader, ProgressBar, Subtitle1 } from "@fluentui/react-components";
import { ArrowRight20Regular, Rocket20Regular } from "@fluentui/react-icons";
import { getCourses, type CourseRecord } from "@data/courses";
import { getAllLessonCompletionsOfUser } from "@api/index";

export function meta() {
  return [
    { title: 'Securaware - Online-Kurse' },
    {
      name: 'description',
      content: 'Finde den passenden Kurs, um Phishing zu erkennen und dich online zu schützen. Lerne in deinem Tempo und werde Schritt für Schritt sicherer im Netz.'
    },
    {
      name: 'keywords',
      content: 'Phishing Kurse, Online Sicherheit lernen, Securaware, Betrug erkennen, Cybersecurity Training, Phishing Schulung, Internetbetrug, Selbstschutz online, Kursübersicht'
    }
  ]
}

export async function clientLoader() {
  const courses = await getCourses();
  const { data: completions, error } = await getAllLessonCompletionsOfUser();
  if (error) {
    return { courses, completions: [] };
  }
  return { courses, completions };
}

export default function Courses({ loaderData }: Route.ComponentProps) {
  const { courses, completions } = loaderData;

  return (
    <div style={{
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fill, minmax(350px, 1fr))',
      gap: 16
    }}>
      { courses.map((course: CourseRecord, idx: number) => (
        <CourseCard
          key={idx}
          course={course}
          completedLessonIds={completions.find(c => c.courseId === course.id)?.completedLessons ?? []}
        />
      )) }
    </div>
  )
}

type CourseCardProps = {
  course: CourseRecord;
  completedLessonIds: string[];
};

function CourseCard({
  course,
  completedLessonIds,
}: CourseCardProps) {
  const nonCompletedLessons = course.lessons.filter(lesson => !completedLessonIds.includes(lesson.id));
  const firstNonCompletedLesson = nonCompletedLessons.at(0);
  const completedPercentage = ((course.lessons.length - nonCompletedLessons.length) / course.lessons.length) || 0;

  let path: string;
  if (firstNonCompletedLesson) {
    path = `/dashboard/courses/${course.handle}/${firstNonCompletedLesson.handle}`;
  } else {
    path = `/dashboard/courses/${course.handle}/${course.lessons.length > 0 ? course.lessons[0].handle : ''}`;
  }

  return (
    <Card size="large">
      <CardHeader
        style={{
          marginBottom: 'auto',
          gridAutoRows: 'min-content auto'
        }}
        header={<Subtitle1>{course.label}</Subtitle1>}
        description={<Body1>{course.description}</Body1>}
      />
      <div className="card-revert-padding">
        <ProgressBar value={completedPercentage} />
      </div>
      <CardFooter
        action={
          <Link to={path}>
            <Button appearance="primary">
              { completedPercentage <= 0 ? (
                <>
                  Starten <Rocket20Regular style={{ marginLeft: 8 }} />
                </>
              ) : completedPercentage >= 1 ? (
                <>
                  Neu starten <Rocket20Regular style={{ marginLeft: 8 }} />
                </>
              ) : (
                <>
                  Fortfahren <ArrowRight20Regular style={{ marginLeft: 8 }} />
                </>
              )}
            </Button>
          </Link>
        }
      />
    </Card>
  )
}