import { isAxiosError } from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { getRecommendation } from "../../api";
import { RECOMMENDATION } from "../../constants";

export interface IRecommendationForm {
  patient_name: string;
  indicator_name: "heart_rate" | "blood_pressure";
  indicators: Array<number>;
}

export const RecommendationForm = () => {
  const {
    register,
    handleSubmit,
    formState: { errors },
  } = useForm();

  const navigate = useNavigate();

  const onSubmit = async (fields: any) => {
    const data: IRecommendationForm = {
      patient_name: fields.patientName,
      indicator_name: fields.indicatorName,
      indicators: JSON.parse(fields.indicators) as Array<number>,
    };
    const resp = await getRecommendation(data);
    if (isAxiosError(resp)) {
      console.error("foo");
    }

    //@ts-ignore
    const respData = resp.data as { code: number; data: { answer: boolean } };

    navigate(RECOMMENDATION, { state: { answer: respData.data.answer } });
  };

  return (
    <form onSubmit={handleSubmit(onSubmit)}>
      <input {...register("patientName", { required: true })} />
      {errors.patientName && <p>Patient name is required.</p>}
      <input {...register("indicatorName", { required: true })} />
      {errors.indicatorName && <p>Indicator name is required.</p>}
      <input {...register("indicators", { required: true })} />
      {errors.indicators && <p>Indicator numbers are required.</p>}
      <input type="submit" />
    </form>
  );
};
