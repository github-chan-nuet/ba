import { Button, Drawer, DrawerBody, DrawerHeader, DrawerHeaderTitle } from "@fluentui/react-components";
import { Dismiss24Regular } from "@fluentui/react-icons";
import { useEffect, useState } from "react";
import LoginForm from "../LoginForm";
import RegisterForm from "../RegisterForm";

type AuthDrawerProps = {
  isOpen: boolean,
  setIsOpen: (value: React.SetStateAction<boolean>) => void
};

export default function AuthDrawer({ isOpen, setIsOpen }: AuthDrawerProps) {
  const [authContent, setAuthContent] = useState("login");

  useEffect(() => {
    if (isOpen === false) {
      setAuthContent("login");
    }
  }, [isOpen]);

  return (
    <Drawer
      position="end"
      open={isOpen}
      onOpenChange={(_, { open }) => setIsOpen(open)}
      style={{ width: "500px" }}
    >
      <DrawerHeader>
        <DrawerHeaderTitle
          action={
            <Button
              appearance="subtle"
              aria-label="close"
              icon={<Dismiss24Regular />}
              onClick={() => setIsOpen(false)}
            />
          }
        >
          { authContent === "login" ?
            "Login"
            : (authContent === "register") ?
            "Registrieren"
            : ""
          }
        </DrawerHeaderTitle>
      </DrawerHeader>
      <DrawerBody>
        { authContent === "login" ? (
          <LoginForm onSwitchToRegister={() => setAuthContent("register")} />
        ) : (authContent === "register") ? (
          <RegisterForm />
        ) : (
          <></>
        )}
      </DrawerBody>
    </Drawer>
  )
}