import HeaderSectionStyles from './HeaderSection.module.scss';

type HeaderSectionProps = {
  eyebrow?: string;
  title: string;
  paragraph: string;
}

export default function HeaderSection({ eyebrow, title, paragraph }: HeaderSectionProps) {
  return (
    <section className={HeaderSectionStyles.HeaderSection}>
      { eyebrow &&
        <p className={HeaderSectionStyles.HeaderSection__Eyebrow}>
          { eyebrow }
        </p>
      }
      <h2 className={HeaderSectionStyles.HeaderSection__Title}>{ title }</h2>
      <p className={HeaderSectionStyles.HeaderSection__Paragraph}>{ paragraph }</p>
    </section>
  )
}