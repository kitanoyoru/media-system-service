import { Link, Outlet, useFetcher, useRouteLoaderData } from "react-router-dom";

export const Layout = () => {
  return (
    <div>
      <h1>Auth Example using RouterProvider</h1>

      <p>
        This example demonstrates a simple login flow with three pages: a public
        page, a protected page, and a login page. In order to see the protected
        page, you must first login. Pretty standard stuff.
      </p>

      <p>
        First, visit the public page. Then, visit the protected page. You're not
        yet logged in, so you are redirected to the login page. After you login,
        you are redirected back to the protected page.
      </p>

      <p>
        Notice the URL change each time. If you click the back button at this
        point, would you expect to go back to the login page? No! You're already
        logged in. Try it out, and you'll see you go back to the page you
        visited just *before* logging in, the public page.
      </p>

      <AuthStatus />

      <ul>
        <li>
          <Link to="/">Home Page</Link>
        </li>
        <li>
          <Link to="/form/recommendation">Recommendation Page</Link>
        </li>
        <li>
          <Link to="/analytics">Analytics Page</Link>
        </li>
      </ul>

      <Outlet />
    </div>
  );
};

function AuthStatus() {
  console.log(useRouteLoaderData("root"));
  let { token } = useRouteLoaderData("root") as { token: string | null };
  let fetcher = useFetcher();

  if (!token) {
    return <p>You are not logged in.</p>;
  }

  let isLoggingOut = fetcher.formData != null;

  return (
    <div>
      <p>Welcome {token}!</p>
      <fetcher.Form method="post" action="/logout">
        <button type="submit" disabled={isLoggingOut}>
          {isLoggingOut ? "Signing out..." : "Sign out"}
        </button>
      </fetcher.Form>
    </div>
  );
}
