import axios, { AxiosError, AxiosRequestConfig } from "axios";

export const isAxiosError = <ErrorData = any>(
  data: unknown | AxiosError<ErrorData>,
): data is AxiosError<ErrorData> => data instanceof AxiosError;

export const request = async <ResponceData = any, ErrorData = any>(
  config: AxiosRequestConfig,
) => {
  try {
    return await axios.request<ResponceData>(config);
  } catch (e) {
    return e as AxiosError<ErrorData>;
  }
};
