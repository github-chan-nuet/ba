import { Subtitle2Stronger, Card, CardHeader, ProgressBar } from '@fluentui/react-components';

function Courses() {
  return (
    <>
      <CourseCard />
    </>
  )
}

function CourseCard() {
  return (
    <Card size="large">
      <CardHeader
        header={
          <Subtitle2Stronger>
            Beginner's Guide to Figma
          </Subtitle2Stronger>
        }
      />
      <ProgressBar value={0.25}  />
    </Card>
  )
}

export default Courses