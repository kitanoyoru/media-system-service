import { useState, useEffect } from "react";
import { useRouteLoaderData } from "react-router-dom";
import { getPatients } from "../../api";
import { IPatient } from "../../models";
import { getUsernameFromToken } from "../../utils";

const Patients = () => {
  const [patients, setPatients] = useState<IPatient[]>([]);

  let { token } = useRouteLoaderData("root") as { token: string | null };

  useEffect(() => {
    const fetchData = async () => {
      try {
        const response = await getPatients(getUsernameFromToken(token!));
        //@ts-ignore
        console.log(response.data.data.patients)
        setPatients(response.data.data.patients);
      } catch (error) {
        console.error("Error fetching patients:", error);
      }
    };

    fetchData();
  }, []);

  return (
    <div>
      <h2>Patients</h2>
      {patients.length > 0 ? (
        <table>
          <thead>
            <tr>
              <th>ID</th>
              <th>Name</th>
            </tr>
          </thead>
          <tbody>
            {patients.map((patient, id) => (
              <tr key={id}>
                <td>{id}</td>
                <td>{patient.name}</td>
              </tr>
            ))}
          </tbody>
        </table>
      ) : (
        <p>No patients available.</p>
      )}
    </div>
  );
};

export default Patients;
