import type { Route } from "./+types/phishing-simulation"
import { getPhishingSimulationRun } from "../../api"
import HeaderSection from "@components/(Marketing)/HeaderSection"
import VulnerabilityEducator from "@components/(Marketing)/VulnerabilityEducator"
import { Navigate } from "react-router"

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

  return phishingSimulationRun === undefined ? (
    <Navigate to="/" replace />
  ) : (
    <>
      <HeaderSection
        eyebrow="Lerne wie du handeln solltest"
        title="Gut aufgepasst!"
        paragraph="Keine Sorge - das war nur eine Simulation! Diese Seite zeigt dir, woran du die gefälschte E-Mail hättest erkennen können. Achte das nächste Mal besser auf diese Aspekte."
      />
      <VulnerabilityEducator run={phishingSimulationRun} />
    </>
  );
}