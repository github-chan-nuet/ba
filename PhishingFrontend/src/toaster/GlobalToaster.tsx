import { Toaster } from "@fluentui/react-components"
import { GLOBAL_TOASTER_ID } from "./useToaster"

const GlobalToaster = () => {
  return (
    <Toaster toasterId={GLOBAL_TOASTER_ID} />
  )
}

export default GlobalToaster;