import { Subtitle2Stronger, Card, CardHeader, ProgressBar, CardFooter, Button, Body1 } from '@fluentui/react-components';
import { ArrowRight20Regular, Rocket20Regular } from '@fluentui/react-icons';
import { Link } from 'react-router';
import { courseData, CourseDef } from '../../data/courses';
import useAuth from '../../auth/useAuth';
import { CourseCompletion, getAllLessonCompletionsOfUser } from '../../api';
import { useEffect, useState } from 'react';

function Courses() {
  const { token } = useAuth();
  const [completions, setCompletions] = useState<CourseCompletion[]>([]);

  useEffect(() => {
    const fetchCompletions = async () => {
      try {
        const result = await getAllLessonCompletionsOfUser({
          headers: {
            Authorization: `Bearer ${token}`
          }
        });
        if (result.response.status === 200 && result.data) {
          setCompletions(() => result.data);
        }
      } catch (e) {
        console.error('Failed to fetch completions', e);
      }
    };
    fetchCompletions();
  }, [token]);

  return (
    <div style={{
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fill, minmax(350px, 1fr))',
      gap: 16
    }}>
      { courseData.map((course, idx) => (
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
  course: CourseDef;
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
          flexGrow: 1,
          gridAutoRows: 'min-content auto'
        }}
        header={<Subtitle2Stronger>{course.label}</Subtitle2Stronger>}
        description={<Body1>{course.description}</Body1>}
      />
      <div className="card-revert-padding">
        <ProgressBar value={completedPercentage}  />
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
              ) : 
              (
                <>
                  Fortfahren <ArrowRight20Regular style={{ marginLeft: 8 }} />
                </>
              )
            }
            </Button>
          </Link>
        }
      >
      </CardFooter>
    </Card>
  )
}

export default Courses