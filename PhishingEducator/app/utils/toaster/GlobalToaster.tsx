import { GLOBAL_TOASTER_ID } from "./useToaster";
import { Toaster } from "@fluentui/react-components";

export default function GlobalToaster() {
  return (
    <Toaster toasterId={GLOBAL_TOASTER_ID} />
  )
}