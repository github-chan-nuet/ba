import { useToastController } from "@fluentui/react-components";

const GLOBAL_TOASTER_ID = "global-toaster"

const useToaster = () => useToastController(GLOBAL_TOASTER_ID);

export { useToaster, GLOBAL_TOASTER_ID };