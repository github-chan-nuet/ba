import { Body1, Button, Card, CardFooter, CardHeader, ProgressBar, Subtitle1 } from "@fluentui/react-components";
import { getExams, type Exam } from "../../../api";
import type { Route } from "./+types";
import { Link } from "react-router";
import { Rocket20Regular } from "@fluentui/react-icons";

export function meta() {
  return [
    { title: 'Securaware - Prüfungen' },
    {
      name: 'description',
      content: 'Stelle dein Wissen mit Prüfungen auf die Probe. Sieh auf einen Blick, was du bestanden hast und was noch vor dir liegt.'
    },
    {
      name: 'keywords',
      content: 'Phishing Test, Prüfungen, Cybersecurity Quiz, Wissen testen, Securaware Prüfung, Online Sicherheit prüfen, Phishing Simulation, Selbsttest, Sicherheitstraining'
    }
  ]
}

export async function clientLoader() {
  const { data, error } = await getExams();
  if (error) {
    console.error(error);
    return { exams: [] };
  }
  return { exams: data };
}

export default function Exams({ loaderData }: Route.ComponentProps) {
  const { exams } = loaderData;

  return (
    <div style={{
      display: 'grid',
      gridTemplateColumns: 'repeat(auto-fill, minmax(350px, 1fr))',
      gap: 16
    }}>
      { exams.map((exam: Exam, idx: number) => (
        <ExamCard
          key={idx}
          exam={exam}
        />
      ))}
    </div>
  )
}

type ExamCardProps = {
  exam: Exam;
};

function ExamCard({
  exam
}: ExamCardProps) {
  const path = `/dashboard/exams/${exam.id}`;

  return (
    <Card size="large">
      <CardHeader
        style={{
          marginBottom: 'auto',
          gridAutoRows: 'min-content auto'
        }}
        header={<Subtitle1>{exam.title}</Subtitle1>}
        description={<Body1>{exam.description}</Body1>}
      />
      <div className="card-revert-padding">
        <ProgressBar value={0} />
      </div>
      <CardFooter
        action={
          <Link to={path}>
            <Button appearance="primary">
              Starten <Rocket20Regular style={{ marginLeft: 8 }} />
            </Button>
          </Link>
        }
      />
    </Card>
  )
}