import axios from "axios";
import { BaseUrl, PublicKey } from "@/constants";
import { UserLoginDataResponseProps, UserLoginProps, UserLoginResponseProps, UserProps, UserResponseDataProps, UserResponseProps } from "@/types/userType";

const api = axios.create({
  baseURL: process.env.BaseUrl,
})



export const registerUserAPI = async (userData: UserProps): Promise<UserResponseDataProps> => {
  try {
    const response: UserResponseProps = await api.post('/api/users', userData);
    return response.data; // Assuming your API returns the user data upon registration
  } catch (error) {
    throw error;
  }
};

export const loginUserAPI = async (credentials: UserLoginProps): Promise<UserLoginDataResponseProps> => {
  try {
    console.log('hello', BaseUrl)
    const response: UserLoginResponseProps = await api.post('/api/users/login', credentials);
    return response.data; // Assuming your API returns the user data upon successful login
  } catch (error) {
    throw error;
  }
};
