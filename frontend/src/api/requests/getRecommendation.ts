import { API_URL } from "../../constants";
import { IRecommendationForm } from "../../pages/Recommendation/Form";
import { request } from "../utils";

export const getRecommendation = (data: IRecommendationForm) => {
  console.log(data);
  return request({
    method: "POST",
    url: `${API_URL}/api/v0/recommendation`,
    data: data,
  });
};
