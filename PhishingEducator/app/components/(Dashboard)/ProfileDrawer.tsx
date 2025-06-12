import { useState, type ChangeEvent, type FormEvent } from "react";
import useAuth from "@utils/auth/useAuth";
import { useToaster } from "@utils/toaster/useToaster";
import { Button, Checkbox, Divider, Drawer, DrawerBody, DrawerHeader, DrawerHeaderTitle, Field, Input, Spinner, Toast, ToastBody, ToastTitle } from "@fluentui/react-components";
import { Dismiss24Regular } from "@fluentui/react-icons";

import AuthFormStyles from '../AuthForm.module.scss';

type ProfileDrawerProps = {
  isOpen: boolean;
  setIsOpen: (value: React.SetStateAction<boolean>) => void
};

export default function ProfileDrawer({ isOpen, setIsOpen }: ProfileDrawerProps) {
  const { user, updateUser } = useAuth();
  const { dispatchToast } = useToaster();
  const [isLoading, setIsLoading] = useState<boolean>(false);
  const [showSpinner, setShowSpinner] = useState<boolean>(false);
  const [formData, setFormData] = useState({
    firstname: user?.firstname,
    lastname: user?.lastname, 
    email: user?.email,
    password: "",
    participatesInPhishingSimulation: user?.participatesInPhishingSimulation
  });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, checked, type, value } = e.target;
    setFormData((prev) => ({
      ...prev,
      [name]: type === 'checkbox' ? checked : value
    }));
  }

  const handleSave = async (event: FormEvent) => {
    event.preventDefault();

    if (isLoading || !user?.id) return;

    setIsLoading(true);
    const spinnerTimeout = setTimeout(() => {
      setShowSpinner(true);
    }, 250);

    try {
      await updateUser({
        firstname: formData.firstname,
        lastname: formData.lastname,
        participatesInPhishingSimulation: formData.participatesInPhishingSimulation
      });
      dispatchToast(
        <Toast>
          <ToastTitle>Profil-Einstellungen gespeichert!</ToastTitle>
        </Toast>,
        { intent: "success" }
      )
    } catch (e) {
      console.error(e)
      dispatchToast(
        <Toast>
          <ToastTitle>Es ist ein Fehler aufgetreten!</ToastTitle>
          <ToastBody>Bitte versuchen Sie es später erneut.</ToastBody>
        </Toast>,
        { intent: "error" }
      );
    }

    clearTimeout(spinnerTimeout);
    setShowSpinner(false);
    setIsLoading(false);
  } 

  return (
    <Drawer
      className={AuthFormStyles.Drawer}
      position="end"
      open={isOpen}
      onOpenChange={(_, { open }) => setIsOpen(open)}
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
          className={AuthFormStyles.AuthForm}
          onSubmit={handleSave}
        >
          <Field label="Vorname">
            <Input name="firstname" type="text" value={formData.firstname} onChange={handleChange} />
          </Field>
          <Field label="Nachname">
            <Input name="lastname" type="text" value={formData.lastname} onChange={handleChange} />
          </Field>
          <Field label="E-Mail">
            <Input name="email" type="email" value={formData.email} disabled={true} />
          </Field>
          <Divider />
          <Checkbox
            name="participatesInPhishingSimulation"
            type="checkbox"
            defaultChecked={formData.participatesInPhishingSimulation}
            onChange={handleChange}
            label="Ich möchte von der automatisieren Phishing Simulation profitieren."
          />

          <div className={AuthFormStyles.AuthForm__Actions}>
            <Button type="submit" appearance="primary">
              { showSpinner ? (
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