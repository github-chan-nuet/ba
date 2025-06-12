import { useState } from "react";
import useAuth from "@utils/auth/useAuth";
import { Button, ProgressBar } from "@fluentui/react-components";
import { SettingsCogMultiple24Regular } from "@fluentui/react-icons";
import ProfileDrawer from "@components/(Dashboard)/ProfileDrawer";

import UserPopoverStyles from './UserPopover.module.scss';

export default function UserPopover() {
  const { user } = useAuth();

  const [isProfileDrawerOpen, setIsProfileDrawerOpen] = useState<boolean>(false);

  return (
    <div className={UserPopoverStyles.UserPopover}>
      <ProfileDrawer
        isOpen={isProfileDrawerOpen}
        setIsOpen={setIsProfileDrawerOpen}
      />

      <div className={UserPopoverStyles.UserPopover__Content}>
        <span className={UserPopoverStyles.UserPopover__CurrentLevel}>
          { user?.level }
        </span>
        <div className={UserPopoverStyles.UserPopover__LevelProgress}>
          <ProgressBar
            value={user?.totalExperience}
            max={user?.totalXpForNextLevel}
          />
          <span className={UserPopoverStyles.UserPopover__LevelProgressLabel}>
            Noch { (user?.totalXpForNextLevel ?? 0) - (user?.totalExperience ?? 0) } XP bis zum nächsten Level
          </span>
        </div>
        Level { (user?.level ?? 0) + 1 }
      </div>

      <Button
        onClick={() => setIsProfileDrawerOpen(true)}
      >
        <SettingsCogMultiple24Regular /> Einstellungen
      </Button>
    </div>
  );
}