// pages/index.tsx
import PopUp from '@/components/PopUp';
import UserForm from '@/components/UserForm';
import { MantineProvider } from '@mantine/core';
import React, { useState, useEffect } from 'react';
import useSWR from "swr";

interface User {
  id: number;
  name: string;
  age: number;
}

// API endpoint for data fetching
export const ENDPOINT = "http://localhost:8080";

// Fetcher function for SWR data fetching
const fetcher = (url: string) =>
  fetch(`${ENDPOINT}/${url}`).then((r) => r.json());

const Home: React.FC = () => {

  const { data , mutate} = useSWR('users', fetcher)

  // State variables
  const [selectedUser, setSelectedUser] = useState<User | null>(null);
  const [isCreateMode, setIsCreateMode] = useState<boolean>(false);

  const handleNewButtonClick = () => {
    setSelectedUser(null);
    setIsCreateMode(true);
  };

  const handleEditButtonClick = (user: User) => {
    setSelectedUser(user);
    setIsCreateMode(false);
  };
  
  const handleSaveButtonClick = async (formData: User) => {
    const { id, ...rest } = formData;
    rest.age = Number(formData.age);
    if (isCreateMode) {
      const response = await fetch('http://localhost:8080/users', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(rest),
      });
    } else {
        await fetch(`http://localhost:8080/users/${formData.id}`, {
        method: 'PUT',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(rest),
      });
    }

    mutate()
    
    handleClose()
  };
  
  const handleDeleteButtonClick = async (user: User) => {
    const response = await fetch(`http://localhost:8080/users/${user.id}`, {
      method: 'DELETE',
      headers: {
        'Content-Type': 'application/json',
      },
    });
    mutate()
  };
  
  const handleClose = () => {
        setSelectedUser(null);
        setIsCreateMode(false);
  };

  // useEffect for handling side effects when user data changes
  useEffect(() => {
    // Code to be executed when the component is mounted or user data changes
  }, [data]);


  return (
    <div>
      <h1>User List</h1>
      <table>
        <thead>
          <tr>
            <th >ID</th>
            <th >Name</th>
            <th >Age</th>
            <th >Actions</th>
          </tr>
        </thead>
        <tbody>
          {data?.map((user: User) => (
            <tr key={user.id}>
              <td>{user.id}</td>
              <td>{user.name}</td>
              <td>{user.age}</td>
              <td >
                <button className='button-edit' onClick={() => handleEditButtonClick(user)}>Edit</button>
                <button className='button-delete' onClick={() => handleDeleteButtonClick(user)}>Delete</button>
              </td>
            </tr>
          ))}
        </tbody>
      </table>
      <button onClick={handleNewButtonClick}>New</button>

      {(isCreateMode || selectedUser) && (
        <MantineProvider>

          <UserForm
            user={selectedUser}
            isCreateMode={isCreateMode}
            onSave={handleSaveButtonClick}
            closeForm={handleClose} />
        </MantineProvider>
      )}
    </div>
  );
};

export default Home;