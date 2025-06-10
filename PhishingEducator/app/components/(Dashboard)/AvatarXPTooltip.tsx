import { useEffect, useState } from "react";
import { animate, motion } from "framer-motion";

import AvatarXPTooltipStyles from './AvatarXPTooltip.module.scss';

type AvatarXPTooltipProps = {
  xp: number;
  prevXp: number;
  visible: boolean;
  onAnimationComplete?: () => void
};

export default function AvatarXPTooltip({
  xp,
  prevXp,
  visible,
  onAnimationComplete,
}: AvatarXPTooltipProps) {
  const [displayXp, setDisplayXp] = useState(prevXp);

  useEffect(() => {
    const controls = animate(prevXp, xp, {
      duration: 1,
      onUpdate(value) {
        setDisplayXp(Math.round(value));
      },
      onComplete() {
        onAnimationComplete?.();
      },
    });

    return () => controls.stop();
  }, [prevXp, xp, onAnimationComplete]);

  if (!visible) return null;

  return (
    <motion.div
      className={AvatarXPTooltipStyles.AvatarXPTooltip}
      initial={{ opacity: 0, y: -10 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -10 }}
    >
      🎉 +{xp - prevXp} XP (Total: {displayXp})
    </motion.div>
  )
}