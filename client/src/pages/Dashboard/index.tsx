import React from 'react';
import { Button } from 'antd';
import { useSelector } from 'react-redux'
import { RootState } from '../../store'
import UserLayout from "../../layouts/default"
import './index.scss'

export default function Dashboard() {
  const token = useSelector((state: RootState) => state.token)
  return <div>
    <UserLayout>
    {token.token}
    <br />
    这里是主页233
    <br />
    <Button type="primary">这里是主页</Button>
    </UserLayout>
  </div>
}