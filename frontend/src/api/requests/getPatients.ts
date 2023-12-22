import { API_URL } from "../../constants";
import { request } from "../utils";

export const getPatients = (username: string) => {
  return request({
    method: "POST",
    url: `${API_URL}/api/v0/patients`,
    data: {
        username,
    },
  });
};
