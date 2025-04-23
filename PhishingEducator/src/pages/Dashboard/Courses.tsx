import { Subtitle2Stronger, Card, CardHeader, ProgressBar, CardFooter, Button, Body1 } from '@fluentui/react-components';

function Courses() {
  return (
    <div style={{
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fill, minmax(350px, 1fr))',
      gap: 16
    }}>
      <CourseCard
        title="Common Ground"
        description="Dieser Kurs bietet einen kompakten Einstieg in das Thema Phishing. Du erfährst, was Phishing ist, wie verbreitet es ist, welche Schäden entstehen können und warum es für Securaware relevant ist."
        completedPercentage={0.25}
      />
      <CourseCard
        title="Angriffsvektoren"
        description="Lerne, über welche Wege Phishing-Angriffe verbreitet werden - von E-Mail über SMS bis hin zu Telefonanrufen. Dieser Kurs zeigt dir die typischen Einfallstore für Angreifer."
        completedPercentage={0}
      />
      <CourseCard
        title="Sensitive Informationen"
        description="Erfahre, welche persönlichen Daten besonders schützenswert sind und warum Phishing-Angriffe genau auf sie abzielen."
        completedPercentage={0}
      />
      <CourseCard
        title="URL-Spoofing"
        description="In diesem Kurs lernst du, wie manipulierte Links dich in die Falle locken - inklusive Techniken wie URL-Verkürzung oder homographische Angriffe."
        completedPercentage={0}
      />
      <CourseCard
        title="Indizien"
        description="Lerne, woran du Phishing-Versuche erkennst: von verdächtigen Absendern und Anhängen bis hin zu sprachlichen Auffälligkeiten und untypischem Kontext."
        completedPercentage={0}
      />
      <CourseCard
        title="Tools gegen Phishing"
        description="Entdecke hilfreiche Tools und Techniken zur Abwehr von Phishing - wie Multi-Faktor-Authentifizierung, Browser-Erweiterungen und Domain-Checker."
        completedPercentage={0}
      />
    </div>
  )
}

type CourseCardProps = {
  title: string,
  description: string,
  completedPercentage: number
};

function CourseCard({
  title,
  description,
  completedPercentage
}: CourseCardProps) {
  return (
    <Card size="large">
      <CardHeader
        style={{
          flexGrow: 1,
          gridAutoRows: 'min-content auto'
        }}
        header={<Subtitle2Stronger>{title}</Subtitle2Stronger>}
        description={<Body1>{description}</Body1>}
      />
      <div className="card-revert-padding">
        <ProgressBar value={completedPercentage}  />
      </div>
      <CardFooter
        action={
          completedPercentage > 0 ?
          <Button appearance="primary">Fortfahren</Button> :
          <Button appearance="primary">Starten</Button>
        }
      >
      </CardFooter>
    </Card>
  )
}

export default Courses