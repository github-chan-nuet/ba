import { Button, Checkbox, Divider, Drawer, DrawerBody, DrawerHeader, DrawerHeaderTitle, Field, Input, Spinner } from "@fluentui/react-components";
import { Dismiss24Regular } from "@fluentui/react-icons";
import useAuth from "../utils/auth/useAuth";
import { updateUser } from "../api";
import { useState, type FormEvent } from "react";

type ProfileDrawerProps = {
  isOpen: boolean;
  setIsOpen: (value: React.SetStateAction<boolean>) => void
};

export default function ProfileDrawer({ isOpen, setIsOpen }: ProfileDrawerProps) {
  const { user } = useAuth();
  const [isLoading, setIsLoading] = useState<boolean>(false);

  const handleSave = async (event: FormEvent) => {
    event.preventDefault();

    if (isLoading) return;

    setIsLoading(true);
    const { error } = await updateUser({
      path: {
        userId: ''
      }
    });
    if (error) {
      console.error("Failed to update user", error);
    } 
    setIsLoading(false);
  } 

  return (
    <Drawer
      position="end"
      open={isOpen}
      onOpenChange={(_, { open }) => setIsOpen(open)}
      style={{ width: 500 }}
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
          Profil-Einstellungen
        </DrawerHeaderTitle>
      </DrawerHeader>
      <DrawerBody>
        <form
          onSubmit={handleSave}
          style={{
            display: 'flex',
            flexDirection: 'column',
            gap: 16
          }}
        >
          <Field label="Vorname">
            <Input name="firstname" type="text" value={user?.firstname} />
          </Field>
          <Field label="Nachname">
            <Input name="lastname" type="text" value={user?.lastname} />
          </Field>
          <Field label="E-Mail">
            <Input name="email" type="email" value={user?.email} disabled />
          </Field>
          <Divider />
          <Checkbox name="phishing-simulation" type="checkbox" label="Ich möchte von der automatisieren Phishing Simulation profitieren." />

          <div
            style={{
              display: 'flex',
              gap: 16
            }}
          >
            <Button type="submit" appearance="primary">
              { isLoading ? (
                <Spinner size="tiny" />
              ) : (
                'Speichern'
              )}
            </Button>
          </div>
        </form>
      </DrawerBody>
    </Drawer>
  );
}