import { API_URL } from "../../constants";
import { LoginCreds } from "../../providers/auth/auth";
import { request } from "../utils";

export const getAuthToken = (creds: LoginCreds) => {
  return request({
    method: "POST",
    url: `${API_URL}/login`,
    data: creds,
  });
};
