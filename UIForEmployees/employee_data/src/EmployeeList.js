import React, { useState, useEffect } from 'react';

const EmployeeList = () => {
    const [employees, setEmployees] = useState([]);

    useEffect(() => {
        // Fetch data from your API
        fetch('http://localhost:8080/Employees')
          .then((response) => {
            if (!response.ok) {
              throw new Error(`HTTP error! Status: ${response.status}`);
            }
            return response.json();
          })
          .then((data) => {
            console.log("Data fetched:", data); // Debugging log
            setEmployees(data);
          })
          .catch((error) => console.error('Error fetching employees:', error));
      }, []);

    
  return (
    <div>
      <h2>Employees</h2>
      {employees.length > 0 ? (
        <ul>
          {employees.map((emp) => (
            <li key={emp.emp_no}>{emp.first_name} {emp.last_name}</li>
          ))}
        </ul>
      ) : (
        <p>No employees found or unable to fetch data.</p>
      )}
    </div>
  );
};

export default EmployeeList;
