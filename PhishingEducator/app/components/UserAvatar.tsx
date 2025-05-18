import { useEffect, useState } from "react";
import { Avatar } from "@fluentui/react-components";
import { AnimatePresence } from "framer-motion";
import AvatarXPTooltip from "./AvatarXPTooltip";
import useAuth from "../utils/auth/useAuth";

export default function UserAvatar() {
  const { user } = useAuth();
  const [prevXp, setPrevXp] = useState<number|undefined>(user?.totalExperience);
  const [showPopup, setShowPopup] = useState(false);

  useEffect(() => {
    if (typeof prevXp === "number" && user?.totalExperience !== undefined && user.totalExperience !== prevXp) {
      setShowPopup(true);
    } else if (user?.totalExperience !== undefined) {
      setPrevXp(user.totalExperience);
    }
  }, [user?.totalExperience, prevXp]);

  return (
    <div style={{
      position: "relative"
    }}>
      <Avatar
        name={`${user?.firstname} ${user?.lastname}`}
      />
      <AnimatePresence>
        { showPopup && (
          <AvatarXPTooltip
            xp={user?.totalExperience ?? 0}
            prevXp={prevXp ?? 0}
            visible={showPopup}
            onAnimationComplete={() => {
              setTimeout(() => setShowPopup(false), 1500);
              setTimeout(() => setPrevXp(user?.totalExperience ?? 0), 1500);
            }}
          />
        ) }
      </AnimatePresence>
    </div>
  )
}