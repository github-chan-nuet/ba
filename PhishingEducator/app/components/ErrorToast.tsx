import { Toast, ToastTitle } from "@fluentui/react-components";
import type { ProblemDetail } from "app/api";

const isProblemDetail = (error: unknown) => {
  return (
    typeof error === "object" &&
    error !== null &&
    "title" in error &&
    typeof error.title === "string"
  );
}

type ErrorToastProps = {
  error: unknown;
};

export default function ErrorToast({ error }: ErrorToastProps) {
  return (
    <Toast>
      <ToastTitle>{ isProblemDetail(error) ? (error as ProblemDetail).title : 'Unerwarteter Fehler' }</ToastTitle>
    </Toast>
  );
}