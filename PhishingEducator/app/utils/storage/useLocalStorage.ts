import { useState } from "react";

export default function useLocalStorage(key: string, initialValue: unknown) {
  const [value, setValue] = useState(() => {
    try {
      const item = typeof window !== "undefined" ? window.localStorage.getItem(key) : null;
      return item ? JSON.parse(item) : initialValue;
    } catch (e) {
      return initialValue;
    }
  });

  const setItem = (newVal: unknown) => {
    try {
      window.localStorage.setItem(key, JSON.stringify(value));
      setValue(newVal);
    } catch (e) {
    }
  };

  return [value, setItem];
}