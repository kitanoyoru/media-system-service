import { Link, Outlet, useFetcher, useRouteLoaderData } from "react-router-dom";
import { getUsernameFromToken } from "../../utils/auth";

import styles from "./Layout.module.scss";

export const Layout = () => {
  return (
    <div className={styles.container}>
      <h1>Medical system service</h1>

      <p>
        A medical service system is a digital platform designed to facilitate
        the delivery of healthcare services. It typically includes features such
        as appointment scheduling, medical records management, prescription
        management, and communication tools between patients and healthcare
        providers.
      </p>

      <p>
        First, visit the public page. Then, visit the protected page. You're not
        yet logged in, so you are redirected to the login page. After you login,
        you are redirected back to the protected page.
      </p>

      <p>
        For healthcare providers, the system offers tools to manage patient
        appointments, access and update patient records, prescribe medications
        electronically, and communicate with patients. It can streamline
        administrative tasks, improve efficiency, and enhance patient care by
        providing easy access to relevant medical information.
      </p>

      <AuthStatus />

      <ul className={styles.navLinks}>
        <li>
          <Link to="/form/recommendation">Recommendation Page</Link>
        </li>
        <li>
          <Link to="/form/analytics">Analytics Page</Link>
        </li>
        <li>
          <Link to="/patients">Patients Page</Link>
        </li>
      </ul>

      <Outlet />
    </div>
  );
};

function AuthStatus() {
  let { token } = useRouteLoaderData("root") as { token: string | null };
  let fetcher = useFetcher();

  if (!token) {
    return <p>You are not logged in.</p>;
  }

  let isLoggingOut = fetcher.formData != null;

  return (
    <div className={styles.authStatus}>
      <p>Welcome {getUsernameFromToken(token)}!</p>
      <fetcher.Form method="post" action="/logout">
        <button type="submit" disabled={isLoggingOut}>
          {isLoggingOut ? "Signing out..." : "Sign out"}
        </button>
      </fetcher.Form>
    </div>
  );
}
