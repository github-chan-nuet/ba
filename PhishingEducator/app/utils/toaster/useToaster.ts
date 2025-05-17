import { useToastController } from "@fluentui/react-components";

export const GLOBAL_TOASTER_ID = "global-toaster";
export function useToaster() {
  return useToastController(GLOBAL_TOASTER_ID);
}