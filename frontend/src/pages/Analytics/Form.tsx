import { isAxiosError } from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { getAnalytics } from "../../api";
import { ANALYTICS } from "../../constants";
import { saveAs } from "file-saver";

export interface IAnalyticsForm {
  patient_name: string;
  indicator_name: "heart_rate" | "blood_pressure";
}

export const AnalyticsForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const navigate = useNavigate();

  const onSubmit = async (fields: any) => {
    const data: IAnalyticsForm = {
      patient_name: fields.patientName,
      indicator_name: fields.indicatorName,
    };
    const resp = await getAnalytics(data);
    if (isAxiosError(resp)) {
      console.error("foo");
    }

    //@ts-ignore
    const respData = resp.data;

    const blob = new Blob([respData], { type: "text/html" });
    saveAs(blob, "htmlFile.html");

    navigate(ANALYTICS, {
      state: { data: "Report was saved on your computer" },
    });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register("patientName", { required: true })} />
      {errors.patientName && <p>Patient name is required.</p>}
      <input {...register("indicatorName", { required: true })} />
      {errors.indicatorName && <p>Indicator name is required.</p>}
      <input type="submit" />
    </form>
  );
};
