import { useLocation } from "react-router-dom";
import styles from "./Recommendation.module.scss";

interface INavigateState {
  answer: boolean;
}

export const Recommendation = () => {
  const location = useLocation();

  const { answer } = location.state as INavigateState;

  return (
    <div className={styles.recommendation}>
      {answer ? (
        <h2 className={styles.title}>
          Here's the recommendation for the requested patient:{" "}
          {answer ? <span className={styles.okText}>Everything is ok</span> : <span className={styles.errorText}>Something is wrong</span>}
        </h2>
      ) : (
        <h2 className={styles.title}>No recommendation available.</h2>
      )}
    </div>
  );
};
