import { useEffect, useState } from "react";
import { animate, motion } from "framer-motion";

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
      initial={{ opacity: 0, y: -10 }}
      animate={{ opacity: 1, y: 0 }}
      exit={{ opacity: 0, y: -10 }}
      style={{
        position: "absolute",
        bottom: "-75%",
        right: 0,
        backgroundColor: "#742774",
        color: "#fff",
        padding: "6px 10px",
        borderRadius: 6,
        boxShadow: "0 2px 8px rgba(0,0,0,0.15)",
        fontSize: 14,
        zIndex: 100,
        whiteSpace: "nowrap"
      }}
    >
      🎉 +{xp - prevXp} XP (Total: {displayXp})
    </motion.div>
  )
}