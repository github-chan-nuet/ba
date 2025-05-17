import { CheckmarkCircle24Filled, Circle24Regular, PlayCircle24Regular } from "@fluentui/react-icons";
import type { LessonRecord } from "../data/courses";

import CourseProgressStyles from "../styles/CourseProgress.module.scss";
import { Body1, Body2, tokens } from "@fluentui/react-components";

type CourseProgressProps = {
  lessons: LessonRecord[];
  currentLesson: LessonRecord;
  completedLessons: string[];
};

export default function CourseProgress({ lessons, currentLesson, completedLessons }: CourseProgressProps) {
  return (
    <div className={CourseProgressStyles.Container}>
      { lessons.map((lesson, idx, lessons) => {
        const isCurrent = lesson === currentLesson;
        return (
          <div className={CourseProgressStyles.Step} key={idx}>
            <div style={{
              display: 'flex',
              flexDirection: 'column',
              alignItems: 'center',
              gap: '.25rem'
            }}>
              { isCurrent ? (
                <PlayCircle24Regular color="black" />
              ) : completedLessons.includes(lesson.id) ? (
                <CheckmarkCircle24Filled color={tokens.colorStatusSuccessBackground3} />
              ) : (
                <Circle24Regular color="black" />
              ) }
              { idx !== lessons.length - 1 &&
                <div style={{
                  width: 1,
                  flex: 1,
                  backgroundColor: "black"
                }}></div> }
            </div>
            <div>
              <Body2>{lesson.label}</Body2>
              <Body1>{lesson.description}</Body1>
            </div>
          </div>
        )
      }) }
    </div>
  )
}