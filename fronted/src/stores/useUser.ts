import useCookie from "@/lib/tool/useTools/useCookie";
import { defineStore } from "pinia";
import { computed } from "vue";

const useUser = defineStore("useUser", () => {
  const token = useCookie("token", "");

  const getToken = computed(() => token.cookie.value);

  function setUser(tokenValue: string) {
    token.cookie.value = tokenValue;
  }
  function updateToken(tokenValue: string) {
    token.cookie.value = tokenValue;
  }

  return {
    getToken,
    setUser,
    updateToken,
  };
});

export default useUser;
