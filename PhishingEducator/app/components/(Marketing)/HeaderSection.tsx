import { tokens } from '@fluentui/react-components';
import HeaderSectionStyles from './HeaderSection.module.scss';

export default function HeaderSection() {
  return (
    <section className={HeaderSectionStyles.HeaderSection}>
      <p
        className={HeaderSectionStyles.HeaderSection__Eyebrow}
        style={{
          color: tokens.colorBrandForeground1
        }}
      >
        Lerne wie du handeln solltest
      </p>
      <h2 className={HeaderSectionStyles.HeaderSection__Title}>Gut aufgepasst!</h2>
      <p className={HeaderSectionStyles.HeaderSection__Paragraph}>
        Keine Sorge - das war nur eine Simulation! Diese Seite zeigt dir, woran du die gefälschte E-Mail hättest erkennen können. Achte das nächste Mal besser auf diese Aspekte.
      </p>
    </section>
  )
}