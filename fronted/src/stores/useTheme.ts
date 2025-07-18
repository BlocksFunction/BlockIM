import useCookie from "@/lib/tool/useTools/useCookie";
import { defineStore } from "pinia";
import { computed } from "vue";

enum ThemeType {
  Dark = "dark",
  Light = "light",
  System = "system",
}

const useTheme = defineStore("useTheme", () => {
  const theme = useCookie("theme", "light");

  const getTheme = computed(() => theme.cookie.value);

  function setTheme(type: ThemeType) {
    theme.cookie.value = type;
  }

  function toggleTheme() {
    if (theme.cookie.value == "dark") theme.cookie.value = "light";
    else theme.cookie.value = "dark";
  }

  return { getTheme, toggleTheme, setTheme };
});

export default useTheme;
