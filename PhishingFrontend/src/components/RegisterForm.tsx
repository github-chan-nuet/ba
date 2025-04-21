import { Button, Field, Input, Toast, ToastTitle } from "@fluentui/react-components"
import { ChangeEvent, FormEvent, useState } from "react";
import { register } from "../api";
import { useToaster } from "../toaster/useToaster";

const RegisterForm = () => {
  const { dispatchToast } = useToaster();
  const [formData, setFormData] = useState({ firstname: "", lastname: "", email: "", password: "" });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
    console.log(formData);
  }

  const handleRegister = async (event: FormEvent) => {
    event.preventDefault();
    try {
      const result = await register();
      console.log(result);
    } catch (e) {
      dispatchToast(<Toast>
        <ToastTitle>Aktion fehlgeschlagen!</ToastTitle>
      </Toast>);
      console.error(e);
    }
    return;
  }

  return (
    <form
      onSubmit={handleRegister}
      style={{
        display: 'flex',
        flexDirection: 'column',
        gap: 16
      }}
    >
      <Field label="Vorname">
        <Input name="firstname" type="text" onChange={handleChange} />
      </Field>
      <Field label="Nachname">
        <Input name="lastname" type="text" onChange={handleChange} />
      </Field>
      <Field label="E-Mail">
        <Input name="email" type="email" onChange={handleChange} />
      </Field>
      <Field label="Passwort">
        <Input name="password" type="password" onChange={handleChange} />
      </Field>
      <div
        style={{
          display: 'flex',
          gap: 16
        }}
      >
        <Button type="submit" appearance="primary">Registrieren</Button>
      </div>
    </form>
  )
}

export default RegisterForm;