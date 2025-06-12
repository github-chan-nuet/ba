import { useState, type ChangeEvent, type FormEvent } from "react";
import useAuth from "@utils/auth/useAuth";
import { useToaster } from "@utils/toaster/useToaster";
import { Button, Field, Input } from "@fluentui/react-components";
import ErrorToast from "@components/ErrorToast";

import AuthFormStyles from './AuthForm.module.scss';

type LoginFormProps = {
  onSwitchToRegister?: () => void
};

export default function LoginForm({
  onSwitchToRegister = () => {}
}: LoginFormProps) {
  const { onLogin } = useAuth();
  const { dispatchToast } = useToaster();
  const [formData, setFormData] = useState({ email: "", password: ""});

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  }

  const handleLogin = async (event: FormEvent) => {
    event.preventDefault();
    try {
      await onLogin(formData.email, formData.password);
    } catch (e) {
      if (
        typeof e === "object" &&
        e !== null &&
        'title' in e &&
        typeof e.title === "string"
      ) {
        dispatchToast(
          <ErrorToast error={e} />,
          { intent: "error"}
        )
      }
    }
  }

  return (
    <>
      <form
        className={AuthFormStyles.AuthForm}
        onSubmit={handleLogin}
      >
        <Field label="E-Mail" required>
          <Input name="email" type="email" onChange={handleChange} />
        </Field>
        <Field label="Passwort" required>
          <Input name="password" type="password" onChange={handleChange} />
        </Field>
        <div className={AuthFormStyles.AuthForm__Actions}>
          <Button type="submit" appearance="primary">Login</Button>
          <Button appearance="secondary" onClick={onSwitchToRegister}>Ich besitze noch kein Konto</Button>
        </div>
      </form>
    </>
  )
}