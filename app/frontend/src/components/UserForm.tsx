import React, { useState, useEffect } from 'react';

interface User {
  id: number;
  name: string;
  age: number;
}

interface UserFormProps {
  user?: User | null;
  isCreateMode: boolean;
  onSave: (formData: User) => void;
  closeForm: () => void;
}

const UserForm: React.FC<UserFormProps> = ({ user, isCreateMode, onSave ,closeForm }) => {

  const [formData, setFormData] = useState<User>({ id: 0, name: '', age: 0 });

  useEffect(() => {
    if (user) {
      setFormData(user);
    } else {
      setFormData({ id: 0, name: '', age: 0 });
    }
  }, [user]);

  const handleInputChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData((prevData) => ({ ...prevData, [e.target.name]: e.target.value }));
  };

  const handleSaveClick = () => {
    onSave(formData);
  };

  return (
    <div className='div-form'>
      <h2>{isCreateMode ? 'New User' : 'Edit User'}</h2>
      <form>
        <button className='close-button' type="button" onClick={closeForm}>
          Cancel
        </button>
        <input type="text" name="id" value={formData.id || ''} hidden />
        <label>
          Name:
          <input type="text" name="name" value={formData.name} onChange={handleInputChange} />
        </label>
        <br />
        <label>
          Age:
          <input type="number" name="age" value={formData.age || ''} onChange={handleInputChange} />
        </label>
        <br />
        <button type="button" onClick={handleSaveClick}>
          {isCreateMode ? 'Create' : 'Update'}
        </button>
      </form>
    </div>
  );
};

export default UserForm;