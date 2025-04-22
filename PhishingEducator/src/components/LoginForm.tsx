import { ChangeEvent, FormEvent, useState } from "react";
import useAuth from "../auth/useAuth"
import { Button, Field, Input } from "@fluentui/react-components";

type LoginFormProps = {
  onSwitchToRegister?: () => void
}

const LoginForm = ({
  onSwitchToRegister = () => {}
}: LoginFormProps) => {
  const { onLogin } = useAuth();
  const [formData, setFormData] = useState({ email: "", password: "" });

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    const { name, value } = e.target;
    setFormData((prev) => ({ ...prev, [name]: value }));
  }

  const handleLogin = (event: FormEvent) => {
    event.preventDefault();
    return onLogin(formData.email, formData.password);
  }
  
  return (
    <>
      <form
        onSubmit={handleLogin}
        style={{
          display: 'flex',
          flexDirection: 'column',
          gap: 16
        }}
      >
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
          <Button type="submit" appearance="primary">
            Login
          </Button>
          <Button
            appearance="secondary"
            onClick={onSwitchToRegister}
          >
            Ich besitze noch kein Konto
          </Button>
        </div>
      </form>
    </>
  )
}

export default LoginForm