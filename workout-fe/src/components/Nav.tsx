'use client'
import Link from 'next/link';
import NavStyles from './styles/NavStyles';
import { useSelector } from 'react-redux';
import { RootState } from '@/redux/store';

export default function Nav() {
  const user = useSelector((state: RootState) => state.userReducer.UserData)
  return <NavStyles>
    <Link href="/products">Products</Link>
    <Link href="/sell">Sell</Link>
    <Link href="/orders">Order</Link>
    {user?.access_token ? (

      <Link href="/account">Account</Link>
    ) : (

      <Link href="/login">Login</Link>
    )}
  </NavStyles>;
}
