import { isAxiosError } from "axios";
import { getAuthToken } from "../../api";

export interface LoginCreds {
  username: string | null;
  password: string | null;
}

interface AuthProvider {
  isAuthenticated: boolean;
  token: null | string;

  signin(creds: LoginCreds): Promise<void>;
  signout(): Promise<void>;
}

export const authProvider: AuthProvider = {
  isAuthenticated: false,
  token: null,

  async signin(creds: LoginCreds) {
    const resp = await getAuthToken(creds);

    if (isAxiosError(resp)) {
      console.error("foo");
    }

    //@ts-ignore
    const respData = resp.data as { code: number; data: { token: string } };

    console.log(respData);

    authProvider.isAuthenticated = true;
    authProvider.token = respData.data.token;
  },
  async signout() {
    await new Promise((r) => setTimeout(r, 500)); // fake delay
    authProvider.isAuthenticated = false;
    authProvider.token = "";
  },
};
