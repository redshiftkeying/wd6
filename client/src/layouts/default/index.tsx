
import React, { useState } from 'react';
import { Layout, Menu } from 'antd';
import { useHistory, } from 'react-router-dom'
import 'antd/dist/antd.css'
import './index.scss'
import {
    MenuUnfoldOutlined,
    MenuFoldOutlined,
    UserOutlined,
    VideoCameraOutlined,
    UploadOutlined,
} from '@ant-design/icons';

const UserLayout: React.FunctionComponent = (props) => {
    const { Header, Sider, Content } = Layout;
    const [collapsed, setCollapsd] = useState(false)
    const history = useHistory()


    return <Layout className="main-layout">
        <Sider trigger={null} collapsible collapsed={collapsed}>
            <div className="logo" />
            <Menu theme="dark" mode="inline" defaultSelectedKeys={['1']}>
                <Menu.Item key="1" icon={<UserOutlined />} onClick={() => { history.push('/task/1') }}>
                    list-temp
            </Menu.Item>
                <Menu.Item key="2" icon={<VideoCameraOutlined />} onClick={() => { history.push('/user/home') }}>
                    home
            </Menu.Item>
                <Menu.Item key="3" icon={<VideoCameraOutlined />} onClick={() => { history.push('/user/dashbord') }}>
                    Dashboard
            </Menu.Item>
                <Menu.Item key="4" icon={<UploadOutlined />} onClick={() => { history.push('/404') }}>
                    404
            </Menu.Item>
            </Menu>
        </Sider>
        <Layout className="site-layout">
            <Header className="site-layout-background" >
                {collapsed ? <MenuUnfoldOutlined onClick={() => {
                    setCollapsd(!collapsed)
                }} /> : <MenuFoldOutlined onClick={() => {
                    setCollapsd(!collapsed)
                }} />}
            </Header>
            <Content className="site-layout-content" >
                {props.children}
            </Content>
        </Layout>
    </Layout>
}
export default UserLayout;