import { isAxiosError } from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { getAnalytics } from "../../api";
import { ANALYTICS } from "../../constants";
import { saveAs } from "file-saver";
import styles from "./Form.module.scss";

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
    <form onSubmit={handleSubmit(onSubmit)} className={styles.recommendationForm}>
      <div className={styles.formGroup}>
        <label htmlFor="patientName" className={styles.label}>
          Patient Name
        </label>
        <input
          {...register("patientName", { required: true })}
          className={styles.input}
        />
        {errors.patientName && <p className={styles.error}>Patient name is required.</p>}
      </div>

      <div className={styles.formGroup}>
        <label htmlFor="indicatorName" className={styles.label}>
          Indicator Name
        </label>
        <input
          {...register("indicatorName", { required: true })}
          className={styles.input}
        />
        {errors.indicatorName && <p className={styles.error}>Indicator name is required.</p>}
      </div>

      <input type="submit" className={styles.submitButton} value="Submit" />
    </form>
  );
};
