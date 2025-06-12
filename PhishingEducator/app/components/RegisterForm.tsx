import { Button, Field, Input } from "@fluentui/react-components";
import { createUser } from "@api/index";
import useAuth from "@utils/auth/useAuth";
import { useToaster } from "@utils/toaster/useToaster";
import { useState, type ChangeEvent, type FormEvent } from "react";
import ErrorToast from "./ErrorToast";

export default function RegisterForm() {
  const { dispatchToast } = useToaster();
  const { onLogin } = useAuth();
  const [formData, setFormData] = useState({ firstname: "", lastname: "", email: "", password: "" });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  }

  const handleRegister = async (event: FormEvent) => {
    event.preventDefault();
    const { error } = await createUser({ body: formData });
    if (error) {
      dispatchToast(
        <ErrorToast error={error} />,
        { intent: "error" }
      );
      return;
    }
    try {
      await onLogin(formData.email, formData.password);
    } catch (e) {
      dispatchToast(
        <ErrorToast error={e} />,
        { intent: "error" }
      );
    }
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
      <Field label="Vorname" required>
        <Input name="firstname" type="text" onChange={handleChange} />
      </Field>
      <Field label="Nachname" required>
        <Input name="lastname" type="text" onChange={handleChange} />
      </Field>
      <Field label="E-Mail" required>
        <Input name="email" type="email" onChange={handleChange} />
      </Field>
      <Field label="Passwort" required>
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