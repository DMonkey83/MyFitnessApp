import { ChangeEvent, useEffect, useState } from 'react';

export interface LoginProps {
  Username: string;
  Password: string;
}

export interface RegistrationProps {
  Email: string;
  Username: string;
  Password: string;
}

export default function useForm(initial = { Username: '', Password: '' }) {
  // create a state object for our inputs
  const [inputs, setInputs] = useState<LoginProps | RegistrationProps>(initial);
  const initialValues = Object.values(initial).join('');

  useEffect(() => {
    // This function runs when the things we are watching change
    setInputs(initial);
  }, [initialValues]);

  // {
  //   name: 'wes',
  //   description: 'nice shoes',
  //   price: 1000
  // }

  function handleChange(e: React.ChangeEvent<HTMLInputElement>) {
    let { value, name } = e.target;
    setInputs({
      // copy the existing state
      ...inputs,
      [name]: value,
    });
  }

  function resetForm() {
    setInputs(initial);
  }

  function clearForm() {
    const blankState: unknown = Object.fromEntries(
      Object.entries(inputs).map(([key, value]) => [key, value = ''])
    );
    setInputs(blankState as LoginProps | RegistrationProps);
  }

  // return the things we want to surface from this custom hook
  return {
    inputs,
    handleChange,
    resetForm,
    clearForm,
  };
}
