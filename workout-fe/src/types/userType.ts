export interface UserProps {
  username: string;
  email: string;
  password: string;
}

export interface UserLoginProps {
  username: string;
  password: string;
}

export interface UserResponseDataProps {
  username: string;
  email: string;
  password_changed_at: string;
}

export interface UserResponseProps {
  data: UserResponseDataProps
}

export interface UserLoginDataResponseProps {
  access_token: string;
  user: UserResponseDataProps
}

export interface UserLoginResponseProps {
  data: UserLoginDataResponseProps
}

export interface ErrorDataProps {
  message: string;
}
