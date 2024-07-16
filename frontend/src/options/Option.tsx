
import { Events, WML } from "@wailsio/runtime";
import { Layout, Menu, theme, type MenuProps } from 'antd';
import { Content, Footer, Header } from "antd/es/layout/layout";
import Sider from "antd/es/layout/Sider";
import React, { useEffect, useState } from 'react';
import About from "../about/About";
import General from "../general/OptionGeneral";

type MenuItem = Required<MenuProps>['items'][number];

const items: MenuItem[] = [
  { key: 'General', label: 'General' },
  { key: 'Keymap', label: 'Keymap' },
  { key: 'About', label: 'About' },
];

const contentsMap: { [key: string]: React.FC } = {
  "General": General,
  "About": About,
};

function GetCurrentComponent(key: string): React.FC {
  return contentsMap[key] || (() => <div>not complete yet!</div>);
}

const Option: React.FC = () => {
  const [, setTime] = useState<string>('Listening for Time event...');
  const [current, setCurrent] = useState('General');
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
    WML.Reload();
  }, []);

  const CurrentComponent = GetCurrentComponent(current);

  return (
    <Layout className="h-screen" style={{ minHeight: '100vh' }}>
      <Sider collapsible collapsed={collapsed} onCollapse={(value) => setCollapsed(value)}>
        <div className="demo-logo-vertical" />
        <Menu theme="dark" defaultSelectedKeys={['General']} onClick={onClick} mode="inline" items={items} />
      </Sider>
      <Layout>
        <Header style={{ padding: 0, background: colorBgContainer }} />
        <Content style={{ margin: '0 16px' }}>
          <div
            style={{
              padding: 24,
              minHeight: 360,
              background: colorBgContainer,
              borderRadius: borderRadiusLG,
            }}
          >
            <CurrentComponent />
          </div>
        </Content>
        <Footer style={{ textAlign: 'center' }}>
          Ant Design ©{new Date().getFullYear()} Created by Ant UED
        </Footer>
      </Layout>
    </Layout>
  );
}

export default Option;
