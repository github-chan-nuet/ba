import { CheckmarkCircle24Filled, Circle24Regular, PlayCircle24Regular } from "@fluentui/react-icons";
import type { LessonRecord } from "@data/courses";

import CourseProgressStyles from "./CourseProgress.module.scss";
import { Body1, Body2, tokens } from "@fluentui/react-components";

type CourseProgressProps = {
  lessons: LessonRecord[];
  currentLesson: LessonRecord;
  completedLessons: string[];
};

export default function CourseProgress({ lessons, currentLesson, completedLessons }: CourseProgressProps) {
  return (
    <div className={CourseProgressStyles.CourseProgress}>
      { lessons.map((lesson, idx, lessons) => {
        const isCurrent = lesson === currentLesson;
        return (
          <div className={CourseProgressStyles.CourseProgress__Step} key={idx}>
            <div className={CourseProgressStyles.CourseProgress__StepIndicators}>
              { isCurrent ? (
                <PlayCircle24Regular color="black" />
              ) : completedLessons.includes(lesson.id) ? (
                <CheckmarkCircle24Filled color={tokens.colorStatusSuccessBackground3} />
              ) : (
                <Circle24Regular color="black" />
              ) }
              { idx !== lessons.length - 1 &&
                <div className={CourseProgressStyles.CourseProgress__Line}></div> }
            </div>
            <div className={CourseProgressStyles.CourseProgress__StepTexts}>
              <Body2>{lesson.label}</Body2>
              <Body1>{lesson.description}</Body1>
            </div>
          </div>
        )
      }) }
    </div>
  )
}