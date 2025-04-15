import CoursesStyles from './Courses.module.scss';

function Courses() {
  return (
    <>
      <CourseCard />
    </>
  )
}

function CourseCard() {
  return (
    <div className={CoursesStyles.Courses__card}>
      <div>
        <h3>Beginner's Guide to Figma</h3>
        <p>Learn how to design a beautiful and engaging mobile app with Figma.</p>
      </div>
      <div>

      </div>
    </div>
  )
}

export default Courses