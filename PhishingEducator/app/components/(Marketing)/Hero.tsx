import type { ReactElement } from 'react';
import { motion } from 'framer-motion';
import { ChevronCircleDown48Regular } from '@fluentui/react-icons';

import HeroStyles from './Hero.module.scss';

type HeroProps = {
  title: ReactElement;
  subtitle?: ReactElement;
  showScrollIcon?: boolean;
};

export default function Hero({ title, subtitle, showScrollIcon = false }: HeroProps) {
  return (
    <section className={HeroStyles.Hero}>
      <div className={HeroStyles.ShinyBackground}>
        <div className={HeroStyles.ShinyBackground__Shine} />
      </div>
      <h1 className={HeroStyles.Hero__Title}>{ title }</h1>
      { subtitle &&
        <p className={HeroStyles.Hero__Subtitle}>{ subtitle }</p>
      }
      { showScrollIcon &&
        <div className={HeroStyles.Hero__Icon}>
          <motion.div
            animate={{ y: [0, 10, 0] }}
            transition={{
              duration: 1.5,
              repeat: Infinity,
              ease: "easeInOut"
            }}
          >
            <ChevronCircleDown48Regular className="color-gray-500" />
          </motion.div>
        </div>
      }
    </section>
  )
}