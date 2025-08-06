import { createWebHashHistory, createRouter } from "vue-router";
import HomePage from "./pages/HomePage.vue";
import AboutPage from "./pages/AboutPage.vue";
import LogsPage from "./pages/LogsPage.vue";
import RouteListPage from "./pages/ManageRoute/RouteListPage.vue";
import RouteCreatePage from "./pages/ManageRoute/RouteCreatePage.vue";
import RouteUpdatePage from "./pages/ManageRoute/RouteUpdatePage.vue";

const routes = [
  { path: "/", component: HomePage },
  { path: "/about", component: AboutPage },
  // manage routes
  { path: "/routes", component: RouteListPage },
  { path: "/routes/new", component: RouteCreatePage },
  { path: "/routes/:id/edit", component: RouteUpdatePage },
  // logs
  { path: "/logs", component: LogsPage },
];

const router = createRouter({
  history: createWebHashHistory(),
  routes,
});

export default router;
