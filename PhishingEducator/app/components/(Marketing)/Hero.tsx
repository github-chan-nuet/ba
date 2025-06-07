import { Display } from '@fluentui/react-components'
import HeroStyles from './Hero.module.scss'
import type { ReactElement } from 'react';

type HeroProps = {
  title: string;
  display: ReactElement;
};

export default function Hero({ title, display }: HeroProps) {
  return (
    <section className={HeroStyles.Hero}>
      <div className={HeroStyles.ShinyBackground}>
        <div className={HeroStyles.ShinyBackground__Shine} />
      </div>
      <h1 className={HeroStyles.Hero__Title}>{ title }</h1>
      <hr className={HeroStyles.Hero__Line} />
      <Display className={HeroStyles.Hero__Display}>{ display }</Display>
    </section>
  )
}