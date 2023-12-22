import { useLocation } from "react-router-dom";

interface INavigateState {
  answer: boolean;
}

export const Recommendation = () => {
  const location = useLocation();

  const { answer } = location.state as INavigateState;

  return (
    <div>
      {answer ? (
        <h2>
          Here's the recommendation for requestes patient:{" "}
          {answer ? <div>Everything is ok</div> : <div>Smth wrong</div>}
        </h2>
      ) : (
        <h2>No recommendation available.</h2>
      )}
    </div>
  );
};
