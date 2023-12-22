import { API_URL } from "../../constants";
import { IAnalyticsForm } from "../../pages/Analytics/Form";
import { request } from "../utils";

export const getAnalytics = (data: IAnalyticsForm) => {
  console.log(data);
  return request({
    method: "POST",
    url: `${API_URL}/api/v0/tendency`,
    data: data,
  });
};
