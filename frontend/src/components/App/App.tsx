import type { LoaderFunctionArgs } from "react-router-dom";
import {
  RouterProvider,
  createBrowserRouter,
  redirect,
} from "react-router-dom";
import {
  ANALYTICS,
  ANALYTICS_FORM,
  LOGIN,
  LOGOUT,
  MAIN,
  PATIENTS,
  RECOMMENDATION,
  RECOMMENDATION_FORM,
} from "../../constants";
import { Analytics } from "../../pages/Analytics";
import { AnalyticsForm } from "../../pages/Analytics/Form";

import { Login } from "../../pages/Login";
import Patients from "../../pages/Patients/Patients";
import { Recommendation, RecommendationForm } from "../../pages/Recommendation";
import { authProvider } from "../../providers/auth";
import { Layout } from "../Layout/Layout";

const router = createBrowserRouter([
  {
    id: "root",
    path: "/",
    loader() {
      return { token: authProvider.token };
    },
    Component: Layout,
    children: [
      {
        path: LOGIN,
        action: loginAction,
        loader: loginLoader,
        Component: Login,
      },
      {
        path: RECOMMENDATION_FORM,
        loader: protectedLoader,
        Component: RecommendationForm,
      },
      {
        path: RECOMMENDATION,
        loader: protectedLoader,
        Component: Recommendation,
      },
      {
        path: ANALYTICS_FORM,
        loader: protectedLoader,
        Component: AnalyticsForm,
      },
      {
        path: ANALYTICS,
        loader: protectedLoader,
        Component: Analytics,
      },
      {
        path: PATIENTS,
        loader: protectedLoader,
        Component: Patients,
      },
    ],
  },
  {
    path: LOGOUT,
    async action() {
      await authProvider.signout();
      return redirect(MAIN);
    },
  },
]);

export const App = () => {
  return (
    <RouterProvider router={router} fallbackElement={<p>Initial Load...</p>} />
  );
};

async function loginAction({ request }: LoaderFunctionArgs) {
  let formData = await request.formData();

  let username = formData.get("username") as string | null;
  let password = formData.get("password") as string | null;

  if (!(username || password)) {
    return {
      error: "You must provide a creds to log in",
    };
  }

  try {
    await authProvider.signin({ username, password });
  } catch (error) {
    return {
      error: "Invalid login attempt",
    };
  }

  let redirectTo = formData.get("redirectTo") as string | null;
  return redirect(redirectTo || "/");
}

async function loginLoader() {
  if (authProvider.isAuthenticated) {
    return redirect("/");
  }
  return null;
}

function protectedLoader({ request }: LoaderFunctionArgs) {
  if (!authProvider.isAuthenticated) {
    let params = new URLSearchParams();
    params.set("from", new URL(request.url).pathname);

    return redirect("/login?" + params.toString());
  }

  return null;
}
