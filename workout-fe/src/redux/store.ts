import { configureStore, ThunkAction, Action } from "@reduxjs/toolkit";
import userReducer from './features/userSlice';

export type AppThunk = ThunkAction<void, RootState, null, Action<string>>;
export const store = configureStore({
  reducer: {
    userReducer
  },
  devTools: process.env.NODE_ENV !== "production",
});

export type RootState = ReturnType<typeof store.getState>;
export type AppDispatch = typeof store.dispatch;
