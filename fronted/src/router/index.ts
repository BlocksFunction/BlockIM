import useCookie from "@/lib/tool/useTools/useCookie";
import Auth from "@/router/modules/Auth";
import ChatView from "@/views/ChatView.vue";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "Chat",
      component: ChatView,
    },
    ...Auth,
  ],
});

router.beforeEach((to, from, next) => {
  if (to.name === "Login" || to.name === "Register") {
    next();
  } else {
    const token = useCookie("token", "");

    if (token.cookie.value != "") {
      next();
    } else {
      next({ name: "Login" });
    }
  }
});

export default router;
