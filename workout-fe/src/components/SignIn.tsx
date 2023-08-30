import { useDispatch, useSelector } from 'react-redux';
import { RootState } from '@/redux/store';
import { loginFailure, loginSuccess } from '@/redux/features/userSlice';
import { BaseUrl } from '@/constants';
import { loginUserAPI } from '@/api/apiServices';
import Error, { DisplayErrorProps } from './ErrorMessage';
import Form from './styles/Form';
import useForm from '../lib/useForm';
import { ErrorDataProps } from '@/types/userType';

export default function SignIn() {
  const dispatch = useDispatch()
  const isLoading = useSelector((state: RootState) => state.userReducer.isLoading)
  const errorData = useSelector((state: RootState) => state.userReducer.error)
  const { inputs, handleChange, resetForm } = useForm({
    Username: '',
    Password: '',
  });
  async function handleSubmit(e: React.FormEvent<HTMLFormElement>) {
    e.preventDefault(); // stop the form from submitting
    console.log('before', BaseUrl);
    try {
      const response = await loginUserAPI({ username: inputs.Username || '', password: inputs.Password || '' })

      dispatch(loginSuccess(response))
      resetForm();
      // Send the email and password to the graphqlAPI

    } catch (error) {
      dispatch(loginFailure((error as ErrorDataProps).message))

    }
  }
  const error: DisplayErrorProps = {
    error: {
      message: '',
      error: {
        message: ''
      },
      networkError: {
        result: {
          errors: [
            {
              message: ''
            }
          ]
        }
      }

    }
  };
  return (
    <Form method="submit" onSubmit={handleSubmit}>
      <h2>Sign Into Your Account</h2>
      <Error error={error.error} />
      <fieldset>
        <label htmlFor="Username">
          Email
          <input
            type="Username"
            name="Username"
            placeholder="Your Email Address"
            value={inputs.Username}
            onChange={handleChange}
          />
        </label>
        <label htmlFor="Password">
          Password
          <input
            type="Password"
            name="Password"
            placeholder="Password"
            autoComplete="password"
            value={inputs.Password}
            onChange={handleChange}
          />
        </label>
        <button type="submit">Sign In!</button>
      </fieldset>
      {errorData && <p>{errorData}</p>}
      {isLoading && <p>Loading...</p>}
    </Form>
  );
}
