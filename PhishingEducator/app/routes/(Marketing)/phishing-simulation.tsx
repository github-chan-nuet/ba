import Hero from "@components/(Marketing)/Hero"
import type { Route } from "./+types/phishing-simulation"
import { getPhishingSimulationRun } from "../../api"

export function meta() {
  return [
    { title: 'Securaware - Wir haben eine Schwachstelle entdeckt!'},
  ]
}

export async function clientLoader({ request }: Route.ClientLoaderArgs) {
  const url = new URL(request.url)
  const runIdQueryParam = url.searchParams.get('r');

  const { data: phishingSimulationRun } = await getPhishingSimulationRun({
    path: {
      phishingSimulationRunId: runIdQueryParam ?? ''
    }
  });
  return { phishingSimulationRun };
}

export default function PhishingSimulationEducation({ loaderData }: Route.ComponentProps) {
  const { phishingSimulationRun } = loaderData;
  console.log(phishingSimulationRun);

  return (
    <Hero
      title="Securaware"
      display={<>Wir haben eine <strong>Schwachstelle</strong> von dir entdeckt!</>}
    />
  )
}