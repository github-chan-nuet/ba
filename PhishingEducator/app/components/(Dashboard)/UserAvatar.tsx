import { useEffect, useState } from "react";
import { AnimatePresence } from "framer-motion";
import useAuth from "@utils/auth/useAuth";
import { Avatar, Popover, PopoverSurface, PopoverTrigger } from "@fluentui/react-components";
import AvatarXPTooltip from "@components/(Dashboard)/AvatarXPTooltip";
import UserPopover from "@components/(Dashboard)/UserPopover";

import UserAvatarStyles from './UserAvatar.module.scss'

export default function UserAvatar() {
  const { user } = useAuth();
  const [prevXp, setPrevXp] = useState<number|undefined>(user?.totalExperience);
  const [showPopup, setShowPopup] = useState(false);

  useEffect(() => {
    if (
      typeof prevXp === "number" &&
      user?.totalExperience !== undefined &&
      user.totalExperience !== prevXp
    ) {
      setShowPopup(true);
    } else if (user?.totalExperience !== undefined) {
      setPrevXp(user.totalExperience);
    }
  }, [user?.totalExperience, prevXp]);

  return (
    <div className={UserAvatarStyles.UserAvatar}>
      <Popover withArrow={true} openOnHover={true}>
        <PopoverTrigger>
          <Avatar
            name={`${user?.firstname} ${user?.lastname}`}
          />
        </PopoverTrigger>
        <PopoverSurface>
          <UserPopover />
        </PopoverSurface>
      </Popover>
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