import { useLocation } from "react-router-dom";

interface INavigateState {
  data: string;
}

export const Analytics = () => {
  const location = useLocation();

  const { data } = location.state as INavigateState;

  console.log(data);

  return <div dangerouslySetInnerHTML={{ __html: data }} />;
};
