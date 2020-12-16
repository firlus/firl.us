import Vue from "vue";
import Router from "vue-router";
import "bootstrap";

import ShortcutList from "./components/ShortcutList.vue";
import ShortcutCreator from "./components/ShortcutCreator.vue";
import Shortcut from "./components/Shortcut.vue";
import Login from "./components/Login.vue";

Vue.component("shortcut-list", ShortcutList);
Vue.component("shortcut-creator", ShortcutCreator);
Vue.component("shortcut", Shortcut);

const router = new Router({
  routes: [
    {
      path: "/",
      name: "Login",
      component: Login,
    },
    {
      path: "/list",
      name: "List",
      component: ShortcutList,
    },
  ],
});

new Vue({
  el: "#app",
  router: router,
  components: {
    ShortcutList,
    Shortcut,
  },
});
