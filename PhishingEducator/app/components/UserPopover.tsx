import { Button, ProgressBar, tokens } from "@fluentui/react-components";
import { SettingsCogMultiple24Regular } from "@fluentui/react-icons";
import useAuth from "@utils/auth/useAuth";
import ProfileDrawer from "./ProfileDrawer";
import { useState } from "react";

export default function UserPopover() {
  const { user } = useAuth();

  const [isProfileDrawerOpen, setIsProfileDrawerOpen] = useState<boolean>(false);

  return (
    <div
      style={{
        display: 'flex',
        flexDirection: 'column',
        gap: 16
      }}
    >
      <ProfileDrawer
        isOpen={isProfileDrawerOpen}
        setIsOpen={setIsProfileDrawerOpen}
      />

      <div style={{
        display: 'flex',
        justifyContent: 'center',
        alignItems: 'center',
        paddingBottom: 20,
        width: 350
      }}>
        <span
          style={{
            flexShrink: 0,
            display: 'flex',
            justifyContent: 'center',
            alignItems: 'center',
            backgroundColor: tokens.colorBrandBackground,
            borderRadius: tokens.borderRadiusCircular,
            width: 36,
            height: 36,
            color: 'white',
          }}
        >
          { user?.level }
        </span>
        <div
          style={{
            position: 'relative',
            flex: 1,
            marginRight: 15
          }}
        >
          <ProgressBar
            value={user?.totalExperience}
            max={user?.totalXpForNextLevel}
          />
          <span
            style={{
              position: 'absolute',
              top: 5,
              width: '100%',
              textAlign: 'center'
            }}
          >
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