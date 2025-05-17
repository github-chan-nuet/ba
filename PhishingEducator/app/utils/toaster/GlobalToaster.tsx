import { Toaster } from "@fluentui/react-components";
import { GLOBAL_TOASTER_ID } from "./useToaster";

export default function GlobalToaster() {
  return (
    <Toaster toasterId={GLOBAL_TOASTER_ID} />
  )
}