import { useLocation } from "react-router-dom"

interface INavigateState {
    answer: boolean;
}

export const Recommendation = () => {
    const location = useLocation();

    const { answer } = location.state as INavigateState;

    console.log(answer);

    return (
        <div>
            {answer}
        </div>
    );
}
