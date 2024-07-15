import { Events, WML } from "@wailsio/runtime";
import { Layout, Menu, theme, type MenuProps } from 'antd';
import { Content, Footer, Header } from "antd/es/layout/layout";
import Sider from "antd/es/layout/Sider";
import React, { useEffect, useState } from 'react';
import About from "../about/About";

type MenuItem = Required<MenuProps>['items'][number];


const items: MenuItem[] = [
  { key: 'General', label: 'General' },
  { key: 'Keymap', label: 'Keymap' },
  { key: 'About', label: 'About' },
];

const contentsMap: object = {
  "About": About,
  "_Default": <div>not complete yet!</div>
}

function GetCurrentMap(key: string): any{
  return contentsMap.hasOwnProperty(key) ? contentsMap[key]() : contentsMap["_Default"]
}

const Option: React.FC = (props) => {
  const [, setTime] = useState<string>('Listening for Time event...');
  // const [theme, setTheme] = useState<MenuTheme>('dark');
  const [current, setCurrent] = useState('1');

  // const changeTheme = (value: boolean) => {
  //   setTheme(value ? 'dark' : 'light');
  // };

  // const optionMap: { [key: string]: any } = {
  //   "general": General,
  //   "anothergeneral": General,
  // };
  // const optionNames = Object.keys(optionMap)

  const [collapsed, setCollapsed] = useState(false);
  const {
    token: { colorBgContainer, borderRadiusLG },
  } = theme.useToken();


  const onClick: MenuProps['onClick'] = (e) => {
    console.log('click ', e);
    setCurrent(e.key);
  };

  useEffect(() => {
    Events.On('time', (timeValue: any) => {
      setTime(timeValue.data);
    });
    // Reload WML so it picks up the wml tags
    WML.Reload();
  }, []);



  return (
    <Layout className="h-screen" style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <div className="demo-logo-vertical" />
        <Menu theme="dark" defaultSelectedKeys={['1']} onClick={onClick} mode="inline" items={items} />
        {/* <Footer>fuck</Footer> */}
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }} />
        <Content style={{ margin: '0 16px' }}>
          {/* <Breadcrumb style={{ margin: '16px 0' }}>
            <Breadcrumb.Item>User</Breadcrumb.Item>
            <Breadcrumb.Item>Bill</Breadcrumb.Item>
          </Breadcrumb> */}
          <div
            style={{
              padding: 24,
              minHeight: 360,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            {/* Bill is a cat. */}
            {GetCurrentMap(current)}
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          Ant Design ©{new Date().getFullYear()} Created by Ant UED
        </Footer>
      </Layout>
    </Layout>
  );
}

export default Option
