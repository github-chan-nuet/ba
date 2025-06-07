import Hero from "@components/(Marketing)/Hero"

export function meta() {
  return [
    { title: 'Securaware - Wir haben eine Schwachstelle entdeckt!'},
  ]
}

export default function PhishingSimulationEducation() {
  return (
    <Hero
      title="Securaware"
      display={<>Wir haben eine <strong>Schwachstelle</strong> von dir entdeckt!</>}
    />
  )
}