import { CheckmarkCircle24Filled, Circle24Regular, PlayCircle24Regular } from "@fluentui/react-icons";
import { LessonDef } from "../../../data/courses";
import CourseProgressStyles from "./CourseProgress.module.scss";
import { Body1, Body2, tokens } from "@fluentui/react-components";

type CourseProgressProps = {
  lessons: LessonDef[];
  currentLesson: LessonDef;
  completedLessons: string[];
}

const CourseProgress = ({ lessons, currentLesson, completedLessons }: CourseProgressProps) => {
  return (
    <div className={CourseProgressStyles.Container}>
      {
        lessons.map((lesson, idx, lessons) => {
          const isCurrent = lesson === currentLesson;
          return (
            <div className={CourseProgressStyles.Step} key={idx}>
              <div style={{
                display: 'flex',
                flexDirection: 'column',
                alignItems: 'center',
                gap: '.25rem'
              }}>
                { /* {tokens.colorStatusSuccessBackground3} */ }
                {isCurrent ? (
                  <PlayCircle24Regular color="black" />
                ) : 
                  completedLessons.includes(lesson.id) ? (
                    <CheckmarkCircle24Filled color={tokens.colorStatusSuccessBackground3} />
                  ) : (
                    <Circle24Regular color="black" />
                  )
                }
                {idx !== lessons.length - 1 ? (
                  <div style={{
                    width: 1,
                    flex: 1,
                    backgroundColor: "black"
                  }}></div>
                ) : null}
              </div>
              <div>
                <Body2>{lesson.label}</Body2><br />
                <Body1>{lesson.description}</Body1>
              </div>
            </div>
          )
        })
      }
    </div>
  );
}

export default CourseProgress;