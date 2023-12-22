import { isAxiosError } from "axios";
import { useForm } from "react-hook-form";
import { useNavigate } from "react-router-dom";
import { getRecommendation } from "../../api";
import { RECOMMENDATION } from "../../constants";
import classNames from "classnames";

import styles from "./Form.module.scss";

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
    <form onSubmit={handleSubmit(onSubmit)} className={styles.recommendationForm}>
      <div className={classNames(styles.formGroup, styles.label)}>
        <label htmlFor="patientName">Patient Name</label>
        <input
          {...register("patientName", { required: true })}
          className={classNames(styles.input, { [styles.inputError]: errors.patientName })}
        />
        {errors.patientName && <p className={styles.error}>Patient name is required.</p>}
      </div>

      <div className={classNames(styles.formGroup, styles.label)}>
        <label htmlFor="indicatorName">Indicator Name</label>
        <input
          {...register("indicatorName", { required: true })}
          className={classNames(styles.input, { [styles.inputError]: errors.indicatorName })}
        />
        {errors.indicatorName && <p className={styles.error}>Indicator name is required.</p>}
      </div>

      <div className={classNames(styles.formGroup, styles.label)}>
        <label htmlFor="indicators">Indicators</label>
        <input
          {...register("indicators", { required: true })}
          className={classNames(styles.input, { [styles.inputError]: errors.indicators })}
        />
        {errors.indicators && <p className={styles.error}>Indicator numbers are required.</p>}
      </div>

      <input type="submit" className={styles.submitButton} value="Submit" />
    </form>
  );
};
