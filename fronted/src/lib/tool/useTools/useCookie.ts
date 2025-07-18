import Cookies from "js-cookie";
import { ref, watch } from "vue";

const useCookie = (
  cookieName: string,
  defaultValue: string,
  options?: Cookies.CookieAttributes,
) => {
  const cookie = ref(Cookies.get(cookieName) ?? defaultValue);

  function setCookie(value: string) {
    Cookies.set(cookieName, value, options);
  }

  watch(cookie, (newValue) => setCookie(newValue));

  return { cookie };
};

export default useCookie;
