import { createSlice, createAsyncThunk } from "@reduxjs/toolkit";
import { registerUserAPI, loginUserAPI } from '../../api/apiServices';
import { ErrorDataProps, UserLoginDataResponseProps, UserLoginProps, UserProps, UserResponseDataProps, UserResponseProps } from "@/types/userType";

// Define your state interface
interface UserState {
  UserData: null | UserLoginDataResponseProps; // Replace 'User' with your user data type
  isLoading: boolean;
  error: null | string;
}

// Define your initial state
const initialState: UserState = {
  UserData: {
    access_token: '',
    user: {} as UserResponseDataProps,
  },
  isLoading: false,
  error: null,
};

// Create an async thunk for user registration
const registerUser = createAsyncThunk<UserResponseDataProps, UserProps>(
  'user/register',
  async (userData) => {
    try {
      const response = await registerUserAPI(userData);
      return response; // Return the response as-is
    } catch (error) {
      throw error;
    }
  }
);

// Create an async thunk for user login
const loginUser = createAsyncThunk<UserLoginDataResponseProps, UserLoginProps>(
  'user/login',
  async (credentials) => {
    try {
      loginRequest()
      const response = await loginUserAPI(credentials)

      loginSuccess(response)
      return response; // Assuming the API returns user data upon login
    } catch (error) {
      loginFailure((error as ErrorDataProps).message)
      throw error;
    }
  }
);

export const userSlice = createSlice({
  name: "user",
  initialState,
  reducers: {
    loginRequest: (state) => {
      console.log('loading')
      state.isLoading = true;
      state.error = null;
    },
    loginSuccess: (state, action) => {
      state.UserData = action.payload;
      state.isLoading = false;
      state.error = null;
    },
    loginFailure: (state, action) => {
      console.log('failure')
      state.isLoading = false;
      state.error = action.payload;
    },
  },
  extraReducers(builder) {
    builder
      .addCase(registerUser.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(registerUser.fulfilled, (state, action) => {
        state.isLoading = false;
        state.UserData = { user: action.payload, access_token: '' }
      })
      .addCase(registerUser.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.error.message || '';
      })
      .addCase(loginUser.pending, (state) => {
        state.isLoading = true;
        state.error = null;
      })
      .addCase(loginUser.fulfilled, (state, action) => {
        state.isLoading = false;
        state.UserData = action.payload;
      })
      .addCase(loginUser.rejected, (state, action) => {
        state.isLoading = false;
        state.error = action.error.message || '';
      });
  },
},
)

export const { loginRequest, loginSuccess, loginFailure } = userSlice.actions;
export { registerUser, loginUser };
export default userSlice.reducer;
